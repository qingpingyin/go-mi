package collect

import (
	"MI/models"
	"MI/models/req"
	"MI/pkg/jwt"
	"MI/pkg/logger"
	"MI/pkg/validate"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

var CollectTrans = map[string]string{"Uid":"用户id","Pid":"商品id"}
func CreateCollect(c *gin.Context){
	var collectReq req.CollectReq
	if err := c.BindJSON(&collectReq);err != nil {
		msg := validate.TransTagName(&CollectTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	collect := models.Collect{
		Uid: collectReq.Uid,
		Pid: collectReq.Pid,
	}
	//判断该商品是否已经加入收藏
	if collect,_ := models.GetCollectByWhere("pid=? and uid=?",collectReq.Pid,collectReq.Uid);collect.Id>0{
		response.RespError(c,"该商品已收藏")
		return
	}

	if err := collect.CreateCollect();err != nil{
		logger.Logger.Error("create collect err:",err)
		response.RespError(c,"收藏失败")
		return
	}
	response.RespSuccess(c,"收藏成功")
}

func List(c *gin.Context){
	uid := c.Query("uid")
	id,_ := strconv.Atoi(uid)

	collects := models.GetCollectBy("uid=?",id)
	var data []models.Product
	for _,v := range collects {
		if product,err := models.GetProductByWhere("id=?",v.Pid);err == nil{
			data = append(data,product)
		}
	}
	response.RespData(c,"",data)
}
func DelCollect(c *gin.Context){
	uid,_ := strconv.Atoi(c.Param("uid"))
	pid,_ := strconv.Atoi(c.Param("pid"))
	if err := models.DelCollect("uid=? and pid=?",uid,pid);err != nil{
		logger.Logger.Error("del collect err:",err)
		response.RespError(c,"移除失败")
	}
	response.RespSuccess(c,"")
}

func CollectCount(c *gin.Context){
	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)

	count := models.GetCollectCountBy("uid=?", userInfo.Id)

	response.RespData(c,"",count)
}