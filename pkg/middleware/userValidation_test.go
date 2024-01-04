package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/middleware"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	testCases := []struct {
		desc        string
		userRequest models.UserRequest
		shouldFail  bool
		status      int
	}{
		{
			desc: "Valid Request",
			userRequest: models.UserRequest{
				Username: "john",
				Password: "secretpassword",
			},
			shouldFail: false,
			status:     http.StatusOK,
		},
		{
			desc: "Invalid Request",
			userRequest: models.UserRequest{
				Username: "",
				Password: "secretpassword",
			},
			shouldFail: true,
			status:     http.StatusBadRequest,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Setup
			jsonUserRequest, _ := json.Marshal(tC.userRequest)
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUserRequest))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			validator := &MockValidatorService{shouldFail: tC.shouldFail}
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				value := r.Context().Value(contextkeys.UserRequestKey)
				assert.Equal(t, tC.userRequest, value)
			})

			handler := middleware.ValidateUser(validator, nextHandler)

			// Execute
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// Assert
			assert.Equal(t, tC.status, rr.Code, "Unexpected status code")
		})
	}
}

// mock validator
type MockValidatorService struct {
	shouldFail bool
}

func (m *MockValidatorService) ValidateCredentials(username, password string) bool {
	if m.shouldFail {
		return false
	} else {
		return true
	}
}

func (m *MockValidatorService) ValidateFoodData(foodData models.FoodData) error {
	if m.shouldFail {
		return assert.AnError
	} else {
		return nil
	}

}
