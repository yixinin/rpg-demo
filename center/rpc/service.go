package rpc

import (
	"fmt"
	"log"
	"net"
	"rpg-demo/center/logic"
	"rpg-demo/config"
	"rpg-demo/protocol"

	"google.golang.org/grpc"
)

func StartGrpcService() {
	var addr = fmt.Sprintf("0.0.0.0:%s", config.DefaultConfig.Grpc)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protocol.RegisterCenter4GateServiceServer(s, &logic.Center4Gate{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}
