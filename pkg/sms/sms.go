package sms

import (
	"MI/pkg/cache"
	"context"
)

//将手机号跟验证码存入redis
func SmsSet(key,value string)error{
	return cache.Set(context.Background(),key,value)

}
func SmsGet(key string)string {
	return cache.Get(context.Background(),key)
}