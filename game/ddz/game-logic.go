package ddz

import (
	"com/protocol"
	"context"
)

type GameLogic struct {
	game *DDZGame
}

func (d *GameLogic) KickUser(ctx context.Context, req *protocol.KickUserReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}

func (g *GameLogic) EnterGame(ctx context.Context, req *protocol.EnterGameReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}

func (g *GameLogic) UpdateUserInfo(ctx context.Context, req *protocol.UpdateUserInfoReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}

func (g *GameLogic) LeaveGame(ctx context.Context, req *protocol.LeaveGameReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}

func (d *GameLogic) CloseGame(ctx context.Context, req *protocol.CloseGameReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *GameLogic) AddCoin(ctx context.Context, req *protocol.AddCoinReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
func (d *GameLogic) AddPlayer(ctx context.Context, req *protocol.AddPlayerReq) (*protocol.CallAck, error) {
	ack := &protocol.CallAck{}
	return ack, nil
}
