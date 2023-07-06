package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sunflower10086/cloud-station/store"
)

var (
	_ store.Uploader = &AliOssStore{}
)

type AliOssStore struct {
	client      *oss.Client
	aliOssStore *Config
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	conf, err := ReadEnvFile()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return NewAliOssStore(conf)
}

func NewAliOssStore(conf *Config) (*AliOssStore, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	client, err := oss.New(conf.OssEndpoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client, aliOssStore: conf}, nil
}

func (a *AliOssStore) Upload(bucketName, objectKey, fileName string) error {
	// 2.获得我们的bucket对象
	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 3.上传文件
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// 4. 打印下载链接
	downloadUrl, err := bucket.SignURL(fileName, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Println("上传云商: 阿里云[oss-cn-beijing.aliyuncs.com]")
	fmt.Printf("上传用户: [%s]\n", a.aliOssStore.OssEndpoint)
	fmt.Printf("文件下载url: [%s]\n\n", downloadUrl)
	fmt.Println("注意:下载链接有效期一天，请在一天内下载")
	return nil
}
