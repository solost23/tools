package md5

import (
	"fmt"
	"testing"
)

func TestNewMd5(t *testing.T) {
	md5Str := NewMd5("123", "hhh", "qqq")
	fmt.Println(md5Str)
}
