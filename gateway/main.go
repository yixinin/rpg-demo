package main

import (
	"flag"
	"fmt"
	"rpg-demo/config"
	"rpg-demo/db"
	"rpg-demo/lib/bus"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/hook"
	"rpg-demo/lib/log"
	"rpg-demo/utils"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "conf", "C:\\Users\\yixin\\go\\rpg-demo\\config\\gateway.yaml", "-conf app/conf")
	flag.Parse()

	config.Init(configPath)
	db.InitRedis()

	n, err := bus.NewBus(config.DefaultConfig.Nats)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(n)
	var serviceInfo = etcd.ServiceInfo{
		IP: fmt.Sprintf("%s:%s", utils.IPAddress(), config.DefaultConfig.Grpc),
	}

	log.Info("grpc listen in %d", serviceInfo.IP)

	gateway := NewGateway()
	gateway.StartGateway("tcp")

	service, err := etcd.NewService("gateway", serviceInfo, config.DefaultConfig.EtcdAddress)
	if err != nil {
		log.Error("etcd conn error:%v", err)
	}
	service.Start()

	var app = hook.NewApp()
	app.Start()
}
