package oss

import "github.com/aliyun/aliyun-sts-go-sdk/sts"

type AliOss struct {
	Config
}

func NewAliOss(config Config) *AliOss {
	if config.STSConfig.ExpireDur == 0 {
		config.STSConfig.ExpireDur = 3600
	}
	return &AliOss{
		Config: config,
	}
}

func (a *AliOss) AssumeRole() (response *sts.Credentials, err error) {
	stsClient := sts.NewClient(a.AccessKeyID, a.AccessKeySecret, a.RoleArn, "my-project")
	resp, err := stsClient.AssumeRole(a.ExpireDur)
	if err != nil {
		return response, err
	}
	return &resp.Credentials, nil
}
