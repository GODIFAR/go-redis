package main

import (
	"context"

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

	// //LPush  从列表左边插入数据
	// rdb.LPush(ctx, "demo", "ai")
	// err := rdb.LPush(ctx, "demo", "1", "2", "3", "4").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //LPushX  仅当列表存在的时候才插入数据
	// err := rdb.LPushX(ctx, "demo", "5").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //RPop  从列表的右边删除第一个数据，并返回删除的数据
	// val, err := rdb.RPop(ctx, "demo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// RPush插入一个数据
	//rdb.RPush(ctx, "demo", "data1")
	// 支持一次插入任意个数据
	// err := rdb.RPush(ctx, "demo", 1, 2, 3, 4, 5).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //RPushX  仅当列表存在的时候才插入数据
	// err := rdb.RPushX(ctx, "demo", "k").Err()
	// if err != nil {
	// 	panic(err)
	// }

	//LPop  从列表左边删除第一个数据，并返回删除的数据
	// val, err := rdb.LPop(ctx, "demo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)
	// //RPop 从列表右边删除第一个数据，并返回数据
	// val, err := rdb.RPop(ctx, "demo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //LLen  返回列表的大小
	// val, err := rdb.LLen(ctx, "demo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //LRange  返回列表的一个范围内的数据，也可以返回全部数据 0到-1代表全部数据
	// val, err := rdb.LRange(ctx, "demo", 0, -1).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	//LRem  删除列表中的数据
	// err := rdb.LRem(ctx, "demo", 2, 1).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// //从右边开始删除
	// rdb.LRem(ctx, "demo", -2, 1)
	// // 删除全部指定值
	// rdb.LRem(ctx, "demo", 0, 2)

	// //LIndex 根据索引坐标，查询列表中的数据
	// val, err := rdb.LIndex(ctx, "demo", 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	//LInsert   在指定位置插入数据
	err := rdb.LInsert(ctx, "demo", "before", 5, "love").Err()
	if err != nil {
		panic(err)
	}
	rdb.LInsert(ctx, "demo", "before", "love", "i")
	rdb.LInsert(ctx, "demo", "after", "love", "you")
}
