package  main

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"go-background-survey/alipayMgr"
	"go-background-survey/routes"
)

var router *gin.Engine

func main(){
	var(
		//configFilePath string
		logger seelog.LoggerInterface
		err error
	)
	fmt.Println("GoPay Version: ", gopay.Version)
	logger,err =seelog.LoggerFromConfigAsFile("config/seelog.xml")
	if err != nil {
		goto ERR
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()
	alipayMgr.InitAlipayMgr(false)
	router=gin.New()
	routes.InitializeRoutes(router.Group("/v1"),router)
	seelog.Info("[Serve]","[Run]", "服务开启")
	router.Run(":8753")
	return


ERR:
	return
	seelog.Error("[Error]","[Message]", err)
}


