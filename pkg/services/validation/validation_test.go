package validation

import (
	"NutritionCalculator/data/models"
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
			validationService := &DefaultValidationService{}

			userRequest := models.UserRequest{
				Username: tC.username,
				Password: tC.password,
			}

			// Execute
			err := validationService.ValidateUserRequest(userRequest)

			// Assert
			if tC.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
