package mycrypt

import (
	"crypto/sha256"
	"encoding/base64"
)

// encrypt string to base64 crypto using SHA256
func Mycrypt(text string) []byte {

	h := sha256.New()

	h.Write([]byte(text))
	return h.Sum(nil)
}

// decrypt from base64 to decrypted []byte
func Mydecrypt(cryptoText string) []byte {

	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	h := sha256.New()

	h.Write(ciphertext)

	return h.Sum(nil)
}
