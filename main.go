package main

import (
	"MI/models"
	"MI/pkg/cache"
	"MI/pkg/logger"
	"MI/pkg/setting"
	"MI/pkg/validate"
	"MI/routers"
	"flag"
	"fmt"
)

func main(){

	load()
	router := routers.InitRouter()
	panic(router.Run(fmt.Sprintf("%s:%d", setting.ApplicationConf.Host, setting.ApplicationConf.Port)))
}

func load(){
	flag.Parse()
	setting.Setup()
	models.Setup()
	cache.SetUp()
	logger.Setup()
	validate.InitValidate()
}
