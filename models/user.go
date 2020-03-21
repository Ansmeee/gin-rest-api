package models

import (
	"crypto/md5"
	"io"
)

type User struct {
	userName string
	userPass string
	email string
}

// encrypt the userpass
func EncryptUserPass(userPass string) []byte {
	encrypter := md5.New()
	io.WriteString(encrypter, userPass)

	return encrypter.Sum(nil)
}