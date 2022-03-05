package digest

import "encoding/base64"

// EncodeBase64 converts an input string to base64 string.
func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeBase64 decode a base64 string.
func DecodeBase64(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
