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
	PageSize := c.DefaultQuery("pageSize", "8")

	cid, _ := strconv.Atoi(Cid)
	page,_ := strconv.Atoi(Page)
	pageSize, _ := strconv.Atoi(PageSize)


	service.GetProductByCid(c,page,pageSize,cid)
}

func GetProductDetail(c *gin.Context){
	pid := c.Query("pid")
	id, err := strconv.Atoi(pid)
	if err != nil {
		logger.Logger.Info(err)
	}
	service.GetProductByPid(c,id)

}