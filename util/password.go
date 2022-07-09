package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//生成哈希散列密码值
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}

	return string(hashedPassword), nil
}

//检查密码是否正确
func CheckPasssword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
