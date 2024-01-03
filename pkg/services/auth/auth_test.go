package auth

import (
	"NutritionCalculator/data/models"
	"NutritionCalculator/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var hashedPassword1, _ = bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
var hashedPassword2, _ = bcrypt.GenerateFromPassword([]byte("password2"), bcrypt.DefaultCost)
var hashedPassword3, _ = bcrypt.GenerateFromPassword([]byte("password3"), bcrypt.DefaultCost)

var mockUserCredentials = []models.UserCredentials{
	{
		Username:     "user1",
		PasswordHash: string(hashedPassword1),
	},
	{
		Username:     "user2",
		PasswordHash: string(hashedPassword2),
	},
	{
		Username:     "user3",
		PasswordHash: string(hashedPassword3),
	},
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name     string
		input    models.UserRequest
		hasError bool
		errMsg   string
	}{
		{
			name:     "user exists, right password",
			input:    models.UserRequest{Username: "user1", Password: "password1"},
			hasError: false,
		},
		{
			name:     "user exists, wrong password",
			input:    models.UserRequest{Username: "user1", Password: "wrongPassword"},
			hasError: true,
			errMsg:   "invalid credentials",
		},
		{
			name:     "user does not exist",
			input:    models.UserRequest{Username: "nonexistentUser", Password: "password1"},
			hasError: true,
			errMsg:   "invalid credentials",
		},
		{
			name:     "empty username and password",
			input:    models.UserRequest{Username: "", Password: ""},
			hasError: true,
			errMsg:   "invalid credentials",
		},
		{
			name:     "empty username, valid password",
			input:    models.UserRequest{Username: "", Password: "password1"},
			hasError: true,
			errMsg:   "invalid credentials",
		},
		{
			name:     "valid username, empty password",
			input:    models.UserRequest{Username: "user1", Password: ""},
			hasError: true,
			errMsg:   "invalid credentials",
		},
		{
			name:     "nonexistent username, valid password",
			input:    models.UserRequest{Username: "nonexistentUser", Password: "password1"},
			hasError: true,
			errMsg:   "invalid credentials",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create temp file
			tempFile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write mock users to temp file
			for _, u := range mockUserCredentials {
				err := utils.WriteUserCredInJSONFile(u, tempFile.Name())
				if err != nil {
					t.Fatal(err)
				}
			}

			// Create auth service
			authService := DefaultAuthService{FilePath: tempFile.Name()}

			// Test auth service
			auth, err := authService.Auth(tc.input)
			if tc.hasError {
				assert.Error(t, err)
				assert.Equal(t, tc.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.True(t, auth)
			}
		})
	}

}
