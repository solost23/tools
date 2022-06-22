package minio_storage

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go"
)

func NewMinio(config *Config) (client *minio.Client, err error) {
	client, err = minio.New(config.EndPoint, config.SecretAccessKey, config.AccessKeyID, config.UserSSL)
	if err != nil {
		log.Panicf("minio connect failed: %v \n", err.Error())
	}
	return client, err
}

func CreateBucket(_ context.Context, minio *minio.Client, bucketName string) (err error) {
	exists, err := minio.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// 不存在，则创建
	err = minio.MakeBucket(bucketName, "")
	if err != nil {
		return err
	}
	return nil
}

func FileUpload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, filePath string, contextType string) (err error) {
	_, err = minioClient.FPutObjectWithContext(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		return err
	}
	return nil
}

func FileDownload(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, filePath string) (err error) {
	err = minioClient.FGetObjectWithContext(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	return err
}

// 生成对象的url
func GetFileUrl(_ context.Context, minioClient *minio.Client, bucketName string, objectName string, expire time.Duration, request url.Values) (url string, err error) {
	urlObject, err := minioClient.PresignedGetObject(bucketName, objectName, expire, request)
	if err != nil {
		return "", err
	}
	return urlObject.String(), nil
}
