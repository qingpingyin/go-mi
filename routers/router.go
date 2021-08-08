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
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(setting.ApplicationConf.Env)

	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.CorsMiddleware(),
		middleware.Logger(),
		middleware.NoFound(),
		)
		v1 := r.Group("/api")
		{
			v1.GET("/", func(context *gin.Context) {
				ip := context.ClientIP()
				response.RespSuccess(context, ip)
			})
			v1.POST("/singUpByMobile", user.SingUpByMobile)
			v1.POST("/sendSms", user.SendSms)
			v1.POST("/checkCode", user.CheckCode)
			v1.PUT("/forgetPassword", user.ForgetPassword)
			v1.POST("/login", user.Login)
			v1.GET("/validate/email", user.ValidateEmail)
			v1.GET("/carousel", carousel.Carousel)
			v1.GET("/cate", cate.Cate)
			v1.GET("/search", product.GetProductBySearch)
			v1.GET("/product", product.Product)
			v1.GET("/productDetail", product.GetProductDetail)
		}
		v1.Use(middleware.JWTAuthMiddleware())
		{
			//user
			v1.GET("/userInfo", user.UserInfo)
			v1.POST("/logout", user.Logout)
			v1.PUT("/updateUser", user.UpdateUserInfo)
			//email
			v1.POST("/bindEmail", user.BindEmail)
			//cart
			v1.POST("/addCart", cart.AddCart)
			v1.GET("/cartList", cart.CartList)
			v1.PUT("/updateCartNum/:uid/:pid/:num", cart.CartUpdateNum)
			v1.DELETE("/deleteCart/:uid/:pid", cart.DeleteCart)
			//collect
			v1.GET("/collect", collect.List)
			v1.POST("/collect", collect.CreateCollect)
			v1.DELETE("/collect/:uid/:pid", collect.DelCollect)
			v1.GET("/getCollectCount",collect.CollectCount)
			//address
			v1.GET("/address", address.Address)
			v1.GET("/addressById", address.GetAddressById)
			v1.POST("/address", address.CreateAddress)
			v1.DELETE("/address/:id", address.DeleteAddres)
			//order
			v1.GET("/order", order.OrderList)
			v1.GET("/orderById", order.GetOrderById)
			v1.POST("/createOrder", order.CreateOrder)
			v1.GET("/getOrderCount",order.GetOrderCountBy)
			//upload
			v1.POST("/avatarUpload", upload.AvatarUpload)
		}
	return r
}


