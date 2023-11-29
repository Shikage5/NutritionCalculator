package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
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
			reqURL:             "/login",
			reqBody:            nil,
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Login form",
		},
		{
			desc:               "POST request with valid credentials should return 200 OK",
			reqMethod:          http.MethodPost,
			reqURL:             "/login",
			reqBody:            []byte(`{"username": "testuser", "password": "testpassword"}`),
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Login successful",
		},
		{
			desc:               "POST request with invalid credentials should return 401 Unauthorized",
			reqMethod:          http.MethodPost,
			reqURL:             "/login",
			reqBody:            []byte(`{"username": "testuser", "password": "wrongpassword"}`),
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       "Invalid credentials",
		},
		{
			desc:               "POST request with missing username should return 400 Bad Request",
			reqMethod:          http.MethodPost,
			reqURL:             "/login",
			reqBody:            []byte(`{"password": "testpassword"}`),
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Missing username",
		},
		{
			desc:               "POST request with missing password should return 400 Bad Request",
			reqMethod:          http.MethodPost,
			reqURL:             "/login",
			reqBody:            []byte(`{"username": "testuser"}`),
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Missing password",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockAuthService := &mockAuthService{shouldFail: tC.expectedStatusCode == http.StatusUnauthorized}

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
	shouldFail bool
}

func (m mockAuthService) Auth(user models.User) (bool, error) {
	if m.shouldFail {
		return false, nil
	}
	return true, nil
}
