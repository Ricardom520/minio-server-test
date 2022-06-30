package router

import (
	minioApi "minioServer/api/minio"
	"minioServer/api/cloudServer"
	"github.com/gin-gonic/gin"
	"minioServer/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 设置全局跨域访问
	router.Use(handler.CrosHandler())

	v1 := router.Group("/v1")
	{
		v1.POST("/cloudServer/:action", cloudServer.CsHandle)
		v1.POST("/uploadLogFile", minioApi.UploadLogFile)

		v1.OPTIONS("/cloudServer/:action", cloudServer.CsHandle)
	}

	return  router
}