package upload

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/logger"
	"MI/service/upload"
	"MI/utils/common"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)
const maxUploadSize int64 = 2 * 1024 * 1024 // 2 mb

func AvatarUpload(c *gin.Context){
	//从token claims获取个人信息
	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		logger.Logger.Error("upload err:",err)
		response.RespError(c,"图片上传失败")
		return
	}
	if fileHeader.Size>maxUploadSize{
		response.RespError(c,"图片文件超出2M，请重新上传！")
		return
	}
	//检查文件后类型
	if has :=checkFileType(fileHeader.Filename);!has{
		response.RespError(c,"文件格式错误")
		return
	}
	//文件名去重 hash
	en := common.Md5En(fileHeader.Filename)
	//图片验证通过之后  存储bucket
	imgPath, err := upload.UploadFile(file, fileHeader.Size, en)
	if err != nil {
		logger.Logger.Error("上传Bucket err:",err)
		return
	}

	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)

	if err := models.UpdateAvatarBy(imgPath,"id=?",userInfo.Id);err != nil{
		logger.Logger.Error("update user avatar err:",err)
		response.RespError(c,"图片保存失败")
		return
	}

	response.RespSuccess(c,"上传成功")
}

func checkFileType(path string)bool{
	var flag bool
	supports :=[]string{".png", ".jpg"}
	for _, v := range supports {
		if has:=strings.HasSuffix(path,v);has{
			flag = true
		}
	}
	return flag
}
