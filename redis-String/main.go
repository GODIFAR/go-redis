package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func redisinit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	redisinit()
	//创建上下文
	ctx := context.Background()
	//set方法设置key和value，处理返回的错误，参数（上下文，key名，value值，过期时间）
	// err := rdb.Set(ctx, "gorediskey", "gorediskey", 0).Err()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// //Get方法获取value
	// val, err := rdb.Get(ctx, "goredistestkey").Result()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }

	// //Do方法使用原生命令,返回值是一个interface类型
	// result, err := rdb.Do(ctx, "get", "key").Result()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Println("get:", val)
	// fmt.Print("原生命令：", result.(string))

	// //GetSet 设置新值但是返回旧值
	// oldVal, err := rdb.GetSet(ctx, "gorediskey", "new value").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// // 打印key的旧值
	// fmt.Println("key", oldVal)

	// //SetNX 如果不存在key1，则设置值为value
	// err := rdb.SetNX(ctx, "key3", "value3", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //MGet 批量查询key值
	// vals, err := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(vals)

	// //MSet 批量设置值
	// err := rdb.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //Incr,IncrBy  针对一个key的数值进行递增操作
	// // Incr函数每次加一
	// val, err := rdb.Incr(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", val)
	// // IncrBy函数，可以指定每次递增多少
	// valBy, err := rdb.IncrBy(ctx, "key", 3).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", valBy)
	// // IncrByFloat函数，可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
	// valFloat, err := rdb.IncrByFloat(ctx, "key1", 2.2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", valFloat)

	// //Decr,DecrBy     针对一个key的数值进行递减操作
	// // Decr函数每次减一
	// val, err := rdb.Decr(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", val)
	// // DecrBy函数，可以指定每次递减多少
	// valBy, err := rdb.DecrBy(ctx, "key", 3).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", valBy)

	// // Del 删除key
	// rdb.Del(ctx, "key")
	// // 删除多个key, Del函数支持删除多个key
	// err := rdb.Del(ctx, "key1", "key2", "key3").Err()
	// if err != nil {
	// 	panic(err)
	// }

	//Expire 设置过期时间
	//rdb.Expire(ctx, "key", 5*time.Second)
	rdb.Set(ctx, "key", "value", 10*time.Second) //expiration:0 永久

}
