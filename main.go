package main

import (
	"flag"
	"fmt"
	"rpg-demo/center"
	"rpg-demo/client"
	"rpg-demo/config"
	"rpg-demo/db"
	"rpg-demo/gateway"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/hook"
	"rpg-demo/lib/log"
	"rpg-demo/utils"
)

var appGateway bool
var appClient bool
var appCenter bool
var appGame bool
var configPath string

var appMaster bool //监控程序

func main() {

	flag.BoolVar(&appGateway, "gateway", false, "-gateway")
	flag.BoolVar(&appCenter, "center", false, "-center")
	flag.BoolVar(&appGame, "game", false, "-game")

	flag.StringVar(&configPath, "conf", "C:\\Users\\yixin\\go\\rpg-demo\\config\\gateway.yaml", "-conf app/conf")

	flag.BoolVar(&appClient, "client", false, "-client")
	flag.BoolVar(&appMaster, "master", false, "-master")

	flag.Parse()

	config.Init(configPath)
	db.InitMysql()
	db.InitRedis()
	var serviceName string
	var serviceInfo = etcd.ServiceInfo{
		IP: fmt.Sprintf("%s:%s", utils.IPAddress(), config.DefaultConfig.Grpc),
	}

	log.Info("grpc listen in %d", serviceInfo.IP)

	if appGateway {
		var gateway = gateway.NewGateway()
		serviceName = "gateway"
		go gateway.StartServer()
	} else if appClient {
		client.Client()
		serviceName = "client"
	} else if appCenter {
		serviceName = "center"
		go center.StartCenter()
	} else if appGame {
		serviceName = "game"
	} else {
		fmt.Println("unknow app")
		return
	}

	service, err := etcd.NewService(serviceName, serviceInfo, config.DefaultConfig.EtcdAddress)
	if err != nil {
		log.Error("etcd conn error:%v", err)
	}
	service.Start()

	var app = hook.NewApp()
	app.Start()
}
