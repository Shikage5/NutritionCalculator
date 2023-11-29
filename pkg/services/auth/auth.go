package auth

import (
	"NutritionCalculator/data/models"
	"errors"
)

type AuthService interface {
	Auth(inputUser models.User) (bool, error)
}

type DefaultAuthService struct {
	FilePath string
}

func (a DefaultAuthService) Auth(inputUser models.User) (bool, error) {
	userList, err := models.ReadUsersFromJSONFile(a.FilePath)
	if err != nil {
		return false, err
	}
	for _, u := range userList {
		if inputUser.Username == u.Username {
			// User found, check password hash
			if inputUser.PasswordHash == u.PasswordHash {
				return true, nil // User exists with the correct password
			}
			return false, errors.New("invalid credentials") // User exists with the wrong password
		}
	}

	return false, errors.New("invalid credentials") // User not found
}
