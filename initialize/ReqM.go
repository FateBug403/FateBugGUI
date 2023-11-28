package initialize

import (
	"github.com/FateBug403/FateBugGUI/global"
	"github.com/adjust/rmq/v4"
	redis2 "github.com/go-redis/redis/v8"
)

// InitQueue 初始化消息队列
func InitQueue() error {
	var err error
	redisClient := redis2.NewClient(&redis2.Options{Network: "tcp", Addr: "127.0.0.1:6379", DB: 1, Password: ""})
	global.GVA_RmqC, err = rmq.OpenConnectionWithRedisClient("FateBugGUI-service", redisClient, nil) // 声明链接名字为FateBugGUI-service
	if err != nil {
		return err
	}
	return nil
}