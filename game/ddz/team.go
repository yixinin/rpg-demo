package ddz

import "sync"

type Team struct {
	sync.Mutex
	teamId  int64
	players []int64
	leader  int64 //小队队长
	invite  bool  //所有人可以邀请
}

func CreateTeam(id int64, playerId int64) *Team {
	t := &Team{
		teamId:  id,
		players: make([]int64, 0, 2),
	}
	t.players = append(t.players, playerId)
	t.leader = playerId
	return t
}

func (t *Team) Add(playerId ...int64) {
	t.Lock()
	defer t.Unlock()
	t.players = append(t.players, playerId...)
}
func (t *Team) Remove(playerId int64) {
	t.Lock()
	defer t.Unlock()
	var index int
	for i, v := range t.players {
		if v == playerId {
			index = i
		}
	}
	t.players = append(t.players[:index], t.players[index+1:]...)
}

//移交队长
func (t *Team) SetLeader(playerId int64) bool {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.players {
		if v == playerId {
			t.leader = v
			return true
		}
	}
	return false
}

//设置邀请权限
func (t *Team) SetInvite(b bool) {
	t.Lock()
	defer t.Unlock()
	t.invite = b
}
