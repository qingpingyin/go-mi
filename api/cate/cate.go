package cate

import (
	"MI/models"
	"MI/pkg/logger"
	"MI/service/cate"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Cate(c *gin.Context){
	isNav := c.Query("is_nav")
	if isNav == "" {
		response.RespError(c,"参数不能为空")
		return
	}
	//当参数为空 返回全部类别信息
	is_nav, err := strconv.Atoi(isNav)
	if err != nil {
		logger.Logger.Info("类型转化失败：",err)
	}
	var list []models.Categories
	switch is_nav {
		case 0:
			//nav导航栏 递归类别
			list = cate.CateTree(0)
		case 1:
			//header-nav 导航栏 流动显示
			if list,err = models.GetCategoriesBy("is_nav=?",is_nav);err == nil {
				for i, v := range list {
					product ,_ := models.GetAllProductBy(1,6,"cid=?",int(v.Id))
					list[i].Product = product
				}
			}
	}
	response.RespData(c,"",list)
}


