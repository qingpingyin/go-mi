package service

import (
	"MI/models"
	"MI/models/req"
	"MI/pkg/cache"
	"MI/pkg/logger"
	"MI/utils/common"
	"MI/utils/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateOrder(c *gin.Context,req req.OrderReq){

	cacheKey := fmt.Sprintf("cart:user:%d",req.Uid)
	//订单id
	OrderId := common.GenerateOrderSnowId()
	order := models.Order{
		OrderId: OrderId,
		Uid: req.Uid,
		Aid: req.AddressId,

	}
	var payMent float64
	for _, v := range req.Pids {
		//根据购物车中的pid获取redis中的num
		Num, _ := cache.HashHGet(context.Background(), cacheKey, strconv.Itoa(v))
		num,_ := strconv.Atoi(Num)

		//从数据库中查询商品信息
		product, err := models.GetProductByWhere("id=?", v)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
		var orderItem = models.OrderItem{
			OrderId:    OrderId,
			Pid:        v,
			Num:        num,
			Title:      product.Title,
			Price:      product.ShopPrice,
			TotalPrice: float64(num) * product.ShopPrice,
			ImgUrl: product.ImgUrl,
		}
		payMent += orderItem.TotalPrice
		order.OrderItem = append(order.OrderItem,orderItem)
	}
	order.Payment = payMent
	//插入至数据库
	if err := order.Create();err != nil{
		logger.Logger.Error(err)
		response.RespError(c,"创建订单失败")
		return
	}
	for _, item := range req.Pids {
		cache.HashDel(context.Background(),cacheKey,strconv.Itoa(item))
	}

	response.RespData(c,"",order)
}

func OrderList(c *gin.Context,page,pageSize int,uid int){

	list, err := models.GetAllOrderBy(page, pageSize, "uid=?", uid)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	count := models.GetOrderCountBy("uid=?",uid)
	response.RespData(c,"",map[string]interface{}{
		"list":list,
		"count":count,
	})
}

func GetOrderById(c *gin.Context,oid string){

	order,err := models.GetAllOrderByWhere("order_id=?",oid)
	if err != nil {
		logger.Logger.Error("select order info err:",err)
		response.RespError(c,"该订单不存在")
		return
	}
	response.RespData(c,"",order)
}

