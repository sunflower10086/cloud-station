package example_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

//var (
//	AccessKeyId     = os.Getenv("ALI_AK")
//	AccessKeySecret = os.Getenv("ALI_SK")
//	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
//	BucketName      = os.Getenv("ALI_BUCKET_NAME")
//)

var (
	AccessKeyId     = "LTAI5tFn7XFC9Lkv1vSJ55SA"
	AccessKeySecret = "2YZknyvGJGCjaFA4lX9GdbnSmW4ZwY"
	OssEndpoint     = "oss-cn-beijing.aliyuncs.com"
	BucketName      = "lz-devcloud-station"
)

// 测试oss的基本使用
func init() {

	c, err := oss.New(OssEndpoint, AccessKeyId, AccessKeySecret)
	fmt.Println(OssEndpoint)
	if err != nil {
		panic(err)
	}
	client = c
}

func TestBucketList(t *testing.T) {

	lsRes, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUpLoadFile(t *testing.T) {

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}
}
