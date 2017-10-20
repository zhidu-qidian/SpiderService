package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	S "workspace/SpiderService/storage"
)

const (
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"
)

var HttpClient = http.Client{
	Timeout: 60 * time.Second,
}

func getUniqueID(buffer []byte) (id string, err error) {
	h := sha256.New()
	if _, err = io.Copy(h, bytes.NewReader(buffer)); err != nil {
		return
	}
	id = hex.EncodeToString(h.Sum(nil))
	return id, err
}

func UploadImage(url, refer string) (u string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", UserAgent)
	if len(refer) > 10 {
		req.Header.Set("Referer", refer)
	}
	r, err := HttpClient.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	id, err := getUniqueID(buffer)
	if err != nil {
		return
	}
	suffix := ""
	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	switch contentType {
	case "image/jpeg":
		suffix = "jpg"
	case "image/webp":
		suffix = "webp"
	case "image/png":
		suffix = "png"
	case "image/gif":
		suffix = "gif"
	case "image/bmp":
		suffix = "bmp"
	}
	if len(suffix) == 0 {
		err = fmt.Errorf("Can't download image: %s", contentType)
		return
	}
	objectKey := fmt.Sprintf("%s.%s", id, suffix)
	bucketName := S.ImageBucketName
	bucket, err := S.GetBucket(bucketName)
	if err != nil {
		return
	}
	err = bucket.PutObject(objectKey, bytes.NewReader(buffer))
	if err != nil {
		return
	}
	domain := S.BucketNameDomainMap[bucketName]
	u = fmt.Sprintf("%s/%s", domain, objectKey)
	return u, err
}

func DeleteImage(objectKey string) (err error) {
	bucketName := S.ImageBucketName
	bucket, err := S.GetBucket(bucketName)
	if err != nil {
		return
	}
	err = bucket.DeleteObject(objectKey)
	return err
}
