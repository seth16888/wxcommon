package helpers

func ValidatePortalReq(timestamp, nonce, signature, token string) bool {
	excepted := Signature(token, timestamp, nonce)
	return excepted == signature
}
