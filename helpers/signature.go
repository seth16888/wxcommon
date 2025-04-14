package helpers

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
)

func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// WXUrlSign 微信公众号 url 签名.
func WXUrlSign(token, timestamp, nonce string) (signature string) {
	strArr := sort.StringSlice{token, timestamp, nonce}
	strArr.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strArr[0]...)
	buf = append(buf, strArr[1]...)
	buf = append(buf, strArr[2]...)

	hash := sha1.Sum(buf)
	return hex.EncodeToString(hash[:])
}

// WXMsgSign 微信公众号/企业号 消息体签名.
func WXMsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strArr := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strArr.Sort()

	h := sha1.New()

	buf := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	buf.WriteString(strArr[0])
	buf.WriteString(strArr[1])
	buf.WriteString(strArr[2])
	buf.WriteString(strArr[3])
	buf.Flush()

	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}
