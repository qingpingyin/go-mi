package jwt

import (
	"MI/models"
	"MI/pkg/logger"
	"MI/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	RealName string `json:"real_name"`
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}

var JwtKey = []byte(setting.JwtConf.Key)
//生成令牌
func GenerateToken(user models.Users,expTime time.Time)(string,error){
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,Claims{
		Id:user.Id,
		Name: user.NikeName,
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

