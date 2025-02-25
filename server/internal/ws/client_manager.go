package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// ClientManager 管理ws连接
type ClientManager struct {
	connections map[int64]*websocket.Conn
	mu          sync.RWMutex // 读写锁
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		connections: make(map[int64]*websocket.Conn),
	}
}

func (uc *ClientManager) Add(userID int64, conn *websocket.Conn) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	uc.connections[userID] = conn
}

func (uc *ClientManager) Remove(userID int64) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	delete(uc.connections, userID)
}

func (uc *ClientManager) Get(userID int64) (*websocket.Conn, bool) {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	conn, exists := uc.connections[userID]
	return conn, exists
}
