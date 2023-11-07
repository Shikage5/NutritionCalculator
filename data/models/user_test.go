package models

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUsers = []User{
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
		expected []User
		hasError bool
	}{
		{
			name:     "Valid JSON",
			input:    `[{"username": "user1", "passwordHash": "hash1"}]`,
			expected: []User{{Username: "user1", PasswordHash: "hash1"}},
			hasError: false,
		},
		{
			name:     "Empty File",
			input:    "",
			expected: nil,
			hasError: true,
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
			tempFile := createTempTestJSONFile(t, tc.input)
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
		input    User
		hasError bool
	}{
		{
			desc:     "User added correctly",
			input:    User{Username: "newUser", PasswordHash: "newPasswordHash"},
			hasError: false,
		},
		{
			desc:     "Username already exists",
			input:    User{Username: "user1", PasswordHash: "newPasswordHash"},
			hasError: true,
		},
		{
			desc:     "Empty Username",
			input:    User{Username: "", PasswordHash: "newPasswordHash"},
			hasError: true,
		},
		{
			desc:     "Empty Password",
			input:    User{Username: "newUser", PasswordHash: ""},
			hasError: true,
		},
		{
			desc:     "Empty Input",
			input:    User{Username: "", PasswordHash: ""},
			hasError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//mockUsers to JSON
			usersJSON, err := json.Marshal(mockUsers)
			if err != nil {
				t.Fatal(err)
			}
			tempFile := createTempTestJSONFile(t, string(usersJSON))
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

// Helper function to create a temporary test JSON file
func createTempTestJSONFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "test_users.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()
	return tempFile.Name()
}
