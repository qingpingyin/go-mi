package service

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/logger"
	"MI/pkg/setting"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"time"
)

func DoLogin(c *gin.Context,user models.Users){
	accessExpTime := time.Now().Add(time.Duration(setting.JwtConf.ExpTime)*time.Hour)
	refreshExpTime := time.Now().Add(time.Duration(setting.JwtConf.ExpTime+1800)*time.Hour)
	accessToken, err := jwt.GenerateToken(user, accessExpTime)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	refreshToken,err := jwt.GenerateToken(user,refreshExpTime)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	response.RespData(c,"登陆成功",map[string]string{
		"access_token":accessToken,
		"refreshToken":refreshToken,
	})

	return
}

