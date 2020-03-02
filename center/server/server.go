package server

import (
	"rpg-demo/center/intface"
	"sync"
)

var Server *Center

func init() {
	Server = NewCenter()
}

type Center struct {
	sync.RWMutex
	Players map[int64]intface.Player
	Rooms   map[int64]intface.Room
	Games   map[int32]map[int32]map[int64]intface.Room // [gameId][gameType][roomId]Room
}

func NewCenter() *Center {
	return &Center{
		Players: make(map[int64]intface.Player, 10000),
		Rooms:   make(map[int64]intface.Room, 1000),
		Games:   make(map[int32]map[int32]map[int64]intface.Room),
	}
}

func (s *Center) CreateRoom(r intface.Room) {
	s.Lock()
	defer s.Unlock()
	s.Rooms[r.GetRoomId()] = r
	if _, ok := s.Games[r.GetGameId()]; !ok {
		s.Games[r.GetGameId()] = make(map[int32]map[int64]intface.Room)
	}
	if s.Games[r.GetGameId()][r.GetGameType()] == nil {
		s.Games[r.GetGameId()][r.GetGameType()] = make(map[int64]intface.Room)
	}
	s.Games[r.GetGameId()][r.GetGameType()][r.GetRoomId()] = r
}

func (s *Center) DestroyRoom(r intface.Room) {
	s.Lock()
	defer s.Unlock()
	delete(s.Rooms, r.GetRoomId())
	if v, ok := s.Games[r.GetGameId()]; ok {
		if v1, ok := v[r.GetGameType()]; ok {
			delete(v1, r.GetRoomId())
		}
	}
}

func (s *Center) AddPlayer(user intface.Player) {
	s.Lock()
	defer s.Unlock()
	s.Players[user.GetPlayerId()] = user
}

func (s *Center) RemovePlayer(userid int64) {
	s.Lock()
	defer s.Unlock()
	delete(s.Players, userid)
}
