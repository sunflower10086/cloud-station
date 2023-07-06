package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sunflower10086/cloud-station/store"
	"github.com/sunflower10086/cloud-station/store/aliyun"
	"github.com/sunflower10086/cloud-station/store/aws"
	"github.com/sunflower10086/cloud-station/store/tx"
)

var (
	ossProvider     string
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string
	uploadFile      string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件中转站",
	Short:   "upload 文件中转站",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)

		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Config{
				AccessKeyId:     accessKeyId,
				AccessKeySecret: accessKeySecret,
				OssEndpoint:     endpoint,
				BucketName:      bucketName,
			})
		case "tx":
			uploader, err = tx.NewTxOssStore()
		case "aws":
			uploader, err = aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage provider")
		}

		if err != nil {
			return nil
		}
		// 使用uploader上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "ossProvider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&endpoint, "Endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&accessKeyId, "AccessKeyId", "k", "", "oss storage provider AccessKeyId")
	f.StringVarP(&accessKeySecret, "AccessKeySecret", "s", "", "oss storage provider AccessKeySecret")
	f.StringVarP(&bucketName, "BucketName", "b", "lz-devcloud-station", "oss storage provider BucketName")
	f.StringVarP(&uploadFile, "uploadFile", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
