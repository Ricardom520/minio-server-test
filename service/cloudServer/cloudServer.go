package cloudServer

import (
	"fmt"
	"context"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	minioDocker "minioServer/docker"
	"minioServer/tool"
	"os"
	"time"
)

func UploadMinio(msg, timestamp string) (uri string, err error){
	cfg, err := tool.ParseConfig("./config/app.json")
	var path string
	var log string

	path = "./logs"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		os.Chmod(path, 0777)
	}

	today := time.Now().Format("2006-01-02")
	log = path + "/" + today + ".log"
	objectName := today + "/" + timestamp + ".log"

	err1 := ioutil.WriteFile(log, []byte(msg), 0777)

	if err1 != nil {
		return "", err1
	}

	src, err1 := os.Open(log)

	if err1 != nil {
		panic(err1.Error())
	}



	// 使用PutObject上传文件
	_, err2 := minioDocker.MinioClient.PutObject(context.Background(), cfg.Minio.BucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: ".log"})

	if err2 != nil {
		return "", err2
	}

	src.Close()

	err3 := os.Remove(log)               //删除文件test.txt

	if err3 != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}

	return "http://" + "172.29.139.32:14942" + "/" + cfg.Minio.BucketName + "/" + objectName, nil
}