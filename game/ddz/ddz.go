package ddz

import (
	"com/ip"
	"com/protocol"
	"fmt"
	"game/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartDDZ(game *server.Game) {
	var ddzGame = NewDDZGame(game)
	ddzGame.StartGame()

}

type DDZGame struct {
	game      *server.Game
	rooms     map[int64]*Room
	player    map[int64]Player
	team      map[int64]*Team
	typeRooms map[string][]int64 //房间种类分列表
	roomTypes []string           //房间种类
}

func NewDDZGame(game *server.Game) *DDZGame {
	return &DDZGame{
		game: game,
	}
}

func (g *DDZGame) StartGame() {
	g.game.SubMessageQueue(g.handleMessageQueue)
	g.game.StartGame()
	StartGrpcService(ip.ServerIP, g)
}

func StartGrpcService(address string, game *DDZGame) {
	var addr = fmt.Sprintf("0.0.0.0:%s", address)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protocol.RegisterGameServiceServer(s, &GameLogic{game: game})
	protocol.RegisterGameDdzServiceServer(s, &DdzLogic{game: game})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("listen grpc service error", err)
		return
	}
	log.Println("rpc服务已经开启")
}
