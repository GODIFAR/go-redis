package main

import (
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
	//ctx := context.Background()
	// // ZAdd添加一个集合元素到集合中， 这个元素的分数是2.5，元素名是zhangsan
	// err := rdb.ZAdd(ctx, "key", &redis.Z{Score: 2.5, Member: "zhangsan"}).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// rdb.ZAdd(ctx, "key",
	// 	&redis.Z{Score: 3, Member: "lisi"},
	// 	&redis.Z{Score: 5, Member: "wangwu"},
	// )

	// //ZCard  返回集合元素个数
	// val, err := rdb.ZCard(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //ZCount  统计某个分数范围内的元素个数
	// size, err := rdb.ZCount(ctx, "key", "1", "5").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(size)
	// // 返回： 1<分数<=5 的元素个数
	// // 说明：默认第二，第三个参数是大于等于和小于等于的关系。
	// // 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
	// //size, err := rdb.ZCount(ctx, "key", "(1", "5").Result()

	// //ZIncrBy  增加元素的分数
	// val, err := rdb.ZIncrBy(ctx, "key", 5, "zhangsan").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //ZRange,ZRevRange  返回集合中某个索引范围的元素，根据分数从小到大排序 0 -1代表返回全部数据
	// //ZRevRange 用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
	// val, err := rdb.ZRange(ctx, "key", 0, -1).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, val := range val {
	// 	fmt.Println(val)
	// }

	//ZRevRangeByScore 用法类似ZRangeByScore，区别是元素根据分数从大到小排序。

	// 	//ZRangeByScore  根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	// // 初始化查询条件， Offset和Count用于分页
	// op := redis.ZRangeBy{
	// 	Min:"2", // 最小分数
	// 	Max:"10", // 最大分数
	// 	Offset:0, // 类似sql的limit, 表示开始偏移量
	// 	Count:5, // 一次返回多少数据
	// }
	// vals, err := rdb.ZRangeByScore(ctx,"key", &op).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, val := range vals {
	// 	fmt.Println(val)
	// }
	// // ZRangeByScoreWithScores  根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	// // 初始化查询条件， Offset和Count用于分页
	// op := redis.ZRangeBy{
	// 	Min:    "2",  // 最小分数
	// 	Max:    "10", // 最大分数
	// 	Offset: 0,    // 类似sql的limit, 表示开始偏移量
	// 	Count:  5,    // 一次返回多少数据
	// }
	// vals, err := rdb.ZRangeByScoreWithScores(ctx, "key", &op).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, val := range vals {
	// 	fmt.Println(val.Member) // 集合元素
	// 	fmt.Println(val.Score)  // 分数
	// }

	// //ZRem  删除集合元素
	// err := rdb.ZRem(ctx, "key", "zhangsan").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// // 删除集合中的元素zhangsan和zhangsan1
	// // 支持一次删除多个元素
	// rdb.ZRem(ctx, "key", "zhangsan", "lisi")

	// //ZRemRangeByRank  根据索引范围删除元素
	// err := rdb.ZRemRangeByRank(ctx, "key", 0, 5).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// //ZScore  查询元素对应的分数
	// val, err := rdb.ZScore(ctx, "key", "ljc").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //ZRank  根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	// rk, _ := rdb.ZRank(ctx, "key", "zhangsan").Result()
	// fmt.Println(rk)
	// //ZRevRank的作用跟ZRank一样，区别是ZRevRank是按分数从大到小排序。
}
