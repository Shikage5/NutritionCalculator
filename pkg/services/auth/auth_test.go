package auth

import (
	"NutritionCalculator/data/models"
	"os"
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

func TestAuth(t *testing.T) {
	testCases := []struct {
		name     string
		input    models.User
		expected bool
		hasError bool
	}{
		{
			name:     "user exists, right password",
			input:    models.User{Username: "user1", PasswordHash: "passwordHash1"},
			expected: true,
			hasError: false,
		},
		{
			name:     "user exists, wrong password",
			input:    models.User{Username: "user1", PasswordHash: "wrongPassword"},
			expected: false,
			hasError: false,
		},
		{
			name:     "user does not exist",
			input:    models.User{Username: "nonexistentUser", PasswordHash: "passwordHash1"},
			expected: false,
			hasError: false,
		},
		{
			name:     "empty username and password",
			input:    models.User{Username: "", PasswordHash: ""},
			expected: false,
			hasError: true,
		},
		{
			name:     "empty username, valid password",
			input:    models.User{Username: "", PasswordHash: "passwordHash1"},
			expected: false,
			hasError: true,
		},
		{
			name:     "valid username, empty password",
			input:    models.User{Username: "user1", PasswordHash: ""},
			expected: false,
			hasError: true,
		},
		{
			name:     "nonexistent username, valid password",
			input:    models.User{Username: "nonexistentUser", PasswordHash: "passwordHash1"},
			expected: false,
			hasError: false,
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
			for _, u := range mockUsers {
				err := models.WriteUserInJSONFile(u, tempFile.Name())
				if err != nil {
					t.Fatal(err)
				}
			}

			// Create auth service
			authService := DefaultAuthService{FilePath: tempFile.Name()}

			// Test auth service
			auth, err := authService.Auth(tc.input)
			if tc.hasError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.Equal(t, tc.expected, auth)
			}
		})
	}

}
