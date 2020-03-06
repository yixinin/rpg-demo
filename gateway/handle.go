package main

import (
	"com/lib/log"
	"com/protocol"
	"com/utils"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/gin-gonic/gin"
	nats "github.com/nats-io/nats.go"
)

func (g *Gateway) HandleCallback(ev cellnet.Event) {

	// 处理Peer收到的各种事件
	var sess = ev.Session()
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted: // 接受一个连接
		fmt.Println("server accepted", msg.String())

	case *cellnet.SessionClosed: // 会话连接断开
		fmt.Println("session closed: ", sess.ID())
		g.server.DelSess(sess.ID())
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

func (g *Gateway) handleMessageQueue(m *nats.Msg) {
	msgIntface, _, err := codec.DecodeMessage(int(utils.StringHash(m.Subject)), m.Data)
	if err != nil {
		log.Error(err)
		return
	}

	g.handleMessage(msgIntface, nil, func(m interface{}) {})
}

func (g *Gateway) handleMessage(msgIntface interface{}, sess cellnet.Session, send func(m interface{})) error {
	switch msg := msgIntface.(type) {
	case *protocol.LoginReq: // 收到连接发送的消息
		var rpcConn = g.grpc.GetCenter()
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
			if ok && c != nil { //踢下线
				if sess == nil || c.sess.ID() != sess.ID() {
					//center 下线
					//game 下线
					g.server.Del(ack.Header.Uid)
				}

			}

			g.server.Add(&Connection{
				Uid:      ack.Header.Uid,
				Token:    ack.Token,
				sess:     sess,
				CenterIP: rpcConn.IP,
			})

		}
		send(ack)
	case *protocol.LogoutReq:
		rpcConn, err := g.GetRpcConn(msg.Header, "center")
		if err != nil {
			var ack = &protocol.GetGameRoomTypeListAck{
				Header: &protocol.AckHeader{},
			}
			if err == ErrUserNotLogin {
				ack.Header.Code = 401
				ack.Header.Msg = "user not login"
			}
			send(ack)
		}

		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.Logout(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		send(ack)
		//删除登录信息
		g.server.Del(msg.Header.GetUid())
	case *protocol.GetGameRoomTypeListReq:
		rpcConn, err := g.GetRpcConn(msg.Header, "center")
		if err != nil {
			var ack = &protocol.GetGameRoomTypeListAck{
				Header: &protocol.AckHeader{},
			}
			if err == ErrUserNotLogin {
				ack.Header.Code = 401
				ack.Header.Msg = "user not login"
			}
			send(ack)
		}
		var client = protocol.NewCenter4GateServiceClient(rpcConn)
		ack, err := client.GetGameRoomTypeList(context.Background(), msg)
		if err != nil {
			log.Error("rpc error:%v", err)
		}
		send(ack)
	default:
		log.Warn("unhandled message: %s", reflect.TypeOf(msg).Elem().Name())
	}
	return nil
}
