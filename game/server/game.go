package server

import (
	"fmt"
	"rpg-demo/config"
	"rpg-demo/ip"
	"rpg-demo/lib/bus"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/log"
	"rpg-demo/utils"
)

type Game struct {
	grpc *GrpcPool
	nats *bus.Bus //消息队列
	etcd *etcd.Service
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
	return &Game{
		grpc: NewGrpcPool(conf.Grpc),
		nats: n,
		etcd: e,
	}
}

func (g *Game) StartGame() {
	go g.etcd.Start()
	go g.grpc.ListenGrpc()
	g.SubMessageQueue()
}

func (g *Game) SubMessageQueue() {
	_, err := g.nats.Subscribe("rpg", g.handleMessageQueue)
	if err != nil {
		log.Error(err)
	}
}
