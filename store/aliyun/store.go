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
	client *oss.Client
}

func NewAliOssStore(ossEndpoint, accessKeyId, accessKeySecret string) (*AliOssStore, error) {
	client, err := oss.New(ossEndpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client}, nil
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
	fmt.Printf("文件下载url %s\n", downloadUrl)
	fmt.Println("下载链接有效期一天，请在一天内下载")
	return nil
}
