package ddz

import (
	"com/lib/log"
	"com/utils"
	"reflect"

	"github.com/davyxu/cellnet/codec"
	nats "github.com/nats-io/nats.go"
)

func (g *DDZGame) handleMessageQueue(m *nats.Msg) {
	msgIntface, _, err := codec.DecodeMessage(int(utils.StringHash(m.Subject)), m.Data)
	if err != nil {
		log.Error(err)
		return
	}
	switch msg := msgIntface.(type) {
	default:
		log.Warn("unhandled message: %s", reflect.TypeOf(msg).Elem().Name())
	}
}
