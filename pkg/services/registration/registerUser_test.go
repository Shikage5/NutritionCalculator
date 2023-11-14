package registration

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUsers = []models.User{
	{
		Username:     "user1",
		PasswordHash: "passwordHash1",
	},
	{
		Username:     "user2",
		PasswordHash: "passwordHash2",
	},
	{
		Username:     "user3",
		PasswordHash: "passwordHash3",
	},
}

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
		{
			desc:     "Failed password hashing",
			username: "testuser",
			password: "", // invalid password to trigger hashing error
			hasError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			usersJSON, err := json.Marshal(mockUsers)
			if err != nil {
				t.Fatal(err)
			}

			tempFile := utils.CreateTempTestJSONFile(t, string(usersJSON))
			// Create an instance of DefaultRegistrationService

			service := &DefaultRegistrationService{}

			// Call the RegisterUser function with the test case input
			registerErr := service.RegisterUser(tC.username, tC.password, tempFile)

			if tC.hasError {
				assert.Error(t, registerErr, "Expected an error")
			} else {
				assert.NoError(t, registerErr, "Unexpected error")
			}

			// If no error is expected, you may perform additional assertions
			if !tC.hasError {
				// Add additional assertions as needed
			}
		})
	}
}
