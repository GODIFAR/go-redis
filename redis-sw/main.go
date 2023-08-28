package main

import (
	"context"

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

	// // 开启一个TxPipeline事务
	// pipe := rdb.TxPipeline()
	// // 执行事务操作，可以通过pipe读写redis
	// incr := pipe.Incr(ctx, "tx_pipeline_counter") //Incr 自增
	// pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	// // 上面代码等同于执行下面redis命令:
	// //     MULTI
	// //     INCR pipeline_counter
	// //     EXPIRE pipeline_counts 3600
	// //     EXEC
	// // 通过Exec函数提交redis事务
	// _, err := pipe.Exec(ctx)
	// // 提交事务后，我们可以查询事务操作的结果
	// // 前面执行Incr函数，在没有执行Exec函数之前，实际上还没开始运行。
	// fmt.Println(incr.Val(), err)

	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		v++

		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set(ctx, "key", v, 0)
			return nil
		})
		return err
	}
	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(ctx,fn, "key1", "key2", "key3")
	for i := 0; i < 3; i++ {
		err := rdb.Watch(ctx, fn, "key")
		if err == nil {
			break
		}
		if err == redis.TxFailedErr {
			continue
		}
	}
}
