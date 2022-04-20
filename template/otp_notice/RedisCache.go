package otp_notice

import "fmt"

type RedisCache struct {
	Otp
}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

func(t RedisCache) CacheOtp() {
	fmt.Println("Cache otp use Redis")
}
