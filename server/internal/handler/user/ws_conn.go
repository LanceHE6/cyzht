package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server/internal/ws"
	"server/pkg/jwt"
	"server/pkg/logger"
	"server/pkg/response"
	"time"
)

// ws升级
var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// OnlineHeartbeat 在线心跳
func (s userHandler) OnlineHeartbeat(ctx *gin.Context) {
	// 从查询参数中获取 token
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(11, "未提供token", nil))
		return
	}
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(10, "无法建立ws连接", err))
		return
	}
	// 鉴权获取用户信息
	claims, ok := jwt.Check(token)
	if !ok {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("认证失败"))
		return
	}
	// 使用协程处理ws连接
	go s.handleWebSocketConnection(conn, claims)
}

// handleWebSocketConnection 处理 WebSocket 连接
func (s userHandler) handleWebSocketConnection(conn *websocket.Conn, claims jwt.MyClaims) {
	// 设置用户在线状态
	if err := s.UserRepo.SetUserOnline(claims.ID, claims.SessionID, 0); err != nil { // 假设初始活动ID为0
		logger.Logger.Errorf("设置用户 %d 在线状态失败: %s", claims.ID, err.Error())
		_ = conn.WriteMessage(websocket.TextMessage, []byte("设置在线状态失败"))
		return
	}
	logger.Logger.Debugf("用户 %d 已上线", claims.ID)

	// 添加连接到管理器
	ws.Clients.Add(claims.ID, conn)
	// 关闭连接
	defer func(conn *websocket.Conn, user jwt.MyClaims) {
		// 设置用户离线状态
		if err := s.UserRepo.SetUserOffline(user.ID); err != nil {
			logger.Logger.Errorf("设置用户 %d 离线状态失败: %s", claims.ID, err.Error())
		}
		logger.Logger.Debugf("用户 %d 已下线", claims.ID)
		err := conn.Close()
		if err != nil {
			return
		}
		// 移除连接
		ws.Clients.Remove(claims.ID)
	}(conn, claims)

	go s.heartbeat(conn, claims) // 启动心跳协程

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 设置用户离线状态
			if err := s.UserRepo.SetUserOffline(claims.ID); err != nil {
				logger.Logger.Errorf("设置用户 %d 离线状态失败: %s", claims.ID, err.Error())
			}
			return // 读取消息失败可能表示连接已断开
		}
		if message != nil {
			if string(message) == "pong" {
				logger.Logger.Debugf("收到用户 %d 回复心跳", claims.ID)
				continue
			}
			logger.Logger.Debug(fmt.Sprintf("收到用户 %d 消息: %s", claims.ID, message))

		}
	}
}

// heartbeat
//
//	@Description: 心跳
//	@param conn websocket连接
func (s userHandler) heartbeat(conn *websocket.Conn, claims jwt.MyClaims) {
	// 每5秒发送一次心跳包
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger.Logger.Debugf("向用户 %d 发送心跳", claims.ID)
			// 发送心跳
			if err := conn.WriteMessage(websocket.TextMessage, []byte("来自服务端的ping")); err != nil {
				// 设置用户离线状态
				if err := s.UserRepo.SetUserOffline(claims.ID); err != nil {
					logger.Logger.Errorf("设置用户 %d 离线状态失败: %s", claims.ID, err.Error())
				}
				logger.Logger.Debugf("用户 %d 已上线", claims.ID)
				return // 发送心跳失败可能表示连接已断开
			}
			// 更新用户心跳
			if err := s.UserRepo.UpdateUserHeartbeat(claims.ID); err != nil {
				logger.Logger.Errorf("更新用户 %d 心跳失败: %s", claims.ID, err.Error())
				return
			}
		}
	}
}
