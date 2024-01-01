package registration

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/hashing"
	"encoding/json"
	"os"
)

// RegistrationService defines the interface for user registration.
type RegistrationService interface {
	RegisterUser(username, password string) error
}

// DefaultRegistrationService is the default implementation of RegistrationService.
type DefaultRegistrationService struct {
	HashingService hashing.HashingService
	FilePath       string
	UserDataPath   string
}

// RegisterUser implements the registration logic.
func (s *DefaultRegistrationService) RegisterUser(username, password string) error {

	hashedPassword, err := s.HashingService.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.UserCredentials{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	if err := models.WriteUserInJSONFile(user, s.FilePath); err != nil {
		return err
	}
	// Create a new user data file
	userDataFile, err := os.Create(s.UserDataPath + username + ".json")
	if err != nil {
		return err
	}
	defer userDataFile.Close()

	// Initialize the file with an empty JSON object
	emptyData := make(map[string]interface{})
	if err := json.NewEncoder(userDataFile).Encode(emptyData); err != nil {
		return err
	}

	return nil
}
