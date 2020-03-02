package client

import (
	"rpg-demo/lib/log"
	"rpg-demo/protocol"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"

	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

func Client() {
	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件
	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Connector", "client", "10.0.0.23:8001", queue)
	p.(cellnet.WSConnector).SetReconnectDuration(time.Second)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {

		case *cellnet.SessionConnected:
			log.Debug("server connected")

			ev.Session().Send(&protocol.LoginReq{
				Header:   &protocol.ReqHeader{},
				UserName: "yixin",
				Password: "asdasd",
			})
			// 有连接断开
		case *cellnet.SessionClosed:
			log.Debug("session closed: ", ev.Session().ID())
		case *protocol.LoginAck:

			log.Debug("recv: %+v %v", msg, []byte("鲍勃"))

		}
	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()
}
