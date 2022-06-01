package oss

import (
	"fmt"
	"testing"
)

func TestAliOss_AssumeRole(t *testing.T) {
	client := NewAliOss(Config{
		Endpoint:        "",
		Bucket:          "",
		AccessKeyID:     "",
		AccessKeySecret: "",
		STSConfig: STSConfig{
			RoleArn:   "",
			ExpireDur: 3600,
		},
	})
	resp, err := client.AssumeRole()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(resp)
}
