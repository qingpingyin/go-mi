package service

import (
	"MI/models"
	"MI/pkg/logger"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)


func Address(c *gin.Context,uid string) {
	id, _ := strconv.Atoi(uid)
	address, err := models.GetAddressBy("uid=?",id)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	response.RespData(c,"",address)
}

func CreateAddress(c *gin.Context,req models.Address){

	if count := req.Count(int(req.Uid));count>=8{
		response.RespError(c,"超过最大限制")
		return
	}
	if err := req.CreateAddress();err != nil {
		logger.Logger.Info(err)
		response.RespError(c,"添加地址失败")
		return
	}
	response.RespSuccess(c,"")
}

func DeleteAddressById(c *gin.Context,id int){
	address := models.Address{
		Id: uint(id),
	}
	if err := address.DeleteAddressById();err != nil {
		logger.Logger.Info(err)
		return
	}
	response.RespSuccess(c,"")

}

func GetAddressById(c *gin.Context,aid int){
	address, err := models.GetAddressByWhere("id=?", aid)
	if err != nil {
		logger.Logger.Error("select address err:",err)
		response.RespError(c,"收货地址不存在")
		return
	}
	response.RespData(c,"",address)
}