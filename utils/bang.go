package utils

import (
	"time"

	"strconv"

	"github.com/yakun0622/shop/config"
	"github.com/yakun0622/shop/redis"
)

//LoginFailBangUser 登录失败次数超过最大登录次数，暂时禁止用户登录
func LoginFailBangUser(identify string, ip string) {
	key := "bang." + identify
	data := redis.Instance().Get(key)
	count := 0
	if data != nil {
		temp := string(data.([]byte))
		if result, err := strconv.Atoi(temp); err == nil && result != 0 {
			count = result + 1
		} else {
			count = 1
		}
	} else {
		count = 1
	}

	//以字符存储数字
	redis.Instance().Put(key, count, time.Duration(config.FailLockMinutes)*time.Minute)
}

//IsLoginCountOver 是否超过最大登录次数
func IsLoginCountOver(identify string) bool {
	key := "bang." + identify
	data := redis.Instance().Get(key)
	if data == nil {
		return false
	}

	temp := string(data.([]byte))
	if count, err := strconv.Atoi(temp); err == nil && count >= config.MaxTryLoginCount {
		return true
	}
	return false
}
