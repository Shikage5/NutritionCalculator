package middleware

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func TestValidateUser(t *testing.T) {
	testCases := []struct {
		desc        string
		requestBody string
		hasError    bool
	}{
		{
			desc:        "Valid Request",
			requestBody: `{"username": "john", "password": "secretpassword"}`,
			hasError:    false,
		},
		{
			desc:        "Username field empty",
			requestBody: `{"username": "", "password": "secretpassword"}`,
			hasError:    true,
		},
		{
			desc:        "Password field empty",
			requestBody: `{"username": "john", "password": ""}`,
			hasError:    true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/register", strings.NewReader(tC.requestBody))
			assert.NoError(t, err)

		})
	}
}
