package user

//// ws升级
//var upgrade = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//// OnlineHeartbeat
////
////	@Description: 在线心跳
////	@receiver s userService
////	@return gin.HandlerFunc
//func (s userService) OnlineHeartbeat() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		conn, err := upgrade.Upgrade(context.Writer, context.Request, nil)
//		if err != nil {
//			context.JSON(http.StatusBadRequest, response.ErrorResponse(10, "无法建立ws连接", err))
//			return
//		}
//		// 鉴权获取用户信息
//		ok, user := Auth(context.GetHeader("Authorization"))
//		if ok {
//			_ = s.UserRepo.UpdateOnlineStatus(user.ID, 0)
//			logger.Logger.Debug("用户 " + user.Nickname + " 已上线")
//		} else {
//			_ = conn.WriteMessage(websocket.TextMessage, []byte("认证失败"))
//			return
//		}
//		// 关闭连接
//		defer func(conn *websocket.Conn, user jwt.MyClaims) {
//			_ = s.UserRepo.UpdateOnlineStatus(user.ID, 0)
//			logger.Logger.Debug("用户 " + user.Nickname + " 已下线")
//			err := conn.Close()
//			if err != nil {
//				return
//			}
//		}(conn, user)
//
//		go s.heartbeat(conn, user) // 启动心跳协程
//
//		for {
//			_, message, err := conn.ReadMessage()
//			if err != nil {
//				// TODO 将用户在线状态改为离线
//				_ = s.UserRepo.UpdateOnlineStatus(user.ID, 0)
//				return // 读取消息失败可能表示连接已断开
//			}
//			// 鉴权
//			if message != nil {
//				logger.Logger.Debug(fmt.Sprintf("收到消息: %s", message))
//
//			}
//		}
//	}
//}
//
//// heartbeat
////
////	@Description: 心跳
////	@param conn websocket连接
//func (s userService) heartbeat(conn *websocket.Conn, user jwt.MyClaims) {
//	// 每5秒发送一次心跳包
//	ticker := time.NewTicker(5 * time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			logger.Logger.Debug("发送心跳")
//			// 发送心跳
//			if err := conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
//				// TODO 将用户在线状态改为离线
//				_ = s.UserRepo.UpdateOnlineStatus(user.ID, 0)
//				logger.Logger.Debug("用户 " + user.Nickname + " 已下线")
//				return // 发送心跳失败可能表示连接已断开
//			}
//		}
//	}
//}
