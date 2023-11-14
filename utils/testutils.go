package utils

import (
	"os"
	"testing"
)

// Helper function to create a temporary test JSON file
func CreateTempTestJSONFile(t *testing.T, content string) string {
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
