package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"go-background-survey/handlers"
)

func InitializeRoutes(group *gin.RouterGroup, router *gin.Engine) {
	router.GET("/health", handlers.GetTest)

	//加载静态文件
	router.Use(static.Serve("/static", static.LocalFile("app/dist/static", true)))
	router.Use(static.Serve("/", static.LocalFile("app/dist/", true)))
	router.LoadHTMLGlob("app/dist/index.html")

}
