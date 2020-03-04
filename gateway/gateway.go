package main

import (
	"errors"
	"fmt"
	"rpg-demo/config"
	"rpg-demo/ip"
	"rpg-demo/lib/bus"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/log"
	"rpg-demo/protocol"
	"rpg-demo/utils"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

var ErrUserNotLogin = errors.New("user not login")

type Gateway struct {
	server *Server
	grpc   *GrpcPool
	nats   *bus.Bus //消息队列
	etcd   *etcd.Service
	// cc   chan *GrpcInfo //center GRPC 连接
	// gc   chan *GrpcInfo //game GRPC 连接
}

func NewGateway(conf *config.Config) *Gateway {
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

	return &Gateway{
		server: NewServer(),
		grpc:   NewGrpcPool(conf.Etcd),
		// cc:     make(chan *GrpcInfo),
		// gc:     make(chan *GrpcInfo),
		nats: n,
		etcd: e,
	}
}

func (g *Gateway) StartGateway(protocol string) {
	switch protocol {
	case "tcp":
		g.StartTcp()
	case "websocket", "": //默认websocket
		g.StartWebSocket()
	case "http":
		g.StartHttp()
	default:
		log.Error("unknown protocol")
		return
	}

	go g.etcd.Start()
	go g.grpc.ListenGrpc()
	g.SubMessageQueue()
}

func (g *Gateway) StartTcp() {
	queue := cellnet.NewEventQueue()

	// 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)

	// 将接受器Peer与tcp.ltv的处理器绑定，并设置事件处理回调
	// tcp.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(peerIns, "tcp.ltv", g.HandleCallback)

	// 启动Peer，服务器开始侦听
	peerIns.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	queue.StartLoop()
}

func (g *Gateway) StartWebSocket() {
	queue := cellnet.NewEventQueue()

	// 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	peerIns := peer.NewGenericPeer("gorillaws.Acceptor", "server", peerAddress, queue)

	// 将接受器Peer与tcp.ltv的处理器绑定，并设置事件处理回调
	// tcp.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(peerIns, "gorillaws.ltv", g.HandleCallback)

	// 启动Peer，服务器开始侦听
	peerIns.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	queue.StartLoop()
}

func (g *Gateway) StartHttp() {
	router := gin.Default()
	router.POST("/message", g.handleHttp)
}

func (g *Gateway) SubMessageQueue() {
	_, err := g.nats.Subscribe("rpg", g.handleMessageQueue)
	if err != nil {
		log.Error(err)
	}
}

func (g *Gateway) GetRpcConn(header *protocol.ReqHeader, service string) (*grpc.ClientConn, error) {
	ok, err := g.server.CheckUser(header)
	if !ok || err != nil {
		return nil, ErrUserNotLogin
	}
	switch service {
	case "center":
		centerIP, err := g.server.GetCenter(header.GetUid())
		if err != nil {
			log.Error("check token error:%v", err)
		}
		if centerIP == "" {
			log.Error("check token error, center ip is empty :%v", err)
		}

		rpcConn := g.grpc.Center(centerIP)
		if rpcConn == nil {
			return nil, fmt.Errorf("%s is not connected", centerIP)
		}
	case "game":
		gameIP, err := g.server.GetGame(header.GetUid())
		if err != nil {
			log.Error("check token error:%v", err)
		}
		if gameIP == "" {
			log.Error("check token error, game ip is empty :%v", err)
		}

		rpcConn := g.grpc.Game(gameIP)
		if rpcConn == nil {
			return nil, fmt.Errorf("%s is not connected", gameIP)
		}
	}

	return nil, fmt.Errorf("no such service :%s", service)
}
