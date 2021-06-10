package routers

import (
	"MI/api/carousel"
	"MI/api/cate"
	"MI/api/product"
	"MI/api/user"
	"MI/middleware"
	"MI/pkg/setting"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(setting.ApplicationConf.Env)

	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.NoFound(),
		)
	r.GET("/", func(context *gin.Context) {
		ip := context.ClientIP()
		response.RespSuccess(context,ip)
	})


	r.POST("/singUpByMobile",user.SingUpByMobile)
	r.POST("/sendSms",user.SendSms)
	r.POST("/login",user.Login)
	r.GET("/carousel",carousel.Carousel)
	r.GET("/cate",cate.Cate)
	r.GET("/product",product.Product)
	r.GET("/productDetail",product.GetProductDetail)

	r.Use(middleware.JWTAuthMiddleware())
	{
		r.GET("/userInfo",user.UserInfo)
	}

	return r
}
