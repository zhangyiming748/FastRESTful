package fastrestful

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// HttpDelete ...
// DELETE请求
func HttpDelete(addHeaders map[string]string, urlPath string) (body []byte, err error) {
	req, err := http.NewRequest("DELETE", urlPath, nil)
	if err != nil {
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = io.ReadAll(resp.Body)
	return
}

// HttpDeleteJson ...
// DELETE请求发送JSON数据
func HttpDeleteJson(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("DELETE", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = io.ReadAll(resp.Body)
	return
}