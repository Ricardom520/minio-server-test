package cloudServer

import (
	"github.com/gin-gonic/gin"
	"minioServer/service/cloudServer"
	"net/http"
)

func CsHandle(c *gin.Context) {
	action := c.Param("action")

	switch action {
	case "logBlock":
		logBlock(c)
	}
}

func logBlock(c *gin.Context) {
	type LogBlockData struct {
		Msg			string		`json:"msg"`
		Timestamp 	string		`json:"timestamp"`
		Level		string		`json:"level"`
	}

	logBlockData := LogBlockData{}

	c.ShouldBind(&logBlockData)

	if &logBlockData == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "参数错误",
		})

		return
	}

	uri, err := cloudServer.UploadMinio(logBlockData.Msg, logBlockData.Timestamp)

	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{} {
			"uri": uri,
		},
		"msg": "上传成功",
	})
}