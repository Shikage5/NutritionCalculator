package auth

import (
	"NutritionCalculator/data/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Auth(inputUser models.User, userList []models.User) (bool, error) {
	// Validate input
	if inputUser.Username == "" || inputUser.PasswordHash == "" {
		return false, errors.New("username and password cannot be empty")
	}

	for _, u := range userList {
		if inputUser.Username == u.Username {
			// User found, check password hash
			if inputUser.PasswordHash == u.PasswordHash {
				return true, nil // User exists with the correct password
			}
			return false, nil // User found, but incorrect password
		}
	}

	return false, nil // User not found
}

func HashPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 12
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
