package registration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		desc     string
		username string
		password string
		hasError bool
	}{
		{
			desc:     "Successful registration",
			username: "testuser",
			password: "testpass",
			hasError: false,
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

			// Setup
			mockHashingService := &MockHashingService{}
			registrationService := &DefaultRegistrationService{
				HashingService: mockHashingService,
				DataFilePath:   tempFile.Name(),
			}

			// Execute
			err = registrationService.RegisterUser(tC.username, tC.password)

			// Assert
			if tC.hasError {
				assert.Error(t, err, "Expected an error but got none")
				assert.Contains(t, err.Error(), "password hashing", "Expected error related to password hashing")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
			}

		})
	}
}

type MockHashingService struct {
	password string
}

func (m *MockHashingService) HashPassword(password string) (string, error) {
	// Mock the hashing logic for testing
	return "mockedhash", nil
}
