package gateway

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

type Gateway struct {
	server     *Server
	grpcClient *GrpcClient

	centerCh chan *GrpcConn
	gameCh   chan *GrpcConn
}

func NewGateway() *Gateway {
	return &Gateway{
		server: NewServer(),
	}
}

func (g *Gateway) StartServer() {
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

	g.grpcClient = NewGrpcClient()
	g.grpcClient.GrpcConnect()
}
