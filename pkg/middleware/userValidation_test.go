package middleware_test

import (
	"NutritionCalculator/pkg/middleware"
	"NutritionCalculator/pkg/services/validation"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	validator := validation.NewCredentialsValidator()
	testCases := []struct {
		desc     string
		username string
		password string
		status   int
	}{
		{
			desc:     "Valid Request",
			username: "john",
			password: "secretpassword",
			status:   http.StatusOK,
		},
		{
			desc:     "Username field empty",
			username: "",
			password: "secretpassword",
			status:   http.StatusBadRequest,
		},
		{
			desc:     "Password field empty",
			username: "john",
			password: "",
			status:   http.StatusBadRequest,
		},
		{
			desc:     "EmptyCredentials",
			username: "",
			password: "",
			status:   http.StatusBadRequest,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			form := url.Values{}
			form.Set("username", tC.username)
			form.Set("password", tC.password)

			req, err := http.NewRequest("POST", "/some-endpoint", nil)
			assert.NoError(t, err)
			req.Form = form

			// Create a response recorder to capture the response
			recorder := httptest.NewRecorder()

			// Create an http.Handler that includes the middleware
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			middleware := middleware.ValidateUser(validator, handler)
			// Serve the request through the middleware
			middleware.ServeHTTP(recorder, req)

			// Check if the status code matches the expected result
			assert.Equal(t, tC.status, recorder.Code)
		})
	}
}
