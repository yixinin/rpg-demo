package game

import (
	"flag"
	"os"
	"rpg-demo/config"
	"rpg-demo/db"
	"rpg-demo/game/server"
	"rpg-demo/lib/hook"
	"rpg-demo/lib/log"
)

var gameName string

func main() {
	flag.StringVar(&gameName, "game", "ddz", "-game=ddz")
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
	game.StartGame()

	var app = hook.NewApp()
	app.InstallShutdownHook(game.grpc)
	app.InstallShutdownHook(game.etcd)
	app.InstallShutdownHook(game.nats)
	app.Start()
}
