package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {

	testCases := []struct {
		desc               string
		reqMethod          string
		reqURL             string
		reqBody            []byte
		expectedStatusCode int
		expectedBody       string
	}{
		{
			desc:               "GET request should return 200 OK",
			reqMethod:          http.MethodGet,
			reqURL:             "/register",
			reqBody:            nil,
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Registration form",
		},
		{
			desc:               "POST request should return 201 Created",
			reqMethod:          http.MethodPost,
			reqURL:             "/register",
			reqBody:            []byte(`{"username": "testuser", "password": "testpassword"}`),
			expectedStatusCode: http.StatusCreated,
			expectedBody:       "Registration successful form",
		},
		{
			desc:               "POST request registration fail internal server error",
			reqMethod:          http.MethodPost,
			reqURL:             "/register",
			reqBody:            []byte(`{"username": "testuser"}`),
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "Registration fail form\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockService := &mockRegistrationService{shouldFail: tC.expectedStatusCode == http.StatusInternalServerError}
			req, err := http.NewRequest(tC.reqMethod, tC.reqURL, bytes.NewBuffer(tC.reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := RegisterHandler(mockService) // Pass the mock registration service to the handler
			handler.ServeHTTP(rr, req)

			if rr.Code != tC.expectedStatusCode {
				t.Errorf("Expected status code %d, but got %d", tC.expectedStatusCode, rr.Code)
			}

			if rr.Body.String() != tC.expectedBody {
				t.Errorf("Expected body %q, but got %q", tC.expectedBody, rr.Body.String())
			}
		})
	}
}

type mockRegistrationService struct {
	shouldFail bool
}

func (m *mockRegistrationService) RegisterUser(username, password string, filename string) error {
	if m.shouldFail {
		// Simulate failure
		return errors.New("Registration fail form")
	}
	// Simulate success
	return nil
}
