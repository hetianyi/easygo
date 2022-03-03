package base

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return crypt, nil
}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

/*func TestDES(t *testing.T) {
	originalText := convert.Int64ToStr(time.Now().UnixNano()) + "|123456789012345678901234567890df"
	fmt.Println(originalText)
	mytext := []byte(originalText)

	key := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}
	iv := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}

	cryptoText, _ := DesEncryption(key, iv, mytext)
	base64String := base64.StdEncoding.EncodeToString(cryptoText)
	//																// fDE1NjExMDAwOTY5NTM4MzI2MDD10pSRpEgO4DZti3M2w/YkYNKl0TvWxyQ=
	//base64String := base64.StdEncoding.EncodeToString(cryptoText) // 9dKUkaRIDuA2bYtzNsP2JGDSpdE71sck
	fmt.Println(base64String)
	bs, _ := base64.StdEncoding.DecodeString(base64String)
	decryptedText, _ := DesDecryption(key, iv, bs)
	fmt.Println(string(decryptedText))
}*/
