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

func Set(ctx context.Context,key,value string,time time.Duration)error{
	return rdb.Set(ctx,key,value,time).Err()

}
func Get(ctx context.Context,key string)(string,error){

	return rdb.Get(ctx,key).Result()

}

func HashHSet(ctx context.Context, key string,value interface{})error{

	return rdb.HSet(ctx, key, value).Err()

}
func HashIsExists(ctx context.Context,key,filed string)bool{
	result, _ := rdb.HExists(ctx, key, filed).Result()
	return result
}

func HashIncrBy(ctx context.Context,key,filed string,count int64)error{
	return rdb.HIncrBy(ctx,key,filed,count).Err()
}

func HashAll(ctx context.Context,key string)(map[string]string,error){
	return  rdb.HGetAll(ctx, key).Result()
}
func HashHGet(ctx context.Context,key string,filed string)(string,error){
	return rdb.HGet(ctx,key,filed).Result()
}
func HashDel(ctx context.Context,key,filed string)error{
	return rdb.HDel(ctx,key,filed).Err()
}