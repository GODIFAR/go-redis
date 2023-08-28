package main

import (
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
	//ctx := context.Background()

	// //SAdd  添加集合元素(去重)
	// err := rdb.SAdd(ctx, "S", "100").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// rdb.SAdd(ctx, "S", 100, 200, 300)

	// //SCard  获取集合元素个数
	// val, err := rdb.SCard(ctx, "S").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //SIsMember 判断元素是否在集合中
	// val, _ := rdb.SIsMember(ctx, "S", "100").Result()
	// fmt.Println(val)

	// //SMembers 获取集合中所有的元素
	// val, err := rdb.SMembers(ctx, "S").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)

	// //SRem  删除集合元素
	// err := rdb.SRem(ctx, "S", 100).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//SPop,SPopN  随机返回集合中的元素，并且删除返回的元素

	//add()

	// //SPop,SPopN  随机返回集合中的元素，并且删除返回的元素
	// val, err := rdb.SPop(ctx, "S").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)
	// val1, err := rdb.SPopN(ctx, "S", 5).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val1)
}

// func add() {
// 	ctx := context.Background()
// 	err := rdb.SAdd(ctx, "S", 101, 102, 103, 104, 105, 106, 107, 108).Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }
