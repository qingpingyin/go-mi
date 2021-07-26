package sms

import (
	"MI/pkg/cache"
	"context"
	"time"
)

//将手机号跟验证码存入redis
func SmsSet(key,value string)error{
	return cache.Set(context.Background(),key,value,60*time.Second)

}
func SmsGet(key string)(string,error) {
	return cache.Get(context.Background(),key)
}