package oss

type Config struct {
	Endpoint        string
	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	CdnDomain       string
	STSConfig
}

type STSConfig struct {
	RoleArn   string
	ExpireDur uint
}
