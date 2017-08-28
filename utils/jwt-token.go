package utils

import (
	"encoding/json"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/redis"
	"fmt"
)

//GetUserFromToken 从原始Authorization中获取用户信息
func GetUserFromToken(bearerToken string) (user models.PayLoad, err error) {
	if len(bearerToken) <= 0 {
		err = errors.New("Authorization token is null")
		return
	}
	tokenString := bearerToken[7:len(bearerToken)]
	//2016-9-6 校验缓存中是否存在token
	redisToken := redis.Instance().Get("token." + tokenString)
	if redisToken == nil || string(redisToken.([]byte)) != tokenString {
		err = errors.New("wrong token")
		return
	}
	payloadString := strings.Split(tokenString, ".")[1]
	sDec, err := jwt.DecodeSegment(payloadString)
	fmt.Println("解析用户token.....")
	json.Unmarshal(sDec, &user)
	return
}
