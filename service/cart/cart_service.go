package service

import (
	"MI/models"
	"MI/pkg/cache"
	"MI/pkg/logger"
	"MI/utils/common"
	"MI/utils/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddCart(c *gin.Context ,item models.Item){

	cacheKey := fmt.Sprintf("cart:user:%s",item.Uid)
	//判断 cart:key在redis中是否已存在,商品不存在，新增该购物车，商品存在，商品数量+1
	if has := cache.HashIsExists(context.Background(),cacheKey,item.Pid);!has{
		//将cart插入到redis
		var data = map[string]interface{}{
			item.Pid:item.Num,
		}
		if err := cache.HashHSet(context.Background(),cacheKey,data);err != nil {
			logger.Logger.Error(err)
			response.RespError(c,"添加失败")
			return
		}
	}else{
		if err := addNum(cacheKey,item.Pid,1);err != nil {
			logger.Logger.Error(err)
			response.RespError(c,"添加失败")
			return
		}
	}
	response.RespSuccess(c,"添加成功")
}

func CartList(c *gin.Context,uid string){
	cacheKey := fmt.Sprintf("cart:user:%s",uid)
	//根据key 获取当前用户下购物车的商品数据
	resp,err := cache.HashAll(context.Background(),cacheKey)
	if err != nil {
		logger.Logger.Error(err)
		response.RespError(c,err)
		return
	}
	//生成购物车唯一主键
	id := common.GenerateSnowId()
	var cart models.Cart
	var totalPrice float64
	//resp map[商品id]商品数量
	for i, v := range resp {
		var cartItem = models.CartItem{}
		pid,_ := strconv.Atoi(i)
		num,_ := strconv.ParseFloat(v,64)
		//根据商品id查询商品详情
		product, err := models.GetProductByWhere("id=?", pid)
		if err != nil {
			logger.Logger.Error(err)
			response.RespError(c,"购物车商品不存在")
			return
		}
		cartItem.Cid = id
		cartItem.Product =product
		cartItem.Num, _ =strconv.Atoi(v)

		cart.CartItem =append(cart.CartItem,cartItem)
		totalPrice += num * product.ShopPrice
	}
	cart.Uid,_ =strconv.Atoi(uid)
	cart.Id = id
	response.RespData(c,"ok",cart)
}

func CartUpdateNum(c *gin.Context,item models.Item){

	cacheKey := fmt.Sprintf("cart:user:%s",item.Uid)
	//更新redis 对应 购物车中 商品的数量
	if has := cache.HashIsExists(context.Background(),cacheKey,item.Pid);!has {
		response.RespError(c,"商品不存在")
		return
	}
	var data = map[string]interface{}{
		item.Pid: item.Num,
	}
	if err := cache.HashHSet(context.Background(), cacheKey, data); err != nil {
		logger.Logger.Error(err)
		response.RespError(c, "更新失败")
		return
	}
	response.RespSuccess(c,"更新成功")
}

func DeleteCart(c *gin.Context,uid,pid string){

	cacheKey := fmt.Sprintf("cart:user:%s",uid)

	if has := cache.HashIsExists(context.Background(),cacheKey,pid);!has {
		response.RespError(c,"商品不存在")
		return
	}

	if err := cache.HashDel(context.Background(),cacheKey,pid);err != nil{
		response.RespError(c,"删除失败")
		return
	}

	response.RespSuccess(c,"删除成功")

}

func addNum(key,filed string,count int64)error{
	return cache.HashIncrBy(context.Background(),key,filed,count)
}
