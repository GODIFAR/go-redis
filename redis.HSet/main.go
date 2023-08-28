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

	//HSet 根据key和field字段设置 field 字段的值
	// user_1 是hash key，username 是字段名, zhangsan是字段值
	// err := rdb.HSet(ctx, "user_1", "username", "zhangsan").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //HGet 根据key和field字段，查询 field 字段的值
	// username, err := rdb.HGet(ctx, "user_1", "username").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("username:", username)

	// //HGetAll 根据key查询所有字段和值
	// data, err := rdb.HGetAll(ctx, "user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// for failed, val := range data {
	// 	fmt.Println(failed, val)
	// }

	// //HIncrBy 根据key和field字段，累加字段的数值
	// count, err := rdb.HIncrBy(ctx, "user_1", "count", 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(count)

	// //HKeys  根据key返回所有字段名
	// keys, err := rdb.HKeys(ctx, "user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(keys)

	// 	//HLen  根据key，查询hash的字段数量
	// 	len, err := rdb.HLen(ctx, "user_1").Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(len)

	// //HMGet  根据key和多个字段名，批量查询多个hash字段值
	// mget, err := rdb.HMGet(ctx, "user_1", "username", "count").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(mget)

	// // HMSet  根据key和多个字段名和字段值，批量设置hash字段值
	// data := make(map[string]interface{})
	// data["id"] = 1
	// data["username"] = "list"
	// err := rdb.HMSet(ctx, "key", data).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //HSetNX  如果field字段不存在，则设置hash字段值
	// err := rdb.HSetNX(ctx, "key", "name", "love").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //HDel  根据key和字段名，删除hash字段，支持批量删除hash字段
	// rdb.HDel(ctx,"key","id")
	// rdb.HDel(ctx,"user_1","name","id")

	// HExists  检测hash字段名是否存在
	exist, err := rdb.HExists(ctx, "key", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(exist)
}
