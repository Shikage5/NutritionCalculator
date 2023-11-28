package middleware_test

import (
	"net/http"
	"net/url"
	"testing"
)

func TestValidateUser(t *testing.T) {
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

		})
	}
}

type MockCredentialsValidator struct {
	ShouldValidate bool
}

func (v *MockCredentialsValidator) ValidateCredentials(username, password string) bool {
	// Return the value of ShouldValidate, simulating a successful or unsuccessful validation
	return v.ShouldValidate
}
