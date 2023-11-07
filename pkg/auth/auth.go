package auth

import (
	"NutritionCalculator/data/models"
	"errors"
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
