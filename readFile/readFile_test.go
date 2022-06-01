package readFile

import (
	"fmt"
	"testing"
)

func TestReadFileCsv(t *testing.T) {
	readFile := ReadFile{
		"xx.csv",
		"csv",
	}
	res, err := readFile.ReadFile()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(res)
}

func TestReadFileXls(t *testing.T) {
	readFile := ReadFile{
		"xx.xlsx",
		"excel",
	}
	res, err := readFile.ReadFile()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(res)
}
