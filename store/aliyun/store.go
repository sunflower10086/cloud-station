package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"github.com/sunflower10086/cloud-station/store"
)

var (
	_    store.Uploader = &AliOssStore{}
	Conf                = new(config)
)

type AliOssStore struct {
	client *oss.Client
}

type config struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
}

// 读取配置文件
func readEnvFile() (*config, error) {
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AddConfigPath("./etc")     // 查找配置文件所在的目录
	viper.AddConfigPath("../../etc") // 查找配置文件所在的目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			fmt.Println("配置文件未找到")
			return nil, err
		} else {
			// 其他错误
			fmt.Println("加载配置文件错误")
			return nil, err
		}
	}

	var conf config

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &conf, nil
}

func (c *config) validate() error {
	if c.AccessKeyId == "" || c.AccessKeySecret == "" || c.OssEndpoint == "" {
		return fmt.Errorf("access_key, secret_key, end_pointe has one empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	conf, err := readEnvFile()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return NewAliOssStore(conf)
}

func NewAliOssStore(conf *config) (*AliOssStore, error) {
	if err := conf.validate(); err != nil {
		return nil, err
	}

	client, err := oss.New(conf.OssEndpoint, conf.AccessKeyId, conf.AccessKeySecret)
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
