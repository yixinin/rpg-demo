package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"rpg-demo/lib/log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Master struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

//node is a client
type Node struct {
	State bool
	Key   string
	Info  ServiceInfo
}

func NewMaster(endpoints []string, watchPath string) (*Master, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	master := &Master{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	go master.WatchNodes()
	return master, err
}

func (m *Master) AddNode(key string, info *ServiceInfo) {
	node := &Node{
		State: true,
		Key:   key,
		Info:  *info,
	}

	m.Nodes[node.Key] = node
}

func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Error("err", err)
	}
	return info
}

func (m *Master) WatchNodes() {
	resp, err := m.Client.KV.Get(context.TODO(), m.Path, clientv3.WithFirstKey()...)
	if err != nil {
		log.Error("get keys error:%v", err)
	}
	for _, v := range resp.Kvs {
		var info ServiceInfo
		json.Unmarshal(v.Value, &info)
		m.AddNode(string(v.Key), &info)
	}
	rch := m.Client.Watch(context.TODO(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("join -> [%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				fmt.Printf("delete -> [%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				delete(m.Nodes, string(ev.Kv.Key))
			}
		}
	}
}
