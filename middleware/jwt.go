package middleware

import (
	"MI/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const JwtName = "Authorization"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg": "token不存在",
			})
			return
		}
	claims, err := jwt.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":  "无效token",
			})
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":  "token已过期",
			})
			return
		}
		//从redis中判断该token是否加入黑名单
		has:=jwt.IsBlackExist(string(claims.Id),token)
		if has {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":  "token已失效",
			})
			return
		}
		// 设置用户对象在上下文中，方便后续使用
		c.Set("user", claims)
		c.Next()
	}
}

// 各种方法获取 token
// 为了防范 CSRF 攻击,不获取 query 和 from 里的 token
func getToken(c *gin.Context) (string, error) {
	if token := c.GetHeader(JwtName); token != "" {
		return token, nil
	}

	if token, _ := c.Cookie(JwtName); token != "" {
		return token, nil
	}
	return "", errors.New("没有找到" + JwtName)
}
