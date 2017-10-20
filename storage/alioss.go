package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"workspace/SpiderService/config"
)

var (
	ali                 *oss.Client
	BucketNameDomainMap = map[string]string{
		ImageBucketName: "http://pro-pic.deeporiginalx.com",
	}
)

const (
	ImageBucketName = "bdp-images"
)

func init() {
	var err error
	endpoint, accessID, accessKey := config.C.AliOss.Endpoint, config.C.AliOss.AccessID, config.C.AliOss.AccessKey
	ali, err = oss.New(endpoint, accessID, accessKey)
	if err != nil {
		panic(err)
	}
}

func GetBucket(name string) (bucket *oss.Bucket, err error) {
	bucket, err = ali.Bucket(name)
	return bucket, err
}
