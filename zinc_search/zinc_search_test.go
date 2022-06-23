package zinc_search

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var query = `{
        "search_type": "alldocuments",
        "query":
        {
            "term": "DEMTSCHENKO",
            "start_time": "2021-06-02T14:28:31.894Z",
            "end_time": "2021-12-02T15:28:31.894Z"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`

func TestZincSearch(t *testing.T) {
	data, err := ReadFile("./olympics.ndjson")
	if err != nil {
		t.Log("read file error")
	}
	fmt.Println("data:", data)
	ctx := context.Background()
	head := make(map[string]interface{}, 1)
	head["Authorization"] = "Basic YWRtaW46Q29tcGxleHBhc3MjMTIz"
	// 创建索引
	resp, err := CreateIndex(ctx, &http.Client{}, "http://localhost:4080/api/index", head, strings.NewReader(data))
	if err != nil {
		t.Log(err)
	}
	fmt.Println(resp.Status)
	// 创建文档
	resp, err = CreateDocument(ctx, &http.Client{}, "http://localhost:4080/api/"+"article"+"/_doc", head, strings.NewReader(`{"title": "golang"}`))
	if err != nil {
		t.Log(err)
	}
	fmt.Println(resp.Status)
	// 展示所有
	resp, err = ListIndex(ctx, &http.Client{}, "http://localhost:4080/api/index", head, strings.NewReader(data))
	if err != nil {
		t.Log(err)
	}
	fmt.Println(resp.Status)
	// 搜索索引
	resp, err = Search(ctx, &http.Client{}, "http://localhost:4080/api/"+"article"+"/_search", head, strings.NewReader(query))
	if err != nil {
		t.Log(err)
	}
	respBodyByte, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBodyByte))
	// 删除
	resp, err = DeleteIndex(ctx, &http.Client{}, "http://localhost:4080/api/index/article", head, strings.NewReader(data))
	if err != nil {
		t.Log(err)
	}
	fmt.Println(resp.Status)
}

func ReadFile(filename string) (string, error) {
	fileContentByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(fileContentByte), nil
}
