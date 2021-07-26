package jwt

import (
	"MI/models"
	"MI/models/req"
	"MI/pkg/cache"
	"MI/pkg/logger"
	"MI/pkg/setting"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"time"
)

type Claims struct{
	Id uint `json:"id"`
	NikeName string `json:"nike_name"`
	RealName string `json:"real_name"`
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}

var JwtKey = []byte(setting.JwtConf.Key)

//生成令牌
func GenerateToken(user models.Users,expTime time.Time)(string,error){
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,Claims{
		Id:user.Id,
		NikeName: user.NikeName,
		RealName: user.RealName,
		Mobile: user.Mobile,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			Subject:   "go-mi",
		},
	})
	return tokenClaim.SignedString(JwtKey)
}
//解析令牌
func ParseToken(token string)(*Claims,error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	logger.Logger.Error("解析jwt出错 : ", err)
	return nil, err
}

//加入黑名单
func AddBlack(key,value string) error{
	key = "black-token"+key+value
	return  cache.Set(context.Background(), key, value, 3600*24*time.Second)
}
//检查是否存在
func IsBlackExist(key,token string)bool {
	key = "black-token"+key+token
	val, err := cache.Get(context.Background(), key)
	if err  == redis.Nil && val  != token {
		return false
	}
	return true
}

type EmailClaims struct{
	UserID        int   `json:"user_id"`
	Email         string `json:"email"`
	OperationType int   `json:"operation_type"`
	jwt.StandardClaims
}

func GenerateEmailToken(req req.EmailReq) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)

	claims := EmailClaims{
		req.UserID,
		req.Email,
		req.OperationType,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-mi",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtKey)

	return token, err
}
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	logger.Logger.Error("解析jwt出错 : ", err)
	return nil, err
}