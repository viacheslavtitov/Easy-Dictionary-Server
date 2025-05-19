package middleware

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(password string) bool {
	// Check if the password is at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Check if the password contains at least one uppercase letter
	hasUpper := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return false
	}

	// Check if the password contains at least one lowercase letter
	hasLower := false
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
			break
		}
	}
	if !hasLower {
		return false
	}

	return true
}

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
