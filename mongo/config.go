package mongo

import "time"

type Config struct {
	Hosts      []string
	AuthSource string
	UserName   string
	Password   string
	Database   string
	Timeout    time.Duration // second
}
