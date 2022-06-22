package minio_storage

import (
	"context"
	"fmt"
	"testing"
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
}
