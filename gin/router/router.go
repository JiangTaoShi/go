package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.gin/api/sample"
	_ "go.gin/docs"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	sample := sample.New()
	r.POST("/test01", sample.Test01)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
