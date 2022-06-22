package minio_storage

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"
)

// minio测试
func TestMinioStorage(t *testing.T) {
	minioClient, err := NewMinio(&Config{
		EndPoint:        "localhost:9000",
		AccessKeyID:     "minioadmin",
		SecretAccessKey: "minioadmin",
		UserSSL:         false,
	})
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("minio connect success")
	ctx := context.Background()
	if err = CreateBucket(ctx, minioClient, "bucket1"); err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("minio bucket create success")
	err = FileUpload(ctx, minioClient, "bucket1", "xx.xlsx", "../readFile/xx.xlsx", "Application/text")
	if err != nil {
		panic(err)
	}
	requestParams := make(url.Values)
	fileUrl, err := GetFileUrl(ctx, minioClient, "bucket1", "xx.xlsx", 30*time.Second, requestParams)
	if err != nil {
		t.Logf("%v \n", err.Error())
	}
	fmt.Println(fileUrl)
}
