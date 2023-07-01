package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 程序内置
	accessKeyId     = "xx"
	accessKeySecret = "xx"
	ossEndpoint     = "oss-cn-beijing.aliyuncs.com"

	// 默认配置
	bucketName = "lz-devcloud-station"

	// 命令行上传
	upLoadFile = ""
)

func upLoad(filePath string) error {
	// 1.初始化oss.client
	client, err := oss.New(ossEndpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return err
	}

	// 2.获得我们的bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 3.上传文件
	return bucket.PutObjectFromFile(filePath, filePath)
}

// 参数合法性校验
func validate() error {
	if accessKeyId == "" || accessKeySecret == "" || ossEndpoint == "" {
		return fmt.Errorf("access_key, secret_key, end_pointe has one empty")
	}

	if upLoadFile == "" {
		return fmt.Errorf("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.StringVar(&upLoadFile, "f", "", "请输入文件的名称")
	flag.Parse()
}

func main() {
	// 参数校验
	loadParams()
	// 参数检验
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常 %s\n", err.Error())
		return
	}

	if err := upLoad(upLoadFile); err != nil {
		fmt.Printf("上传文件错误 %s\n", err.Error())
	}

	fmt.Printf("文件 %s 上传成功 \n", upLoadFile)
}
