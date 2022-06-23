package zinc_search

import (
	"context"
	"net/http"
	"strings"
)

func CreateIndex(ctx context.Context, client *http.Client, url string, header map[string]interface{}, data *strings.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequestWithContext(ctx, "POST", url, data)
	if err != nil {
		return resp, err
	}
	// 循环赋值
	for key, value := range header {
		request.Header.Set(key, value.(string))
	}
	resp, err = client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func DeleteIndex(ctx context.Context, client *http.Client, url string, header map[string]interface{}, data *strings.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequestWithContext(ctx, "DELETE", url, data)
	if err != nil {
		return resp, err
	}
	// 循环赋值
	for key, value := range header {
		request.Header.Set(key, value.(string))
	}
	resp, err = client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func ListIndex(ctx context.Context, client *http.Client, url string, header map[string]interface{}, data *strings.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequestWithContext(ctx, "GET", url, data)
	if err != nil {
		return resp, err
	}
	// 循环赋值
	for key, value := range header {
		request.Header.Set(key, value.(string))
	}
	resp, err = client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CreateDocument(ctx context.Context, client *http.Client, url string, header map[string]interface{}, data *strings.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequestWithContext(ctx, "POST", url, data)
	// 循环赋值
	for key, value := range header {
		request.Header.Set(key, value.(string))
	}
	resp, err = client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func Search(ctx context.Context, client *http.Client, url string, header map[string]interface{}, data *strings.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequestWithContext(ctx, "POST", url, data)
	// 循环赋值
	for key, value := range header {
		request.Header.Set(key, value.(string))
	}
	resp, err = client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
