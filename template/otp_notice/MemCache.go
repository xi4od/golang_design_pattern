package otp_notice

import "fmt"

type MemCache struct {
	Otp
}

func NewMemCache() *MemCache {
	return &MemCache{}
}

func(t MemCache) CacheOtp() {
	fmt.Println("Cache otp use memcache")
}
