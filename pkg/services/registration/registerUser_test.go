package registration

import (
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
			// Setup
			mockHashingService := &MockHashingService{shouldFail: tC.mockHashingServiceShouldFail}

			registrationService := &DefaultRegistrationService{
				HashingService: mockHashingService,
				DataFilePath:   tempFile.Name(),
			}

			// Execute
			err = registrationService.RegisterUser(tC.username, tC.password)

			// Assert
			if tC.hasError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
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
