package minio

import (
	"github.com/gin-gonic/gin"
	"minioServer/tool"
	minioService "minioServer/service/minio"
	"net/http"
	"path"
	"time"
)

func UploadLogFile(c *gin.Context) {
	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	form, _ := c.MultipartForm()
	files := form.File["uploadFile"]
	filePathMinio := make([]string, 1, 20)
	i := 0

	for _, file := range files {
		objectName := time.Now().Format("20060102") + "/" + file.Filename
		ext := path.Ext(file.Filename)
		err := minioService.UploadToMinio(objectName, file, "application/" + ext)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg": "上传失败",
			})

			return
		}

		filePathMinio[i] = "http://" + cfg.Minio.Endpoint + "/" + cfg.Minio.BucketName + "/" + objectName
		i = i + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{} {
			"uri": filePathMinio[0],
		},
		"msg": "上传成功",
	})
}