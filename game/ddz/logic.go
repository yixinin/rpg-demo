package ddz

import (
	"com/protocol"
	"context"
)

type DDZLogic struct {
	game *DDZGame
}

func (d *DDZLogic) KickUser(ctx context.Context, req *protocol.KickUserReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *DDZLogic) StartGame(ctx context.Context, req *protocol.StartGameReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *DDZLogic) CloseGame(ctx context.Context, req *protocol.CloseGameReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *DDZLogic) AddCoin(ctx context.Context, req *protocol.AddCoinReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *DDZLogic) AddPlayer(ctx context.Context, req *protocol.AddPlayerReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
