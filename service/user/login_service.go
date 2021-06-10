package service

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/setting"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"time"
)

func DoLogin(c *gin.Context,user models.Users)error{

	accessExpTime := time.Now().Add(time.Duration(setting.JwtConf.ExpTime)*time.Hour)
	refreshExpTime := time.Now().Add(time.Duration(setting.JwtConf.ExpTime+1800)*time.Hour)
	accessToken, err := jwt.GenerateToken(user, accessExpTime)
	if err != nil {
		return err
	}
	refreshToken,err := jwt.GenerateToken(user,refreshExpTime)
	if err != nil {
		return err
	}
	response.RespData(c,"登陆成功",map[string]string{
		"access_token":accessToken,
		"refreshToken":refreshToken,
	})

	return nil
}

