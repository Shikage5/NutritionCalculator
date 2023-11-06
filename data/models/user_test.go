package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
