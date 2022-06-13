package readFile

// 读取文件并将字段去空格

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ReadFile struct {
	FileName string
	// csv or excel
	FileType string
}

func (r *ReadFile) ReadFile() (res [][]string, err error) {
	if r.FileType == "excel" {
		return r.readFileExcel()
	} else if r.FileType == "csv" {
		return r.readFileCsv()
	}
	return res, errors.New("不支持此文件类型")
}

func (r *ReadFile) readFileCsv() (res [][]string, err error) {
	//pathAbs, err := os.Getwd()
	//if err != nil {
	//	return res, err
	//}
	//filePath := fmt.Sprintf("%s/%s", pathAbs, r.FileName)
	fp, err := os.Open(r.FileName)
	if err != nil {
		return
	}
	defer func() {
		fp.Close()
	}()
	reader := csv.NewReader(fp)
	// 针对大文件，一行一行读取
	for {
		row, err := reader.Read()
		if err != nil && err != io.EOF {
			return res, err
		}
		if err == io.EOF {
			break
		}
		res = append(res, row)
	}
	return r.trim(res)
}

func (r *ReadFile) readFileExcel() (res [][]string, err error) {
	//pathAbs, err := os.Getwd()
	//if err != nil {
	//	return res, err
	//}
	//filePath := fmt.Sprintf("%s/%s", pathAbs, r.FileName)
	fp, err := excelize.OpenFile(r.FileName)
	if err != nil {
		return res, err
	}
	defer func() {
		fp.Close()
	}()
	rows, err := fp.GetRows("Sheet1")
	if err != nil {
		return res, err
	}
	for _, value := range rows {
		res = append(res, value)
	}
	return r.trim(res)
}

// 为数据去除空格
func (r *ReadFile) trim(param [][]string) (res [][]string, err error) {
	for _, item := range param {
		var fields []string
		for _, field := range item {
			fields = append(fields, strings.Trim(field, " "))
		}
		res = append(res, fields)
	}
	return res, nil
}
