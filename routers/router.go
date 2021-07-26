package routers

import (
	"MI/api/address"
	"MI/api/carousel"
	"MI/api/cart"
	"MI/api/cate"
	"MI/api/collect"
	"MI/api/order"
	"MI/api/product"
	"MI/api/upload"
	"MI/api/user"
	"MI/middleware"
	"MI/pkg/setting"
	"MI/utils/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(setting.ApplicationConf.Env)

	r := gin.New()
	r.Use(
		gin.Recovery(),
		cors.Default(),
		middleware.Logger(),
		middleware.NoFound(),
		)
	{
		r.GET("/", func(context *gin.Context) {
			ip := context.ClientIP()
			response.RespSuccess(context, ip)
		})
		r.POST("/singUpByMobile", user.SingUpByMobile)
		r.POST("/sendSms", user.SendSms)
		r.POST("/checkCode",user.CheckCode)
		r.PUT("/forgetPassword",user.ForgetPassword)
		r.POST("/login", user.Login)
		r.GET("/validate/email",user.ValidateEmail)
		r.GET("/carousel", carousel.Carousel)
		r.GET("/cate", cate.Cate)
		r.GET("/product", product.Product)
		r.GET("/productDetail", product.GetProductDetail)

	}

	r.Use(middleware.JWTAuthMiddleware())
	{
		//user
		r.GET("/userInfo",user.UserInfo)
		r.POST("/logout",user.Logout)
		r.PUT("/updateUser",user.UpdateUserInfo)
		//email
		r.POST("/bindEmail",user.BindEmail)
		//cart
		r.POST("/addCart",cart.AddCart)
		r.GET("/cartList",cart.CartList)
		r.PUT("/updateCartNum/:uid/:pid/:num",cart.CartUpdateNum)
		r.DELETE("/deleteCart/:uid/:pid",cart.DeleteCart)
		//collect
		r.GET("/collect",collect.List)
		r.POST("/collect",collect.CreateCollect)
		r.DELETE("/collect/:uid/:pid",collect.DelCollect)
		//address
		r.GET("/address",address.Address)
		r.POST("/address",address.CreateAddress)
		r.DELETE("/address/:id",address.DeleteAddres)
		//order
		r.GET("/order",order.OrderList)
		r.POST("/order",order.CreateOrder)
		//upload
		r.POST("/avatarUpload",upload.AvatarUpload)
	}

	return r
}


