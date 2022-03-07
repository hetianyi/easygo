package digest

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 计算字符串的MD5值
func Md5(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	sliceCipherStr := h.Sum(nil)
	sMd5 := hex.EncodeToString(sliceCipherStr)
	return sMd5
}
