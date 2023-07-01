package example_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var client *oss.Client

var Conf = new(config)

type config struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
}

// 测试oss的基本使用
func init() {
	ReadEnvFile()
	c, err := oss.New(Conf.OssEndpoint, Conf.AccessKeyId, Conf.AccessKeySecret)

	if err != nil {
		panic(err)
	}
	client = c
}

func ReadEnvFile() {
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AddConfigPath("../etc") // 查找配置文件所在的目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			fmt.Println("配置文件未找到")
		} else {
			// 其他错误
			fmt.Println("加载配置文件错误")
		}
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println(err)
	}
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

	bucket, err := client.Bucket(Conf.BucketName)
	if err != nil {
		t.Log(err)
	}

	// 上传文件到bucket
	// 云商会根据key的路径结构自动的帮我们创建目录
	// 第一个参数我们文件放在bucket中的对象的名字(包含路径)   myDir/test.go 自动创建myDir
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}
