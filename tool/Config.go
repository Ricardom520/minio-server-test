package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName				string			`json:"appName"`
	AppMode 			string			`json:"appMode"`
	AppHost				string			`json:"appHost"`
	AppPort				string			`json:"appPort"`
	Minio				MinioConfig		`json:"minio"`
}

type MinioConfig struct {
	Endpoint			string			`json:"endpoint"`
	AccessKeyID			string			`json:"accessKeyID"`
	SecretAccessKey		string			`json:"secretAccessKey"`
	BucketName			string			`json:"bucketName"`
	Location			string			`json:"location"`
	UseSSL				bool			`json:"use_ssl"`
}

var _cfg * Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	docoder := json.NewDecoder(reader)

	if err = docoder.Decode(&_cfg); err != nil {
		return nil, err
	}

	return _cfg, err
}
