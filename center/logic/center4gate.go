package logic

import (
	"center/iface"
	"center/player"
	"center/room"
	"center/server"
	"com/cache"
	"com/db"
	"com/lib/log"
	"com/models"
	"com/protocol"
	"com/utils"
	"context"
	"errors"
	"time"

	"github.com/go-xorm/xorm"
)

type Center4Gate struct {
	Center *server.Center
}

func (u *Center4Gate) CreateAccount(user *models.User) error {
	var f = func(s *xorm.Session) (interface{}, error) {
		user.RegTime = time.Now()

		userId, err := s.Insert(user)
		if err != nil {
			return nil, err
		}
		if userId == 0 {
			return nil, errors.New("创建用户失败，id=0")
		}
		user.Id = userId
		var account = &models.Account{
			UserId: userId,
		}
		_, err = s.Insert(account)
		return nil, err
	}
	_, err := db.Mysql.Transaction(f)
	return err
}

func (u *Center4Gate) Login(ctx context.Context, req *protocol.LoginReq) (ack *protocol.LoginAck, err error) {
	// 登录方式 1；用户名密码；2：手机；3：微信；4：qq；5：设备码；10：内部token
	ack = &protocol.LoginAck{
		Header: &protocol.AckHeader{},
	}
	var user = &models.User{}
	switch req.LoginType {
	case 1:
		user.UserName = req.UserName
		var ok bool
		ok, err = db.Mysql.Get(user)
		if err != nil {
			return
		}
		if !ok {
			if !req.Register {
				ack.Header.Code = 1
				ack.Header.Msg = "user not register"
				return
			}
			//创建用户
			user.Password = utils.MD5(req.Password)
			user.DeviceCode = req.DeviceCode
			err = u.CreateAccount(user)
			if err != nil {
				return
			}
		}
		if user.Password != utils.MD5(req.Password) {
			ack.Header.Code = 1
			ack.Header.Msg = "username and password not match"
			return
		}
	default:
		ack.Header.Code = 401
		ack.Header.Msg = "AccessDenied"
		return
	}
	//登录成功
	//更新最后登录时间
	user.LastLoginTime = time.Now()
	n, err := db.Mysql.Id(user.Id).Cols("last_login_time").Update(user)
	if err != nil || n != 1 {
		log.Error("更新最近登录时间错误，n=%d,err=%v", n, err)
	}

	log.Debug("userid:%d,n:%d,err:%v", user.Id, n, err)

	//生成token
	token := utils.UUID()
	ack.Token = token
	ack.Header.Uid = user.Id

	//缓存
	err = cache.SetToken(user.Id, token)
	if err != nil {
		log.Error("set token error:%v", err)
	}
	err = cache.SetUserService(user.Id)
	if err != nil {
		log.Error("set user service error:%v", err)
	}

	log.Info("login success! token:%s", token)
	//添加到players
	//获取用户资产
	var account = new(models.Account)
	u.Center.AddPlayer(&player.User{
		UserId:   user.Id,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		IsFamale: user.IsFamale,
		Coin:     account.Coin,
	})
	return
}

func (u *Center4Gate) Logout(ctx context.Context, req *protocol.LogoutReq) (*protocol.LogoutAck, error) {
	u.Center.RemovePlayer(req.Header.Uid)
	return &protocol.LogoutAck{}, nil
}

//重连
func (u *Center4Gate) ReConnect(ctx context.Context, req *protocol.ReConnecReq) (*protocol.CallAck, error) {
	return nil, nil
}

//踢下线
func (u *Center4Gate) KickUser(ctx context.Context, req *protocol.KickUserReq) (*protocol.CallAck, error) {
	return nil, nil
}

//------------------ 上分/下分 -----------------------

func (u *Center4Gate) Topup(ctx context.Context) {

}

func (u *Center4Gate) Withdraw(ctx context.Context) {

}

//------------------ Room 房间 -----------------------
func (u *Center4Gate) CreateRoom(context.Context, *protocol.CreateRoomReq) (*protocol.CreateRoomAck, error) {
	return nil, nil
}
func (u *Center4Gate) JoinRoom(context.Context, *protocol.JoinRoomReq) (*protocol.JoinRoomAck, error) {
	return nil, nil
}
func (u *Center4Gate) ExitRoom(context.Context, *protocol.ExitRoomReq) (*protocol.ExitRoomAck, error) {
	return nil, nil
}
func (u *Center4Gate) DiscardRoom(context.Context, *protocol.DiscardRoomReq) (*protocol.DiscardRoomAck, error) {
	return nil, nil
}

//------------------ Game 游戏 ------------------------
func (u *Center4Gate) GetGameRoomTypeList(ctx context.Context, req *protocol.GetGameRoomTypeListReq) (*protocol.GetGameRoomTypeListAck, error) {
	var ack = &protocol.GetGameRoomTypeListAck{
		Header: &protocol.AckHeader{},
	}
	//暂时从数据库查
	var gameTypes = make([]*models.GameType, 0, 5)
	err := db.Mysql.Where("game_id=?", req.GameId).Find(&gameTypes)
	log.Debug("err:%v, len:%d", err, len(gameTypes))
	if err != nil {
		return ack, err
	}
	var items = make([]*protocol.GetGameRoomTypeListAck_GameTypeItem, 0, len(gameTypes))
	for _, v := range gameTypes {
		items = append(items, &protocol.GetGameRoomTypeListAck_GameTypeItem{
			GameType: v.GameType,
			GameInfo: v.FullName,
		})
	}
	ack.Items = items
	return ack, err
}
func (u *Center4Gate) EnterGame(ctx context.Context, req *protocol.EnterGameReq) (*protocol.EnterGameRep, error) {
	var ack = &protocol.EnterGameRep{}
	var r iface.Room
	switch req.GameId {
	case room.DDZGameId:
		//查找有无可用房间
		if v1, ok := u.Center.Games[req.GameId]; ok {
			if v2, ok := v1[req.GameType]; ok {
				for _, v := range v2 {
					if !v.GetIsFull() {
						r = v //可用的空房间
						break
					}
				}
			}
		}
		r = room.CreateDDZRoom(req.Header.Uid, req.GameType, "")
		u.Center.CreateRoom(r)
	}
	if r == nil {
		return ack, errors.New("enter game fail")
	}
	err := r.JoinRoom(req.Header.Uid)
	if err != nil {
		return ack, err
	}
	u.Center.Players[req.Header.Uid].SetRoom(r.GetRoomId())
	return ack, nil
}
func (u *Center4Gate) ExitGame(context.Context, *protocol.ExitGameReq) (*protocol.ExitGameAck, error) {
	return nil, nil
}

//------------------- Chat 聊天系统 -------------------------
func (u *Center4Gate) SendToUser(context.Context, *protocol.SendToUserReq) (*protocol.CallAck, error) {
	return nil, nil
}
func (u *Center4Gate) SendToRoom(context.Context, *protocol.SendToRoomReq) (*protocol.CallAck, error) {
	return nil, nil
}
func (u *Center4Gate) Boardcast(context.Context, *protocol.BoardcastReq) (*protocol.CallAck, error) {
	return nil, nil
}
