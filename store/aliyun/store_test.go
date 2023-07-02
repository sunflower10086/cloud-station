package aliyun

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/sunflower10086/cloud-station/store"
)

var (
	uploader store.Uploader
	Conf     = new(config)
)

var (
	// 程序内置
	accessKeyId     = Conf.AccessKeyId
	accessKeySecret = Conf.AccessKeySecret
	ossEndpoint     = Conf.OssEndpoint
	bucketName      = Conf.BucketName
)

type config struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
}

func readEnvFile() {
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AddConfigPath("./etc")     // 查找配置文件所在的目录
	viper.AddConfigPath("../../etc") // 查找配置文件所在的目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			fmt.Println("配置文件未找到")
		} else {
			// 其他错误
			fmt.Println("加载配置文件错误")
		}
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}

	accessKeyId = Conf.AccessKeyId
	accessKeySecret = Conf.AccessKeySecret
	ossEndpoint = Conf.OssEndpoint
	bucketName = Conf.BucketName
}

func init() {
	readEnvFile()
	aliOssStore, err := NewAliOssStore(ossEndpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}

	uploader = aliOssStore
}

// Aliyun Oss Store 测试用例
func TestUpLoad(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(bucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		t.Log("upload ok")
	}
}
