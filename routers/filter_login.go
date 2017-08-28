package routers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/yakun0622/shop/constant"

	"strings"

	"github.com/astaxie/beego/context"
	"github.com/yakun0622/shop/redis"
)

func init() {
	//登录Filter
	result := make(map[string]interface{})
	var FilterUser = func(ctx *context.Context) {

		tokenString := ctx.Request.Header.Get("Authorization")

		//无token
		if strings.Contains(strings.ToLower(tokenString), "bearer") {
			tokenString = tokenString[7:len(tokenString)]

			//2016-8-16 校验缓存中是否存在token
			redisToken := redis.Instance().Get("token." + tokenString)
			if redisToken == nil || string(redisToken.([]byte)) != tokenString {
				result["code"] = constant.JWTTokenInvalid
				result["msg"] = constant.StatusText(constant.JWTTokenInvalid)
				ctx.Output.JSON(result, false, false)
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return constant.JWTPublicKey, nil
			})

			//token 正常
			if token != nil && token.Valid {
				//TODO:根据行为鉴权
				return
			}

			//token异常
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					result["code"] = constant.JWTTokenInvalidFormat
					result["msg"] = constant.StatusText(constant.JWTTokenInvalidFormat)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					result["code"] = constant.JWTTokenExpired
					result["msg"] = constant.StatusText(constant.JWTTokenExpired)
				} else {
					result["code"] = constant.JWTTokenInvalid
					result["msg"] = constant.StatusText(constant.JWTTokenInvalid)
				}
			} else {
				result["code"] = constant.JWTTokenInvalid
				result["msg"] = constant.StatusText(constant.JWTTokenInvalid)
			}
			ctx.Output.JSON(result, false, false)

		} else {
			result["code"] = constant.JWTTokenNeeded
			result["msg"] = constant.StatusText(constant.JWTTokenNeeded)
			ctx.Output.JSON(result, false, false)
		}
	}
	beego.InsertFilter("/v1/member/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v1/member_ext/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v1/group/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v1/cart/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v1/favorites/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v1/favorites_folder/*", beego.BeforeRouter, FilterUser)
}
