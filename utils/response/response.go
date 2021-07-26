package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespError(c *gin.Context,msg interface{}){
	c.JSON(http.StatusOK,gin.H{
		"status":400,
		"msg":msg,
	})
}

func RespValidatorError(c *gin.Context,msg interface{}){
	c.JSON(http.StatusOK, gin.H{
		"status": 400,
		"msg":  msg,
	})
}
func RespData(c *gin.Context,msg string,data interface{}){
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":  msg,
		"data":data,
	})
}
func RespSuccess(c *gin.Context,msg interface{}){
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":  msg,
	})
}