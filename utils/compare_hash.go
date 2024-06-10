package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHash(hash string, pass string) bool {
	// Texto sin hash
	plaintext := []byte(pass)

	// Hash existente
	existingHash := []byte(hash)

	// Verificar el hash
	err := bcrypt.CompareHashAndPassword(existingHash, plaintext)
	if err != nil {
		return false
	} else {
		return true
	}
}
