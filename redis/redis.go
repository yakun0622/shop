package redis

import (
	"sync"

	"github.com/astaxie/beego/cache"
	//redis driver
	_ "github.com/astaxie/beego/cache/redis"
)

var redisInstace *cache.Cache
var once sync.Once

//Instance Singleton redisInstace
func Instance() cache.Cache {
	once.Do(func() {
		//cache 指定redis dbNum=0
		redisCache, err := cache.NewCache("redis", `{"key":"cache","conn":"139.196.201.45:6379","dbNum":"0","password":"cb_password"}`)
		if err == nil {
			redisInstace = &redisCache
		}
	})
	return *redisInstace
}
