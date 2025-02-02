package wsmanager

import (
	"github.com/gorilla/websocket"
	"sync"
)

// WSManager 管理ws连接
type WSManager struct {
	connections map[int64]*websocket.Conn
	mu          sync.RWMutex // 读写锁
}

func NewWSManager() *WSManager {
	return &WSManager{
		connections: make(map[int64]*websocket.Conn),
	}
}

func (uc *WSManager) Add(userID int64, conn *websocket.Conn) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	uc.connections[userID] = conn
}

func (uc *WSManager) Remove(userID int64) {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	delete(uc.connections, userID)
}

func (uc *WSManager) Get(userID int64) (*websocket.Conn, bool) {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	conn, exists := uc.connections[userID]
	return conn, exists
}
