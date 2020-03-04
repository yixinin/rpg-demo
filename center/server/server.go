package server

import (
	"rpg-demo/center/iface"
	"rpg-demo/config"
	"rpg-demo/lib/bus"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/log"
	"sync"
)

type Center struct {
	sync.RWMutex
	Players map[int64]iface.Player
	Rooms   map[int64]iface.Room
	Games   map[int32]map[int32]map[int64]iface.Room // [gameId][gameType][roomId]Room

	Grpc *GrpcPool
	Nats *bus.Bus //消息队列
	Etcd *etcd.Service
}

func NewCenter(conf *config.Config) *Center {
	n, err := bus.NewBus(conf.Nats)
	if err != nil {
		log.Error(err)
	}
	e, err := etcd.NewService("center", etcd.ServiceInfo{}, conf.Etcd)
	if err != nil {
		log.Error(err)
	}
	return &Center{
		Players: make(map[int64]iface.Player, 10000),
		Rooms:   make(map[int64]iface.Room, 1000),
		Games:   make(map[int32]map[int32]map[int64]iface.Room),
		Grpc:    NewGrpcPool(conf.Etcd),
		Nats:    n,
		Etcd:    e,
	}
}

func (c *Center) StartCenter() {
	go c.Etcd.Start()
	go c.Grpc.ListenGrpc()
	c.SubMessageQueue()
}

func (c *Center) SubMessageQueue() {
	_, err := c.Nats.Subscribe("rpg", c.handleMessageQueue)
	if err != nil {
		log.Error(err)
	}
}

func (s *Center) CreateRoom(r iface.Room) {
	s.Lock()
	defer s.Unlock()
	s.Rooms[r.GetRoomId()] = r
	if _, ok := s.Games[r.GetGameId()]; !ok {
		s.Games[r.GetGameId()] = make(map[int32]map[int64]iface.Room)
	}
	if s.Games[r.GetGameId()][r.GetGameType()] == nil {
		s.Games[r.GetGameId()][r.GetGameType()] = make(map[int64]iface.Room)
	}
	s.Games[r.GetGameId()][r.GetGameType()][r.GetRoomId()] = r
}

func (s *Center) DestroyRoom(r iface.Room) {
	s.Lock()
	defer s.Unlock()
	delete(s.Rooms, r.GetRoomId())
	if v, ok := s.Games[r.GetGameId()]; ok {
		if v1, ok := v[r.GetGameType()]; ok {
			delete(v1, r.GetRoomId())
		}
	}
}

func (s *Center) AddPlayer(user iface.Player) {
	s.Lock()
	defer s.Unlock()
	s.Players[user.GetPlayerId()] = user
}

func (s *Center) RemovePlayer(userid int64) {
	s.Lock()
	defer s.Unlock()
	delete(s.Players, userid)
}
