package server

import (
	"com/lib/etcd"
	"com/lib/log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func Watch() {

}

type GrpcPool struct {
	gateways map[string]*grpc.ClientConn
	games    map[string]*grpc.ClientConn
	game     map[int32][]string //gameid->servers
	etcd     []string
	stop     chan bool
}

type GrpcInfo struct {
	IP   string
	Conn *grpc.ClientConn
}

func NewGrpcPool(etcd []string) *GrpcPool {
	return &GrpcPool{
		gateways: make(map[string]*grpc.ClientConn, 3),
		games:    make(map[string]*grpc.ClientConn, 3),
		etcd:     etcd,
	}
}

func (s *GrpcPool) GetGateway(ip ...string) *GrpcInfo {
	if len(ip) == 1 {
		return &GrpcInfo{
			IP:   ip[0],
			Conn: s.gateways[ip[0]],
		}
	}
	for k, v := range s.gateways {
		return &GrpcInfo{
			IP:   k,
			Conn: v,
		}
	}
	return nil
}

func (s *GrpcPool) Gateway(ip string) *grpc.ClientConn {
	v, _ := s.gateways[ip]
	return v
}
func (s *GrpcPool) Game(ip string) *grpc.ClientConn {
	v, _ := s.games[ip]
	return v
}

//ListenGrpc go s.ListenGrpc()
func (s *GrpcPool) ListenGrpc() {
	m, err := etcd.NewMaster(s.etcd, "rpg/")

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
			case "gateway":
				if _, ok := s.gateways[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.gateways[v.Info.IP] = conn
				}
			case "game":

				if _, ok := s.games[v.Info.IP]; !ok {
					conn, err := grpc.Dial(v.Info.IP, grpc.WithInsecure())
					if err != nil {
						log.Error("err:%v", err)
					}
					s.games[v.Info.IP] = conn
					if len(ks) >= 3 {
						gameid, err := strconv.ParseInt(ks[2], 10, 32)
						if err == nil {
							if _, ok := s.game[int32(gameid)]; !ok {
								s.game[int32(gameid)] = []string{v.Info.IP}
							} else {
								s.game[int32(gameid)] = append(s.game[int32(gameid)], v.Info.IP)
							}
						}
					}
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
	for k := range s.gateways {
		var use = false
		for _, ip := range ips {
			if k == ip {
				use = true
				break
			}
		}
		// 断开连接
		if !use {
			s.gateways[k].Close()
			delete(s.gateways, k)
		}
	}
	for k := range s.games {
		for _, ip := range ips {
			if k == ip {
				continue
			}
		}
		// 断开连接
		s.gateways[k].Close()
		delete(s.games, k)
		//删除游戏节点
		for k1, v1 := range s.game {
			keeps := make([]string, 0, len(v1))
			for _, node := range v1 {
				if k != node {
					keeps = append(keeps, node)
				}
			}
			s.game[k1] = keeps
		}
	}

}

func (s *GrpcPool) Shutdown() {
	s.stop <- true
}
