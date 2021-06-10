package service

import (
	"MI/models"
	"MI/pkg/logger"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
)

func GetProductByCid(c *gin.Context,page,pageSize,cid int){
	categories := models.GetAllCate("parent_id=?", cid)

	//根据 一级类别id查询子类id
	var cates []uint
	for _, v := range categories {
		cates = append(cates, v.Id)
	}
	list, err := models.GetAllProductBy(page, pageSize, "cid in (?) ", cates)
	if err != nil {
		logger.Logger.Error(err)
	}
	response.RespData(c,"",list)
}

func GetProductByPid(c *gin.Context,pid int){

	product, err := models.GetProductByWhere("id=?", pid)
	if err != nil {
		logger.Logger.Error(err)
	}
	response.RespData(c,"",product)
}