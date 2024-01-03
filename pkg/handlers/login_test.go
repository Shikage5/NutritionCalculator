package handlers

import (
	"NutritionCalculator/data/models"
	contextkeys "NutritionCalculator/pkg/contextKeys"
	"NutritionCalculator/pkg/services/auth"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	testCases := []struct {
		desc                         string
		reqMethod                    string
		reqURL                       string
		reqBody                      []byte
		expectedStatusCode           int
		expectedBody                 string
		mockAuthShouldFail           bool
		mockSessionServiceShouldFail bool
	}{
		{
			desc:                         "GET request should return 200 OK",
			reqMethod:                    http.MethodGet,
			reqURL:                       "/login",
			reqBody:                      nil,
			expectedStatusCode:           http.StatusOK,
			expectedBody:                 "Login form",
			mockAuthShouldFail:           false,
			mockSessionServiceShouldFail: false,
		},
		{
			desc:                         "POST request with valid credentials should return 200 OK",
			reqMethod:                    http.MethodPost,
			reqURL:                       "/login",
			reqBody:                      []byte(`{"username": "testuser", "password": "testpassword"}`),
			expectedStatusCode:           http.StatusOK,
			expectedBody:                 "Login successful",
			mockAuthShouldFail:           false,
			mockSessionServiceShouldFail: false,
		},
		{
			desc:                         "POST request with invalid credentials should return 401 Unauthorized",
			reqMethod:                    http.MethodPost,
			reqURL:                       "/login",
			reqBody:                      []byte(`{"username": "testuser", "password": "wrongpassword"}`),
			expectedStatusCode:           http.StatusUnauthorized,
			expectedBody:                 "invalid credentials\n",
			mockAuthShouldFail:           true,
			mockSessionServiceShouldFail: false,
		},
		{
			desc:                         "POST request with valid credentials but session creation fails should return 500 Internal Server Error",
			reqMethod:                    http.MethodPost,
			reqURL:                       "/login",
			reqBody:                      []byte(`{"username": "testuser", "password": "testpassword"}`),
			expectedStatusCode:           http.StatusInternalServerError,
			expectedBody:                 "error creating session\n",
			mockAuthShouldFail:           false,
			mockSessionServiceShouldFail: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockAuthService := &mockAuthService{shouldFail: tC.mockAuthShouldFail}
			mockSessionService := &mockSessionService{shouldFail: tC.mockSessionServiceShouldFail}

			var userRequest models.UserRequest
			json.Unmarshal(tC.reqBody, &userRequest)

			// Create a context with the UserRequest
			ctx := context.WithValue(context.Background(), contextkeys.UserRequestKey, userRequest)

			req, err := http.NewRequest(tC.reqMethod, tC.reqURL, bytes.NewBuffer(tC.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()
			handler := LoginHandler(mockAuthService, mockSessionService)
			handler.ServeHTTP(rr, req)

			if rr.Code != tC.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tC.expectedStatusCode, rr.Code)
			}
			if rr.Body.String() != tC.expectedBody {
				fmt.Printf("Actual body: '%s'\n", rr.Body.String())
				fmt.Printf("Expected body: '%s'\n", tC.expectedBody)
				t.Errorf("Expected body %s, got %s", tC.expectedBody, rr.Body.String())
			}
		})
	}
}

type mockAuthService struct {
	shouldFail bool
}

func (m mockAuthService) Auth(inputUser models.UserRequest) (bool, error) {
	if m.shouldFail {
		return false, auth.ErrInvalidCredentials
	}
	return true, nil
}

type mockSessionService struct {
	shouldFail bool
}

func (m mockSessionService) CreateSession(username string, w http.ResponseWriter) error {
	if m.shouldFail {
		return fmt.Errorf("error creating session")
	}
	return nil
}

func (m *mockSessionService) GenerateSessionID() (string, error) {
	return "", nil
}
func (m *mockSessionService) GetSession(sessionID string) (string, bool) {
	return "", false
}
