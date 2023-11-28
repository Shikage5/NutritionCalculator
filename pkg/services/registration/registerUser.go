package registration

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/hashing"
)

// RegistrationService defines the interface for user registration.
type RegistrationService interface {
	RegisterUser(username, password string) error
}

// DefaultRegistrationService is the default implementation of RegistrationService.
type DefaultRegistrationService struct {
	HashingService hashing.HashingService
	DataFilePath   string
}

// RegisterUser implements the registration logic.
func (s *DefaultRegistrationService) RegisterUser(username, password string) error {

	hashedPassword, err := s.HashingService.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	if err := models.WriteUserInJSONFile(user, s.DataFilePath); err != nil {
		return err
	}

	return nil
}
