package ws

// Clients 全局ws连接管理
var Clients = NewClientManager()

// ActivityMQ 展会消息队列
var ActivityMQ = make(MQ, 100)
