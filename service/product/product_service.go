package service

import (
	. "MI/models"
	"MI/pkg/logger"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
)

func GetProductByCid(c *gin.Context,page,pageSize,cid int,is_recursion bool){
	if is_recursion{
		categories := GetAllCate("parent_id=?", cid)
		//根据 一级类别id查询子类id
		var cates []uint
		for _, v := range categories {
			cates = append(cates, v.Id)
		}
		list, err := GetAllProductBy(page, pageSize, "cid in (?) ", cates)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
		response.RespData(c,"",list)
	}else {
		list, err := GetAllProductBy(page, pageSize, "cid=? ", cid)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
		response.RespData(c,"",list)
	}

}

func GetProductByPid(c *gin.Context,pid int){

	product, err := GetProductByWhere("id=?", pid)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	response.RespData(c,"",product)
}

func Search(c *gin.Context,search string,page,pageSize int){
	//首先判断 search == categories
	//这里拿到 二级目录
	cate := GetAllCate("parent_id !=?",0)
	for _, v := range cate {
		if v.CategoriesName == search{
			list,_ := GetAllProductBy(page,pageSize,"cid=?",v.Id)
			total := GetCountProductByWhere("cid=?",v.Id)
			response.RespData(c,"",map[string]interface{}{
				"total":total,
				"list":list,
			})
			return
		}
	}

	if list,err := GetAllProductBy(page,pageSize,"title like CONCAT('%',?,'%')",search);err == nil {
		total := GetCountProductByWhere("title like CONCAT('%',?,'%')",search)
		response.RespData(c,"",map[string]interface{}{
			"total":total,
			"list":list,
		})
		return
	}
	response.RespData(c,"","")
}