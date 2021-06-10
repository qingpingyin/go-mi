package carousel

import (
	"MI/models"
	"MI/pkg/logger"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
)

func Carousel(c  *gin.Context){
	carousels, err := models.GetCarouselList()
	if err != nil {
		logger.Logger.Info(err)
	}
	response.RespData(c,"",carousels)
}

