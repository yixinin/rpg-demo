package gateway

import (
	"errors"
	"rpg-demo/db"
	"rpg-demo/protocol"
	"sync"

	"github.com/davyxu/cellnet"
	// "golang.org/x/net/websocket"
)

// errors
var (
	RepeatConn   = errors.New("conn is already connected")
	ConnNotExist = errors.New("conn is not exist")
)

type Connection struct {
	Uid int64
	// conn  *websocket.Conn
	sess     cellnet.Session
	Token    string
	CenterIP string
	GameIP   string
}

type Server struct {
	sync.RWMutex
	connMap  map[int64]*Connection
	tokenMap map[string]int64
	sessMap  map[int64]int64
}

func NewServer() *Server {
	return &Server{
		connMap:  make(map[int64]*Connection, 10000),
		tokenMap: make(map[string]int64, 10000),
		sessMap:  make(map[int64]int64, 10000),
	}
}

func (s *Server) Add(conn *Connection) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.connMap[conn.Uid]; ok {
		return RepeatConn
	}
	s.connMap[conn.Uid] = conn
	s.sessMap[conn.sess.ID()] = conn.Uid
	s.tokenMap[conn.Token] = conn.Uid
	return nil
}

func (s *Server) DelConn(conn *Connection) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.connMap[conn.Uid]; ok {
		delete(s.connMap, conn.Uid)
		delete(s.sessMap, conn.sess.ID())
		delete(s.tokenMap, conn.Token)
		return nil
	}
	return ConnNotExist
}

func (s *Server) Del(uid int64) error {
	s.Lock()
	defer s.Unlock()
	if conn, ok := s.connMap[uid]; ok {
		delete(s.connMap, uid)
		delete(s.sessMap, conn.sess.ID())
		delete(s.tokenMap, conn.Token)
		return nil
	}
	return ConnNotExist
}

func (s *Server) Replace(conn *Connection) {
	s.Lock()
	defer s.Unlock()
	s.connMap[conn.Uid] = conn
	s.sessMap[conn.sess.ID()] = conn.Uid
	s.tokenMap[conn.Token] = conn.Uid
}

func (s *Server) Get(uid int64) (*Connection, bool) {
	s.RLock()
	defer s.RUnlock()
	conn, ok := s.connMap[uid]

	return conn, ok
}
func (s *Server) GetByToken(token string) (*Connection, error) {
	s.RLock()
	defer s.RUnlock()
	uid, ok := s.tokenMap[token]
	if !ok {
		return nil, ConnNotExist
	}
	conn, ok := s.connMap[uid]
	if !ok {
		return nil, ConnNotExist
	}
	return conn, nil
}

func (s *Server) Exist(uid int64) bool {
	s.RLock()
	defer s.RUnlock()
	conn, ok := s.connMap[uid]
	if !ok {
		return ok
	}
	return conn.Uid != 0
}

func (s *Server) TokenExist(token string) bool {
	s.RLock()
	defer s.RUnlock()
	uid, ok := s.tokenMap[token]
	if !ok {
		return false
	}
	conn, ok := s.connMap[uid]
	if !ok {
		return false
	}
	return conn.Uid != 0
}

func (s *Server) GetToken(uid int64) string {
	s.RLock()
	defer s.RUnlock()
	conn, ok := s.connMap[uid]
	if !ok {
		return ""
	}
	return conn.Token
}
func (s *Server) GetUid(token string) int64 {
	s.RLock()
	defer s.RUnlock()
	uid, ok := s.tokenMap[token]
	if !ok {
		return 0
	}
	conn, ok := s.connMap[uid]
	if !ok {
		return 0
	}
	return conn.Uid
}

var tokenPrefix = "user:token:"

func (s *Server) CheckUser(header *protocol.ReqHeader) (bool, error) {
	s.Lock()
	defer s.Unlock()

	if uid, ok := s.tokenMap[header.Token]; ok {
		header.Uid = uid
		return true, nil
	}

	userId, err := db.Redis.Get(tokenPrefix + header.Token).Int64()
	if err != nil {
		return false, err
	}
	header.Uid = userId
	return true, err
}
