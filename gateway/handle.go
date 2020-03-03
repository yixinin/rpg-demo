package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"rpg-demo/lib/log"
	"rpg-demo/protocol"
	"rpg-demo/utils"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/gin-gonic/gin"
)

func (g *Gateway) HandleCallback(ev cellnet.Event) {

	// 处理Peer收到的各种事件
	var sess = ev.Session()
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted: // 接受一个连接
		fmt.Println("server accepted", msg.String())

	case *cellnet.SessionClosed: // 会话连接断开
		fmt.Println("session closed: ", ev.Session().ID())

	}
	g.handleMessage(ev.Message(), sess, sess.Send)
}

func (g *Gateway) handleHttp(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	var msgid = utils.BytesToUint16(data[:2])
	msg, _, err := codec.DecodeMessage(int(msgid), data[2:])
	if err != nil {
		return
	}
	var send = func(m interface{}) {
		data, _, err := codec.EncodeMessage(m, nil)
		if err != nil {
			return
		}
		c.Data(http.StatusOK, "data/bytes", data)
	}
	g.handleMessage(msg, nil, send)
}

func (g *Gateway) handleMessage(msgIntface interface{}, sess cellnet.Session, send func(m interface{})) error {
	switch msg := msgIntface.(type) {
	case *protocol.LoginReq: // 收到连接发送的消息
		var rpcConn = g.grpcClient.RandCenter()
		if rpcConn == nil {
			log.Error("没有运行的center服务")
			return nil
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
		send(ack)
	case *protocol.LogoutReq:
		ok, err := g.server.CheckUser(msg.Header)
		if !ok || err != nil {
			return err
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
			return err
		}

		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.Logout(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		send(ack)
	case *protocol.GetGameRoomTypeListReq:
		ok, err := g.server.CheckUser(msg.Header)
		if !ok || err != nil {
			return err
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
			return err
		}

		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.GetGameRoomTypeList(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		send(ack)
	default:

		log.Warn("unkown message", msg)
	}
	return nil
}
