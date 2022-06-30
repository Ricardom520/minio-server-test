package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	minioDocker "minioServer/docker"
	"minioServer/tool"
)

func UploadToMinio(objectName string, file *multipart.FileHeader, contentType string) error {
	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	// 创建一个叫xxx的桶
	err = minioDocker.MinioClient.MakeBucket(context.Background(), cfg.Minio.BucketName, minio.MakeBucketOptions{
		Region: cfg.Minio.Location,
	})

	if err != nil {
		// 检查存储桶是否已经存在
		exists, err := minioDocker.MinioClient.BucketExists(context.Background(), cfg.Minio.BucketName)

		if err == nil && exists {
			log.Printf("Bucket: %s is already exist\n", cfg.Minio.BucketName)
		} else {
			log.Println("不存在")

			return err
		}
	}

	log.Printf("Successfully created bucket: %s\n", cfg.Minio.BucketName)

	src, err1 := file.Open()

	if err1 != nil {
		panic(err1.Error())
	}

	defer src.Close()

	// 使用PutObject上传文件
	_, err2 := minioDocker.MinioClient.PutObject(context.Background(), cfg.Minio.BucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: contentType})

	if err2 != nil {
		return err2
	}

	return nil
}