package aliyun

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunflower10086/cloud-station/store"
)

var (
	uploader store.Uploader
)

// var (
// 	// 程序内置
// 	accessKeyId     = Conf.AccessKeyId
// 	accessKeySecret = Conf.AccessKeySecret
// 	ossEndpoint     = Conf.OssEndpoint
// 	bucketName      = Conf.BucketName
// )

func init() {
	readEnvFile()
	aliOssStore, err := NewDefaultAliOssStore()
	if err != nil {
		return
	}

	uploader = aliOssStore
}

// Aliyun Oss Store 测试用例
func TestUpLoad(t *testing.T) {

	should := assert.New(t)
	err := uploader.Upload(Conf.BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		t.Log("upload ok")
	}
}
