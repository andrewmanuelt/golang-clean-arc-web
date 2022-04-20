package helper

import "golang.org/x/crypto/bcrypt"

func Encrypt(password []byte) string {
	enc, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	ShowError(err)

	return string(enc)
}

func ValidatePassword(hash string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), password)

	return err == nil
}
