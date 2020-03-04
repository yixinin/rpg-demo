package center

import (
	"flag"
	"os"
	"rpg-demo/center/rpc"
	"rpg-demo/center/server"
	"rpg-demo/config"
	"rpg-demo/db"
	"rpg-demo/lib/hook"
	"rpg-demo/lib/log"
)

var configPath string

func StartCenter() {
	flag.StringVar(&configPath, "conf", "C:\\Users\\yixin\\go\\rpg-demo\\config\\gateway.yaml", "-conf app/conf")
	flag.Parse()

	conf, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	db.InitRedis(conf.Redis)
	db.InitMysql(conf.Mysql)

	var center = server.NewCenter(conf)
	rpc.StartGrpcService(conf.Grpc, center)
	var app = hook.NewApp()
	app.InstallShutdownHook(center.Grpc)
	app.InstallShutdownHook(center.Etcd)
	app.InstallShutdownHook(center.Nats)
	app.Start()
}
