package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"server/internal/model"
	"server/internal/repo/activityuser"
	"server/pkg/logger"
)

type MQ chan *model.MsgModel

// InitBroadcast 初始化广播协程
func InitBroadcast(activityUserRepo activityuser.RepoInterface) {
	go BroadcastActivityMsg(&ActivityMQ, activityUserRepo)
}

// BroadcastActivityMsg 广播展会消息
func BroadcastActivityMsg(mq *MQ, activityUserRepo activityuser.RepoInterface) {
	logger.Logger.Info("prepare to push activity message")
	for {
		msg := <-*mq
		byteMsg, _ := json.Marshal(msg)
		// 获取该展会下的所有用户
		uas, err := activityUserRepo.SelectByAID(msg.ActivityID)
		if err != nil {
			continue
		}

		for _, ua := range *uas {
			// 获取该用户对应的 websocket 连接
			// 遍历所有在线用户
			conn, ok := Clients.Get(ua.UserID)
			if !ok {
				continue
			}
			logger.Logger.Infof("send message to user: %d", ua.UserID)
			// 发送消息
			err := conn.WriteMessage(websocket.TextMessage, byteMsg)
			if err != nil {
				continue
			}
		}
	}
}
