package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func redisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func main() {
	redisInit()
	ctx := context.Background()
	// 将"message"消息发送到channel1这个通道上
	rdb.Publish(ctx, "channel1", "message")
	// 查询channel_1通道的订阅者数量
	chs, _ := rdb.PubSubNumSub(ctx, "channel1").Result()
	for ch, count := range chs {
		fmt.Println(ch)    // channel名字
		fmt.Println(count) // channel的订阅者数量
	}

}
