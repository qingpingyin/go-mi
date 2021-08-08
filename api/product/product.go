package product

import (
	"MI/pkg/logger"
	service "MI/service/product"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Product(c *gin.Context){
	Cid := c.Query("cid")
	Page := c.DefaultQuery("page", "1")
	PageSize := c.DefaultQuery("pageSize", "7")
	Is_recursion := c.Query("is_recursion")
	cid, _ := strconv.Atoi(Cid)
	page,_ := strconv.Atoi(Page)
	pageSize, _ := strconv.Atoi(PageSize)
	is_recursion,_:= strconv.ParseBool(Is_recursion)
	service.GetProductByCid(c,page,pageSize,cid,is_recursion)
}

func GetProductDetail(c *gin.Context){
	pid := c.Query("pid")
	id, err := strconv.Atoi(pid)
	if err != nil {
		logger.Logger.Info(err)
	}
	service.GetProductByPid(c,id)

}

func GetProductBySearch(c *gin.Context){
	search := c.Query("search")
	Page := c.DefaultQuery("page", "1")
	PageSize := c.DefaultQuery("pageSize", "20")
	page,_ := strconv.Atoi(Page)
	pageSize, _ := strconv.Atoi(PageSize)
	service.Search(c,search,page,pageSize)
}