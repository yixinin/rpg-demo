package bus

import (
	"strings"
	"time"

	"rpg-demo/lib/log"

	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/nats.go"
)

type Config struct {
	Addr    string
	Cluster string
}

type Bus struct {
	nats *nats.Conn
}

type Mesg struct {
	nats *nats.Msg
	pm   interface{}
}

type Subscription struct {
	nats *nats.Subscription
}

func (m *Mesg) PbMesg() interface{} {
	return m.pm
}

type MesgHandler func(*Mesg)
type NewProtoMesg func() interface{}

func NewBus(cfg *Config) (*Bus, error) {
	conn, err := nats.Connect(cfg.Addr,
		nats.MaxReconnects(-1), // always reconnect
		nats.DisconnectHandler(disconnectHandler),
		nats.ReconnectHandler(reconnectHandler),
		nats.ClosedHandler(closedHandler))
	if err != nil {
		return nil, err
	}

	return &Bus{
		nats: conn,
	}, nil
}

// handle disconnect event
func disconnectHandler(nc *nats.Conn) {
	log.Error("Disconnected from config-servers[%v]!", strings.Join(nc.Opts.Servers, "|"))
}

// handle reconnect event
func reconnectHandler(nc *nats.Conn) {
	log.Warn("Reconnect to [%v], config-servers[%v]!", nc.ConnectedUrl(), strings.Join(nc.Opts.Servers, "|"))
}

// handle close event
func closedHandler(nc *nats.Conn) {
	log.Error("Connection to config-servers[%v] is closed. Reason: %q", strings.Join(nc.Opts.Servers, "|"), nc.LastError())
}

func (b *Bus) Request(topic string, req proto.Message, resp interface{}) error {
	if data, err := proto.Marshal(req); err != nil {
		return err
	} else if respMesg, err := b.nats.Request(topic, data, 2*time.Second); err != nil {
		return err
	} else if err := proto.Unmarshal(respMesg.Data, resp.(proto.Message)); err != nil {
		return err
	}

	return nil
}

// 普通订阅
// 针对某一subject，所有订阅者都收到一份消息
func (b *Bus) Subscribe(topic string, nc NewProtoMesg, cb MesgHandler) (*Subscription, error) {
	s, err := b.nats.Subscribe(topic, func(m *nats.Msg) {
		pm := nc()
		if err := proto.Unmarshal(m.Data, pm.(proto.Message)); err != nil {
			log.Error("Subscription receive data unmarshal error: %v", err)
			return
		} else {
			cb(&Mesg{nats: m, pm: pm})
		}
	})

	if err != nil {
		return nil, err
	} else {
		return &Subscription{
			nats: s,
		}, nil
	}
}

// 队列订阅
// 针对某一subject，所有订阅者只有一个收到消息
// 目前固定queue group名称为default，假设针对某一subject只会需要一个queue grup
func (b *Bus) QueueSubscribe(topic string, nc NewProtoMesg, cb MesgHandler) (*Subscription, error) {
	s, err := b.nats.QueueSubscribe(topic, "default", func(m *nats.Msg) {
		pm := nc()
		if err := proto.Unmarshal(m.Data, pm.(proto.Message)); err != nil {
			log.Error("Subscription receive data unmarshal error: %v", err)
			return
		} else {
			cb(&Mesg{nats: m, pm: pm})
		}
	})

	if err != nil {
		return nil, err
	} else {
		return &Subscription{
			nats: s,
		}, nil
	}
}

func (b *Bus) Publish(topic string, msg interface{}) error {
	if data, err := proto.Marshal(msg.(proto.Message)); err != nil {
		return err
	} else {
		return b.nats.Publish(topic, data)
	}

}

func (b *Bus) Reply(m interface{}, to *Mesg) error {
	if data, err := proto.Marshal(m.(proto.Message)); err != nil {
		return err
	} else {
		return b.nats.Publish(to.nats.Reply, data)
	}
}

func (b *Bus) GetNats() *nats.Conn {
	return b.nats
}
