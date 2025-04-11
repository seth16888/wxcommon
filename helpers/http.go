package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	wx "github.com/seth16888/wxcommon/error"
)

// BuildRequestBody 构建请求体
func BuildRequestBody[T any](body T) (io.Reader, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("marshal body error")
	}

	bodyReader := bytes.NewReader(bodyBytes)

	return bodyReader, nil
}

// ReadResponseBody 读取http response body
func ReadResponseBody(r *http.Response, err error) ([]byte, error) {
	if err != nil { // network error
		return nil, fmt.Errorf("network error:%s", err.Error())
	}

	// 必须先判断err, 否则会panic, r.Body == nil
	defer r.Body.Close()

	responseBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error:%s", err.Error())
	}
	return responseBytes, nil
}

// BuildHttpResponse 构建代理响应结果
func BuildHttpResponse[T any](resp *http.Response, err error) (r *T, e *wx.WXError) {
	rtDataBytes, err := ReadResponseBody(resp, err)
	if err != nil {
		return nil, &wx.WXError{ErrCode: 500, ErrMsg: err.Error()}
	}

	var result *T
	if err := json.Unmarshal(rtDataBytes, &result); err != nil {
		return nil, &wx.WXError{ErrCode: 500, ErrMsg: "unmarshal wx server response error"}
	}

	return result, nil
}

// BuildHttpResponse2 构建代理响应结果
func BuildHttpResponse2[T any](resp *http.Response, err error) (r *T, e error) {
	rtDataBytes, err := ReadResponseBody(resp, err)
	if err != nil {
		return nil, err
	}

	var result *T
	if err := json.Unmarshal(rtDataBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}
