package cache

import (
	"MI/pkg/setting"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)
var (
	rdb *redis.Client
)

// 初始化连接
func SetUp(){
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",setting.RedisConf.Host,setting.RedisConf.Port),
		Password: setting.RedisConf.Password,  // no password set
		DB:       setting.RedisConf.Db,   // use default DB
		PoolSize: setting.RedisConf.PoolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb.Ping(ctx).Result()
}

func Set(ctx context.Context,key,value string)error{

	return rdb.Set(ctx,key,value,time.Second*60).Err()

}
func Get(ctx context.Context,key string)string{

	return rdb.Get(ctx,key).String()

}