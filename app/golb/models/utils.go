package models

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePassword(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))
	return pwdHash
}
