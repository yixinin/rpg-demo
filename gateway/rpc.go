package gateway

import (
	"rpg-demo/config"
	"rpg-demo/lib/etcd"
	"rpg-demo/lib/log"
	"time"

	"google.golang.org/grpc"
)

var (
	centers []string
	games   []string
)

func Watch() {

}

type GrpcClient struct {
	centers map[string]*grpc.ClientConn
	games   map[string]*grpc.ClientConn
}

type GrpcConn struct {
	IP   string
	Conn *grpc.ClientConn
}

func NewGrpcClient() *GrpcClient {
	return &GrpcClient{
		centers: make(map[string]*grpc.ClientConn, 3),
		games:   make(map[string]*grpc.ClientConn, 3),
	}
}

// func （s *GrpcClient)RandCenter() *grpc.ClientConn{
// 	for _,v:=range s.centers{
// 		return v
// 	}
// 	return nil
// }

func (s *GrpcClient) RandCenter() *GrpcConn {
	for k, v := range s.centers {
		return &GrpcConn{
			IP:   k,
			Conn: v,
		}
	}
	return nil
}

//GrpcConnect go s.GrpcConnect()
func (s *GrpcClient) GrpcConnect() {
	m, err := etcd.NewMaster(config.DefaultConfig.EtcdAddress, "services/")

	if err != nil {
		log.Error("err:%v", err)
	}

	for {
		var ips = make([]string, 0, len(m.Nodes))
	FOR:
		for k, v := range m.Nodes {
			ips = append(ips, v.Info.IP)
			switch k {
			case "services/center":
				if _, ok := s.centers[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.centers[v.Info.IP] = conn
					continue FOR
				}
			case "services/game":
				if _, ok := s.games[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.games[v.Info.IP] = conn
					continue FOR
				}
			}
			s.RemoveUnused(ips...)
		}

		time.Sleep(5 * time.Second)
	}
}

func (s *GrpcClient) RemoveUnused(ips ...string) {
	for k := range s.centers {
		var use = false
		for _, ip := range ips {
			if k == ip {
				use = true
				break
			}
		}
		// 断开连接
		if !use {
			s.centers[k].Close()
			delete(s.centers, k)
		}
	}
	for k := range s.games {
		for _, ip := range ips {
			if k == ip {
				continue
			}
		}
		// 断开连接
		s.centers[k].Close()
		delete(s.games, k)
	}
}
