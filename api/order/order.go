package order

import (
	"MI/models"
	"MI/models/req"
	"MI/pkg/validate"
	service "MI/service/order"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

var OrderTrans = map[string]string{"Uid":"用户id","AddressId":"地址id","Pids":"商品id"}
func CreateOrder(c *gin.Context){
	var orderReq req.OrderReq
	if err := c.BindJSON(&orderReq);err != nil {
		msg := validate.TransTagName(&OrderTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	service.CreateOrder(c,orderReq)
}


func OrderList(c *gin.Context){

	Uid := c.Query("uid")
	Page := c.DefaultQuery("page","1")
	PageSize := c.DefaultQuery("pageSize","5")

	uid,_ := strconv.Atoi(Uid)
	page,_ := strconv.Atoi(Page)
	pageSize, _ := strconv.Atoi(PageSize)
	service.OrderList(c,page,pageSize,uid)
}

func GetOrderById(c *gin.Context){
	oid := c.Query("oid")
	service.GetOrderById(c,oid)
}
func GetOrderCountBy(c *gin.Context){

	Uid := c.Query("uid")
	Status := c.Query("status")

	uid,_ := strconv.Atoi(Uid)
	status,_ := strconv.Atoi(Status)

	count := models.GetOrderCountBy("uid=? and pay_status=?",uid, status)

	response.RespData(c,"",count)
}