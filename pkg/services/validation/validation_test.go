package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCredentials(t *testing.T) {
	testCases := []struct {
		desc     string
		username string
		password string
		hasError bool
	}{
		{
			desc:     "Empty username and password",
			username: "",
			password: "",
			hasError: true,
		},
		{
			desc:     "Empty username",
			username: "",
			password: "password",
			hasError: true,
		},
		{
			desc:     "Empty password",
			username: "username",
			password: "",
			hasError: true,
		},
		{
			desc:     "Valid username and password",
			username: "username",
			password: "password",
			hasError: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Setup
			validationService := &CredentialsValidationService{}

			// Execute
			result := validationService.ValidateCredentials(tC.username, tC.password)

			// Assert
			if tC.hasError {
				assert.False(t, result, "Expected validation to fail but it passed")
			} else {
				assert.True(t, result, "Expected validation to pass but it failed")
			}
		})
	}
}
