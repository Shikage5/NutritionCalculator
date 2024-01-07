package auth

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Auth(inputUser models.UserRequest) (bool, error)
}

type DefaultAuthService struct {
	FilePath string
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func NewAuthService(filePath string) *DefaultAuthService {
	return &DefaultAuthService{FilePath: filePath}
}

func (a DefaultAuthService) Auth(inputUser models.UserRequest) (bool, error) {
	userList, err := utils.ReadUserCredFromJSONFile(a.FilePath)
	if err != nil {
		return false, err
	}
	for _, u := range userList {
		if inputUser.Username == u.Username {
			// User found, check password hash
			err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(inputUser.Password))
			if err != nil {
				return false, ErrInvalidCredentials // Wrong password
			}
			return true, nil // User exists with the correct password
		}
	}

	return false, ErrInvalidCredentials // User not found
}
