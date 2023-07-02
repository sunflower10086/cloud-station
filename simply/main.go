package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var Conf = new(config)

var (
	// 程序内置
	accessKeyId     = Conf.AccessKeyId
	accessKeySecret = Conf.AccessKeySecret
	ossEndpoint     = Conf.OssEndpoint

	// 默认配置
	bucketName = "lz-devcloud-station"

	// 命令行上传
	upLoadFile = ""

	help = false
)

type config struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
}

func init() {
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AddConfigPath("./etc")  // 查找配置文件所在的目录
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

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}

	accessKeyId = Conf.AccessKeyId
	accessKeySecret = Conf.AccessKeySecret
	ossEndpoint = Conf.OssEndpoint
}

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
	if err := bucket.PutObjectFromFile(filePath, filePath); err != nil {
		return err
	}

	// 4. 打印下载链接
	downloadUrl, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载url %s\n", downloadUrl)
	fmt.Println("下载链接有效期一天，请在一天内下载")
	return nil
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
	flag.BoolVar(&help, "h", false, "打印帮助信息")

	flag.StringVar(&upLoadFile, "f", "", "请输入文件的名称")
	flag.Parse()

	// 判断CLI 是否需要打印help信息
	if help {
		usage()
		os.Exit(0)
	}
}

// 打印使用说明
func usage() {
	// 1.打印一些描述信息
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	// 2.打印有哪些参数可以使用，像-f
	flag.PrintDefaults()
}

func main() {
	// 1.参数校验
	loadParams()
	// 2.参数检验
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常 %s\n", err.Error())
		return
	}

	// 3.上传文件
	if err := upLoad(upLoadFile); err != nil {
		fmt.Printf("上传文件错误 %s\n", err.Error())
		return
	}

	fmt.Printf("文件 %s 上传成功 \n", upLoadFile)
}
