package main

import (
	"flag"
	"os"
	"rpg-demo/config"
	"rpg-demo/db"
	"rpg-demo/lib/hook"
	"rpg-demo/lib/log"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "conf", "C:\\Users\\yixin\\go\\rpg-demo\\config\\gateway.yaml", "-conf app/conf")
	flag.Parse()

	conf, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	db.InitRedis(conf.Redis)

	gateway := NewGateway(conf)
	gateway.StartGateway("tcp")

	var app = hook.NewApp()
	app.InstallShutdownHook(gateway.grpc)
	app.InstallShutdownHook(gateway.etcd)
	app.InstallShutdownHook(gateway.nats)
	app.Start()
}
