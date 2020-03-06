package game

import (
	"com/config"
	"com/db"
	"com/lib/hook"
	"com/lib/log"
	"flag"
	"game/ddz"
	"game/server"
	"os"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "conf", "C:\\Users\\yixin\\go\\rpg-demo\\config\\game.yaml", "-conf app/conf")
	flag.Parse()

	conf, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	db.InitRedis(conf.Redis)
	db.InitMysql(conf.Mysql)

	game := server.NewGame(conf)
	switch conf.GameName {
	case ddz.GameName:
		var g = ddz.NewDDZGame(game)
		g.StartGame()
	}

	var app = hook.NewApp()
	app.InstallShutdownHook(game.Grpc)
	app.InstallShutdownHook(game.Nats)
	app.InstallShutdownHook(game.Etcd)
	app.Start()
}
