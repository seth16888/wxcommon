package helpers

// ValidatePortalReq 校验微信签名
func ValidatePortalReq(timestamp, nonce, signature, token string) bool {
	excepted := Signature(token, timestamp, nonce)
	return excepted == signature
}

// ValidateEncryptSignature 校验加密签名
func ValidateEncryptSignature(token, timestamp, nonce, msgSignature, encrypt string) bool {
  excepted := Signature(token, timestamp, nonce, encrypt)
  return excepted == msgSignature
}
