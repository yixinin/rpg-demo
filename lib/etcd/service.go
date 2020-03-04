package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"rpg-demo/lib/log"
	"rpg-demo/utils"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//the detail of service
type ServiceInfo struct {
	IP string
}

type Service struct {
	Name    string
	Info    ServiceInfo
	stop    chan error
	leaseid clientv3.LeaseID
	client  *clientv3.Client
}

func NewService(name string, info ServiceInfo, endpoints []string) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		log.Error("err:%v", err)
		return nil, err
	}

	return &Service{
		Name:   name,
		Info:   info,
		stop:   make(chan error),
		client: cli,
	}, err
}

func (s *Service) Start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Error("err:%v", err)
		return err
	}

	for {
		select {
		case err := <-s.stop:
			s.revoke()
			return err
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case _, ok := <-ch:
			if !ok {
				log.Error("keep alive channel closed")
				s.revoke()
				return nil
			} else {
				// log.Info("Recv reply from service: %s, ttl:%d", s.Name, ka.TTL)
			}
		}
	}
}

func (s *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {

	info := &s.Info

	key := fmt.Sprintf("rpg/%s/%s", s.Name, utils.RandStr(10))
	value, _ := json.Marshal(info)

	// minimum lease TTL is 5-second
	resp, err := s.client.Grant(context.TODO(), 10)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	s.leaseid = resp.ID

	return s.client.KeepAlive(context.TODO(), resp.ID)
}

func (s *Service) revoke() error {

	_, err := s.client.Revoke(context.TODO(), s.leaseid)
	if err != nil {
		log.Error("err", err)
	}
	log.Info("servide:%s stop\n", s.Name)
	return err
}

func (s *Service) Shutdown() {
	s.stop <- nil
}
