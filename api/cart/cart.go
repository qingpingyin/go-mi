package cart

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/validate"
	"MI/service/cart"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)
var CartTrans =map[string]string{"Uid":"用户id","Pid":"商品id","Num":"商品数量"}
func AddCart(c *gin.Context){

	var item models.Item
	if err := c.BindJSON(&item);err != nil {
		msg := validate.TransTagName(&CartTrans,err)
		response.RespError(c,msg)
		return
	}
	service.AddCart(c,item)
}

func CartList(c *gin.Context){

	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)
	uid := strconv.Itoa(int(userInfo.Id))
	service.CartList(c,uid)

}
func CartUpdateNum(c *gin.Context){

	uid := c.Param("uid")
	pid := c.Param("pid")
	num := c.Param("num")

	item := models.Item{
		Uid: uid,
		Pid: pid,
		Num: num,
	}
	service.CartUpdateNum(c,item)
}

func DeleteCart(c *gin.Context){

	uid := c.Param("uid")
	pid := c.Param("pid")

	service.DeleteCart(c,uid,pid)
}
