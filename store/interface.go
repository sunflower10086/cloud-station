package store

type Uploader interface {
	Upload(bucketName, objectKey, fileName string) error
}
