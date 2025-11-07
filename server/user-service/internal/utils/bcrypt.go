package utils

import "golang.org/x/crypto/bcrypt"

const DefaultCostLevel = 10

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCostLevel)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
