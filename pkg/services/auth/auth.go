package auth

import (
	"NutritionCalculator/data/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Auth(inputUser models.UserCredentials) (bool, error)
}

type DefaultAuthService struct {
	FilePath string
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (a DefaultAuthService) Auth(inputUser models.UserCredentials) (bool, error) {
	userList, err := models.ReadUsersFromJSONFile(a.FilePath)
	if err != nil {
		return false, err
	}
	for _, u := range userList {
		if inputUser.Username == u.Username {
			// User found, check password hash
			err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(inputUser.PasswordHash))
			if err != nil {
				return false, ErrInvalidCredentials // Wrong password
			}
			return true, nil // User exists with the correct password
		}
	}

	return false, ErrInvalidCredentials // User not found
}
