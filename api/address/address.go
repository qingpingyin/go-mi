package address

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/validate"
	service "MI/service/address"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

var AddressTrans = map[string]string{
	"ReceiverName":"收货人姓名",
	"ReceiverMobile":"手机号",
	"ReceiverProvince":"省份",
	"ReceiverCity":"区",
	"ReceiverDistrict":"县",
	"ReceiverAddress":"详细地址",
}
func Address(c *gin.Context){
	uid := c.Query("uid")
	service.Address(c,uid)
}
func GetAddressById(c *gin.Context){
	aid := c.Query("aid")
	id,_ := strconv.Atoi(aid)
	service.GetAddressById(c,id)
}
func CreateAddress(c *gin.Context){
	var addressReq models.Address
	if err := c.BindJSON(&addressReq);err != nil {
		msg := validate.TransTagName(&AddressTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)
	addressReq.Uid=userInfo.Id
	service.CreateAddress(c,addressReq)
}

func DeleteAddres(c *gin.Context){
	Id := c.Param("id")
	id,_ := strconv.Atoi(Id)

	service.DeleteAddressById(c,id)
}