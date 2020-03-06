package server

import (
	"com/config"
	"com/ip"
	"com/lib/bus"
	"com/lib/etcd"
	"com/lib/log"
	"com/lib/pool"
	"com/utils"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type Game struct {
	Grpc       *GrpcPool
	Nats       *bus.Bus //消息队列
	Etcd       *etcd.Service
	RoomIdPool *pool.IdPool
	TeamIdPool *pool.IdPool
}

func NewGame(conf *config.Config) *Game {
	n, err := bus.NewBus(conf.Nats)
	if err != nil {
		log.Error(err)
	}

	ip.ServerIP = fmt.Sprintf("%s:%s", utils.IPAddress(), conf.Grpc)
	var serviceInfo = etcd.ServiceInfo{
		IP: ip.ServerIP,
	}

	log.Info("grpc listen in %d", serviceInfo.IP)
	e, err := etcd.NewService("gateway", serviceInfo, conf.Etcd)
	if err != nil {
		log.Error("etcd conn error:%v", err)
	}
	var idpool = pool.NewIdPool(1000, 10000, CheckIdAvailable)
	return &Game{
		Grpc:       NewGrpcPool(conf.Etcd),
		Nats:       n,
		Etcd:       e,
		RoomIdPool: idpool,
	}
}

func (g *Game) StartGame() {
	go g.Etcd.Start()
	go g.Grpc.ListenGrpc()
	// g.SubMessageQueue()
}

func (g *Game) SubMessageQueue(f func(m *nats.Msg)) {
	_, err := g.Nats.Subscribe("rpg", f)
	if err != nil {
		log.Error(err)
	}
}

func CheckIdAvailable(n int64) bool {

	return true
}
