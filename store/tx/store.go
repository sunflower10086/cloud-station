package tx

import (
	"fmt"

	"github.com/sunflower10086/cloud-station/store"
)

var (
	_ store.Uploader = &TxOssStore{}
)

type TxOssStore struct {
}

func (t *TxOssStore) Upload(bucketName, objectKey, fileName string) error {

	return nil
}

func NewTxOssStore() (*TxOssStore, error) {
	fmt.Println("tx function not impl")
	return nil, nil
}
