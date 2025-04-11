package error

import "fmt"

type WXError struct {
	ErrCode int    `json:"errcode" xml:"errcode"`
	ErrMsg  string `json:"errmsg" xml:"errmsg"`
}

func (e *WXError) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrCode, e.ErrMsg)
}
