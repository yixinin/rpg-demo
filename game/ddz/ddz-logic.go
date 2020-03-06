package ddz

import (
	"com/protocol"
	"context"
)

type DdzLogic struct {
	game *DDZGame
}

func (d *DdzLogic) JoinTeam(ctx context.Context, req *protocol.JoinTeamReq) (*protocol.CallAck, error) {
	var ack = &protocol.CallAck{}
	if req.TeamId != 0 {
		if team, ok := d.game.team[req.TeamId]; ok {
			team.Add(req.Uid)
		}
	} else { //创建小队
		var t = CreateTeam(d.game.game.TeamIdPool.GetOne(), req.Uid)
		d.game.team[t.teamId] = t
	}

	return ack, nil
}

//房间满员时开始游戏
func (d *DdzLogic) JoinRoom(ctx context.Context, req *protocol.JoinRoomReq) (*protocol.CallAck, error) {
	var ack = &protocol.CallAck{}
	roomIds, ok := d.game.typeRooms[req.GameType]
	if !ok {
		ack.Code = 400
		ack.Msg = "房间类型错误"
		return ack, nil
	}
	for _, roomId := range roomIds {
		room, ok := d.game.rooms[roomId]
		if ok {
			switch room.maxGroup {
			case 0:
				if req.TeamId != 0 { //非组队房间
					ack.Code = 400
					ack.Msg = "不可组队"
					return ack, nil
				}
			case 1:
				team, ok := d.game.team[req.TeamId]
				if ok {
					if len(team.players) != room.maxPlayer { //人数不足
						ack.Code = 400
						ack.Msg = "玩家不足"
						return ack, nil
					}
				}
			default:

			}

		}
	}
	return nil, nil
}

//明牌
func (d *DdzLogic) CallVisible(ctx context.Context, req *protocol.CallVisibleReq) (*protocol.CallVisibleAck, error) {
	return nil, nil
}

//叫地主
func (d *DdzLogic) CallBanker(ctx context.Context, req *protocol.CallBankerReq) (*protocol.CallBankerAck, error) {
	return nil, nil
}

//加倍
func (d *DdzLogic) CallDuoble(ctx context.Context, req *protocol.CallDuobleReq) (*protocol.CallDuobleAck, error) {
	return nil, nil
}

//出牌
func (d *DdzLogic) DiscardPoker(ctx context.Context, req *protocol.DiscardPokerReq) (*protocol.DiscardPokerAck, error) {
	return nil, nil
}
