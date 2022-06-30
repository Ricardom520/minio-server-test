package docker

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"minioServer/tool"
)

var MinioClient *minio.Client

func InitMinio() {
	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("读取成功")

	// 初始化minio client对象
	MinioClient, err = minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})

	if err != nil {
		log.Fatalln("初始化失败")
	}

	fmt.Println("minio初始化成功")
}