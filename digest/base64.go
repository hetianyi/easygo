package digest

import (
	"encoding/base64"
)

// EncodeBase64 将输入字符串转为base64编码
func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeBase64 将输入base64字符串转解码
func DecodeBase64(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
