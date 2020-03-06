package rpc

import (
	"center/logic"
	"center/server"
	"com/protocol"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcService(address string, center *server.Center) {
	var addr = fmt.Sprintf("0.0.0.0:%s", address)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protocol.RegisterCenter4GateServiceServer(s, &logic.Center4Gate{Center: center})
	s.Serve(lis)
	if err != nil {
		log.Fatal("listen grpc service error", err)
		return
	}
	log.Println("rpc服务已经开启")
}
