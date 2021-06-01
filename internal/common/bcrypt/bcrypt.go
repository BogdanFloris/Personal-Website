package bcrypt

import "golang.org/x/crypto/bcrypt"

const HashCost = 14

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(bytes), err
}

func IsPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
