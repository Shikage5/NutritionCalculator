package models

import (
	"NutritionCalculator/utils"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUserCredentials = []UserCredentials{
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

func TestReadUsersFromJSONFile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []UserCredentials
		hasError bool
	}{
		{
			name:     "Valid JSON",
			input:    `[{"username": "user1", "passwordHash": "hash1"}]`,
			expected: []UserCredentials{{Username: "user1", PasswordHash: "hash1"}},
			hasError: false,
		},
		{
			name:     "Empty File",
			input:    "",
			expected: []UserCredentials{},
			hasError: false,
		},
		{
			name:     "Invalid JSON",
			input:    "invalid json",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempFile := utils.CreateTempTestJSONFile(t, tc.input)
			defer os.Remove(tempFile)

			users, err := ReadUsersFromJSONFile(tempFile)

			if tc.hasError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}

			assert.Equal(t, tc.expected, users, "Unexpected users")
		})
	}
}

func TestWriteUserInJSONFile(t *testing.T) {
	testCases := []struct {
		desc     string
		input    UserCredentials
		hasError bool
	}{
		{
			desc:     "User added correctly",
			input:    UserCredentials{Username: "newUser", PasswordHash: "newPasswordHash"},
			hasError: false,
		},
		{
			desc:     "Username already exists",
			input:    UserCredentials{Username: "user1", PasswordHash: "newPasswordHash"},
			hasError: true,
		},
		{
			desc:     "Empty Username",
			input:    UserCredentials{Username: "", PasswordHash: "newPasswordHash"},
			hasError: true,
		},
		{
			desc:     "Empty Password",
			input:    UserCredentials{Username: "newUser", PasswordHash: ""},
			hasError: true,
		},
		{
			desc:     "Empty Input",
			input:    UserCredentials{Username: "", PasswordHash: ""},
			hasError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//mockUserCredentials to JSON
			usersJSON, err := json.Marshal(mockUserCredentials)
			if err != nil {
				t.Fatal(err)
			}
			tempFile := utils.CreateTempTestJSONFile(t, string(usersJSON))
			writeErr := WriteUserInJSONFile(tC.input, tempFile)

			if tC.hasError {
				// Expecting an error, so check if err is not nil
				assert.Error(t, writeErr, "Expected an error")
			} else {
				// Not expecting an error, so check if err is nil
				assert.NoError(t, writeErr, "Expected no error")
			}
		})
	}
}
