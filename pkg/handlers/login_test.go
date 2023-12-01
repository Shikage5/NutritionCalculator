package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	testCases := []struct {
		desc                string
		reqMethod           string
		reqURL              string
		reqBody             []byte
		expectedStatusCode  int
		expectedBody        string
		mockAuthShouldFail  bool
		mockAuthShouldError bool
	}{
		{
			desc:                "GET request should return 200 OK",
			reqMethod:           http.MethodGet,
			reqURL:              "/login",
			reqBody:             nil,
			expectedStatusCode:  http.StatusOK,
			expectedBody:        "Login form",
			mockAuthShouldFail:  false,
			mockAuthShouldError: false,
		},
		{
			desc:                "POST request with valid credentials should return 200 OK",
			reqMethod:           http.MethodPost,
			reqURL:              "/login",
			reqBody:             []byte(`{"username": "testuser", "password": "testpassword"}`),
			expectedStatusCode:  http.StatusOK,
			expectedBody:        "Login successful",
			mockAuthShouldFail:  false,
			mockAuthShouldError: false,
		},
		{
			desc:                "POST request with invalid credentials should return 401 Unauthorized",
			reqMethod:           http.MethodPost,
			reqURL:              "/login",
			reqBody:             []byte(`{"username": "testuser", "password": "wrongpassword"}`),
			expectedStatusCode:  http.StatusUnauthorized,
			expectedBody:        "invalid credentials",
			mockAuthShouldFail:  true,
			mockAuthShouldError: false,
		},
		{
			desc:                "POST request with missing username should return 400 Bad Request",
			reqMethod:           http.MethodPost,
			reqURL:              "/login",
			reqBody:             []byte(`{"password": "testpassword"}`),
			expectedStatusCode:  http.StatusBadRequest,
			expectedBody:        "Missing username or password",
			mockAuthShouldFail:  false,
			mockAuthShouldError: true,
		},
		{
			desc:                "POST request with missing password should return 400 Bad Request",
			reqMethod:           http.MethodPost,
			reqURL:              "/login",
			reqBody:             []byte(`{"username": "testuser"}`),
			expectedStatusCode:  http.StatusBadRequest,
			expectedBody:        "Missing username or password",
			mockAuthShouldFail:  false,
			mockAuthShouldError: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockAuthService := &mockAuthService{shouldFail: tC.mockAuthShouldFail, shouldError: tC.mockAuthShouldError}

			var userRequest UserRequest
			json.Unmarshal(tC.reqBody, &userRequest)

			// Create a context with the UserRequest
			ctx := context.WithValue(context.Background(), contextkeys.UserRequestKey, userRequest)

			req, err := http.NewRequest(tC.reqMethod, tC.reqURL, bytes.NewBuffer(tC.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()
			handler := LoginHandler(mockAuthService)
			handler.ServeHTTP(rr, req)

			if rr.Code != tC.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tC.expectedStatusCode, rr.Code)
			}
			if rr.Body.String() != tC.expectedBody {
				t.Errorf("Expected body %s, got %s", tC.expectedBody, rr.Body.String())
			}
		})
	}
}

type mockAuthService struct {
	shouldFail  bool
	shouldError bool
}

func (m mockAuthService) Auth(user models.User) (bool, error) {
	if m.shouldFail {
		return false, auth.ErrInvalidCredentials
	}
	if m.shouldError {
		return false, errors.New("error")
	}
	return true, nil
}
