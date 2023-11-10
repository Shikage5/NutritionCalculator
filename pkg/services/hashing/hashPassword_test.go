package hashing

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		desc          string
		inputPassword string
	}{
		{
			desc:          "Valid Password",
			inputPassword: "securePassword123",
		},
		{
			desc:          "Empty Password",
			inputPassword: "",
		},
		{
			desc:          "Password with Whitespace",
			inputPassword: "  passwordWithSpaces  ",
		},
		// Add more test cases as needed
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Run the actual test logic
			hashedPassword, err := HashPassword(tC.inputPassword)

			// Check for errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check if the hashed password is not empty
			if hashedPassword == "" {
				t.Error("Hashed password is empty")
			}

			// Check if the hashed password is a valid bcrypt hash
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tC.inputPassword))
			if err != nil {
				t.Error("Generated hash is not a valid bcrypt hash")
			}
		})
	}
}
