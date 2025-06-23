package utils

import "golang.org/x/crypto/bcrypt"


const (
	SALT_VALUE = 10
)

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), SALT_VALUE)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Valid(hashedPassword, plainPassword string) bool {
	 err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	 return err == nil
}