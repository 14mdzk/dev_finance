package helper

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(fmt.Errorf("error at HashPassword: %w", err))
		return "", err
	}

	return string(bytePassword), nil
}
