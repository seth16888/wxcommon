package mp

// R API给客户端的统一结果
type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func (r *R[T]) GetHttpCode() int {
	rt := 200
	if r.Code <= 1024 && r.Code >= 200 {
		rt = r.Code
	}
	return rt
}

// WXError WX 返回的错误信息
type WXError struct {
	ErrCode int    `json:"errcode,omitempty" xml:"errcode"`
	ErrMsg  string `json:"errmsg,omitempty" xml:"errmsg"`
}

// AccessTokenRes 从WX获取access_token的返回结果
type AccessTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Ticket 获取二维码返回结果
type Ticket struct {
	Ticket        string `json:"ticket"`
	URL           string `json:"url"` // URL 解析后的网址，可根据URL自行生成二维码
	ExpireSeconds int64  `json:"expire_seconds"`
}

// GenShortenReq 将不超过4KB的长信息转成短key
type GenShortenReq struct {
	// LongData 需要转换的长信息，不超过4KB
	LongData string `json:"long_data" binding:"required" msg:"long_data required"`
	// 过期秒数，最大值为2592000（即30天），默认为2592000
	ExpireSeconds int64 `json:"expire_seconds"`
}

// FetchShortenRes 获取短key的返回结果
type FetchShortenRes struct {
	// 长信息
	LongData string `json:"long_data"`
	// 创建的时间戳
	CreateTime int64 `json:"create_time"`
	// 剩余的过期秒数
	ExpireSeconds int64 `json:"expire_seconds"`
}
