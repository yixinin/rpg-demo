package server

import (
	"com/lib/etcd"
	"com/lib/log"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type GrpcPool struct {
	Centers  map[string]*grpc.ClientConn
	Gateways map[string]*grpc.ClientConn
	Etcd     []string
	stop     chan bool
}

type GrpcInfo struct {
	IP string

	Conn *grpc.ClientConn
}

func NewGrpcPool(etcd []string) *GrpcPool {
	return &GrpcPool{
		Centers:  make(map[string]*grpc.ClientConn, 3),
		Gateways: make(map[string]*grpc.ClientConn, 3),
		Etcd:     etcd,
	}
}

func (s *GrpcPool) GetCenter(ip ...string) *GrpcInfo {
	if len(ip) == 1 {
		return &GrpcInfo{
			IP:   ip[0],
			Conn: s.Centers[ip[0]],
		}
	}
	for k, v := range s.Centers {
		return &GrpcInfo{
			IP:   k,
			Conn: v,
		}
	}
	return nil
}

func (s *GrpcPool) Center(ip string) *grpc.ClientConn {
	v, _ := s.Centers[ip]
	return v
}
func (s *GrpcPool) Gateway(ip string) *grpc.ClientConn {
	v, _ := s.Gateways[ip]
	return v
}

//ListenGrpc go s.ListenGrpc()
func (s *GrpcPool) ListenGrpc() {
	m, err := etcd.NewMaster(s.Etcd, "rpg/")

	if err != nil {
		log.Error("err:%v", err)
		return
	}

	var listen = func() {
		var ips = make([]string, 0, len(m.Nodes))
		for k, v := range m.Nodes {
			var ks = strings.Split(k, "/")
			if len(ks) < 2 {
				continue
			}
			ips = append(ips, v.Info.IP)
			switch ks[1] {
			case "center":
				if _, ok := s.Centers[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.Centers[v.Info.IP] = conn
				}
			case "gateway":
				if _, ok := s.Gateways[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.Gateways[v.Info.IP] = conn
				}
			}
		}
		s.RemoveOfflineNode(ips...)
	}

	for {
		select {
		case <-s.stop:
			return
		case <-time.After(5 * time.Second):
			listen()
		}
	}
}

func (s *GrpcPool) RemoveOfflineNode(ips ...string) {
	for k := range s.Centers {
		var use = false
		for _, ip := range ips {
			if k == ip {
				use = true
				break
			}
		}
		// 断开连接
		if !use {
			s.Centers[k].Close()
			delete(s.Centers, k)
		}
	}
	for k := range s.Gateways {
		for _, ip := range ips {
			if k == ip {
				continue
			}
		}
		// 断开连接
		s.Centers[k].Close()
		delete(s.Gateways, k)
	}
}

func (s *GrpcPool) Shutdown() {
	s.stop <- true
}
