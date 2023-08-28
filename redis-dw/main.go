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
	// 订阅channel1这个channel
	sub := rdb.Subscribe(ctx, "channel1")
	// sub.Channel() 返回go channel，可以循环读取redis服务器发过来的消息
	for msg := range sub.Channel() {
		// 打印收到的消息
		fmt.Println(msg.Channel) //频道名称
		fmt.Println(msg.Payload) //内容
	}

	// 订阅channel1这个channel
	//sub := rdb.PSubscribe(ctx, "ch_user_*") //PSubscribe
	// 可以匹配ch_user_开头的任意channel

	// 订阅channel1这个channel
	// 取消订阅
	//sub.Unsubscribe(ctx, "channel1")
}
