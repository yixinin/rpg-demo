package gateway

import (
	"context"
	"fmt"
	"rpg-demo/lib/log"
	"rpg-demo/protocol"

	"github.com/davyxu/cellnet"
)

func (g *Gateway) HandleCallback(ev cellnet.Event) {

	// 处理Peer收到的各种事件
	var sess = ev.Session()
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted: // 接受一个连接
		fmt.Println("server accepted")

	case *protocol.LoginReq: // 收到连接发送的消息
		var rpcConn = g.grpcClient.RandCenter()
		if rpcConn == nil {
			log.Error("没有运行的center服务")
			return
		}
		var client = protocol.NewCenter4GateServiceClient(rpcConn.Conn)
		ack, err := client.Login(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		if ack.Header.Code == 0 { //登录成功
			c, ok := g.server.Get(ack.Header.Uid)
			if ok { //踢下线
				if c.sess.ID() != sess.ID() {
					c.sess.Close()
					//center 下线
					//game 下线
					g.server.Del(ack.Header.Uid)
				}
				g.server.Replace(&Connection{
					Uid:      ack.Header.Uid,
					Token:    ack.Token,
					sess:     sess,
					CenterIP: rpcConn.IP,
				})
			} else {
				g.server.Replace(&Connection{
					Uid:      ack.Header.Uid,
					Token:    ack.Token,
					sess:     sess,
					CenterIP: rpcConn.IP,
				})
			}

		}
		sess.Send(ack)
	case *protocol.LogoutReq:
		ok, err := g.server.CheckUser(msg.Header)
		if !ok || err != nil {
			return
		}
		conn, err := g.server.GetByToken(msg.Header.Token)
		if err != nil {
			log.Error("check token error:%v", err)
		}
		if conn == nil || conn.CenterIP == "" {
			log.Error("check token error:%v", err)
		}

		rpcConn, ok := g.grpcClient.centers[conn.CenterIP]
		if !ok {
			return
		}

		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.Logout(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		sess.Send(ack)

	case *cellnet.SessionClosed: // 会话连接断开
		fmt.Println("session closed: ", ev.Session().ID())
	case *protocol.GetGameRoomTypeListReq:
		ok, err := g.server.CheckUser(msg.Header)
		if !ok || err != nil {
			return
		}
		conn, err := g.server.GetByToken(msg.Header.Token)
		if err != nil {
			log.Error("check token error:%v", err)
		}
		if conn == nil || conn.CenterIP == "" {
			log.Error("check token error:%v", err)
		}

		rpcConn, ok := g.grpcClient.centers[conn.CenterIP]
		if !ok {
			return
		}

		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.GetGameRoomTypeList(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		sess.Send(ack)
	default:
		ev.Session().Send(&protocol.LoginAck{
			Header: &protocol.AckHeader{
				Code: 1,
				Msg:  "success",
				Mid:  10,
			},
			Token: "sadsfasf",
		})
		fmt.Println("unkown message", msg)
	}
}
