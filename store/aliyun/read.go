package aliyun

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
}

func (c *Config) Validate() error {
	if c.AccessKeyId == "" || c.AccessKeySecret == "" || c.OssEndpoint == "" {
		return fmt.Errorf("access_key, secret_key, end_pointe has one empty")
	}
	return nil
}

// 读取配置文件
func ReadEnvFile() (*Config, error) {
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

	var conf Config

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &conf, nil
}
