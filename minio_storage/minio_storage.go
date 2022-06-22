package minio_storage

import (
	"context"
	"log"

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
