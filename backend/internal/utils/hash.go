package utils

import "golang.org/x/crypto/bcrypt"

// HashedPassword hashea una contraseña
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compara una contraseña hasheada con una contraseña sin hashear
func CheckPasswordHash(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
