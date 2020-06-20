package security

import (
	"crypto/md5"
	"encoding/hex"
)

// WithMd5
func WithMd5(plaintext string) string {
	h := md5.New()
	h.Write([]byte(plaintext))
	return hex.EncodeToString(h.Sum(nil))
}

// CheckPassword
func CheckPassword(a, b string) bool {
	if a == b {
		return true
	}
	return false
}
