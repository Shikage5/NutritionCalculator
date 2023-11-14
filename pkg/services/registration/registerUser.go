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
type DefaultRegistrationService struct{}

// RegisterUser implements the registration logic.
func (s *DefaultRegistrationService) RegisterUser(username, password string, filename string) error {
	hashedPassword, err := hashing.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	// Write the user to the "data/users.json" file
	if err := models.WriteUserInJSONFile(user, filename); err != nil {
		return err
	}

	return nil
}
