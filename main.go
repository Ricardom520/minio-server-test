package main

import (
	"fmt"
	minio "minioServer/docker"
	"minioServer/router"
	"minioServer/tool"
)

func main() {
	cfg, _ := tool.ParseConfig("./config/app.json")

	router := router.InitRouter()
	minio.InitMinio()
	fmt.Println()
	router.Run(":" + cfg.AppPort)
}
