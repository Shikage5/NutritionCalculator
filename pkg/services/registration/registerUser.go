package registration

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/pkg/services/hashing"
	"NutritionCalculator/utils"
	"encoding/json"
	"os"
)

// RegistrationService defines the interface for user registration.
type RegistrationService interface {
	RegisterUser(models.UserRequest) error
}

// DefaultRegistrationService is the default implementation of RegistrationService.
type DefaultRegistrationService struct {
	HashingService      hashing.HashingService
	CredentialsFilepath string
	UserDataPath        string
}

// NewRegistrationService creates a new instance of DefaultRegistrationService.
func NewRegistrationService(hashService hashing.HashingService, filePath string, userDataPath string) *DefaultRegistrationService {
	return &DefaultRegistrationService{HashingService: hashService, CredentialsFilepath: filePath, UserDataPath: userDataPath}
}

// RegisterUser implements the registration logic.
func (s *DefaultRegistrationService) RegisterUser(userRequest models.UserRequest) error {

	// Check if the userCredentials.json file exists and create it if it doesn't
	if _, err := os.Stat(s.CredentialsFilepath); os.IsNotExist(err) {
		file, err := os.Create(s.CredentialsFilepath)
		if err != nil {
			return err
		}
		file.Close()
	}

	hashedPassword, err := s.HashingService.HashPassword(userRequest.Password)
	if err != nil {
		return err
	}

	user := models.UserCredentials{
		Username:     userRequest.Username,
		PasswordHash: hashedPassword,
	}

	if err := utils.WriteUserCredInJSONFile(user, s.CredentialsFilepath); err != nil {
		return err
	}
	// Create a new user data file
	userDataFile, err := os.Create(s.UserDataPath + userRequest.Username + ".json")
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
