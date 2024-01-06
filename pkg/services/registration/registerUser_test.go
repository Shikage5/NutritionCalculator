package registration

import (
	"NutritionCalculator/data/models"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		desc                         string
		username                     string
		password                     string
		mockHashingServiceShouldFail bool
		hasError                     bool
	}{
		{
			desc:                         "Successful registration",
			username:                     "testuser",
			password:                     "testpass",
			mockHashingServiceShouldFail: false,
			hasError:                     false,
		},
		{
			desc:                         "Hashing service fails",
			username:                     "testuser",
			password:                     "testpass",
			mockHashingServiceShouldFail: true,
			hasError:                     true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Create a temporary file
			tempFile, err := os.CreateTemp("", "user_registration_test")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			// Create a temporary directory for user data files
			tempDir, err := os.MkdirTemp("", "user_data_test")
			if err != nil {
				t.Fatalf("Failed to create temporary directory: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Setup
			mockHashingService := &MockHashingService{shouldFail: tC.mockHashingServiceShouldFail}

			registrationService := &DefaultRegistrationService{
				HashingService: mockHashingService,
				FilePath:       tempFile.Name(),
				UserDataPath:   tempDir + "/",
			}
			userRequest := models.UserRequest{
				Username: tC.username,
				Password: tC.password,
			}

			// Execute
			err = registrationService.RegisterUser(userRequest)

			// Assert
			if tC.hasError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")

				// Assert that the user data file was created
				_, err = os.Stat(tempDir + "/" + tC.username + ".json")
				assert.NoError(t, err, "Expected user data file to exist but it does not")
			}
		})
	}
}

type MockHashingService struct {
	shouldFail bool
}

func (m *MockHashingService) HashPassword(password string) (string, error) {
	if m.shouldFail {
		return "", errors.New("mocked hash error")
	}
	return "mockedhash", nil
}
