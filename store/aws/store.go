package aws

import (
	"fmt"

	"github.com/sunflower10086/cloud-station/store"
)

var (
	_ store.Uploader = &AwsOssStore{}
)

type AwsOssStore struct {
}

func (t *AwsOssStore) Upload(bucketName, objectKey, fileName string) error {

	return nil
}

func NewAwsOssStore() (*AwsOssStore, error) {
	fmt.Println("aws function not impl")
	return nil, nil
}
