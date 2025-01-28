package api_test

import (
	bytes3 "bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hotel-ums/helpers"
	"hotel-ums/internal/api"
	"hotel-ums/internal/models"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockLoginService struct {
	mock.Mock
}

func (m *MockLoginService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.LoginResponse), nil
}

func (m *MockLoginService) RefreshToken(ctx context.Context, refreshToken string, claimsToken *helpers.Claims) (*models.RefreshTokenResponse, error) {
	args := m.Called(ctx, refreshToken, claimsToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.RefreshTokenResponse), nil
}

func TestLoginAPI_Login(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can login", func(t *testing.T) {
		MockLoginSvc := new(MockLoginService)
		loginAPI := api.NewLoginAPI(MockLoginSvc)
		e := echo.New()

		bytes := make([]byte, 32)
		rand.Read(bytes)
		token := base64.URLEncoding.EncodeToString(bytes)

		bytes2 := make([]byte, 32)
		rand.Read(bytes2)
		refreshToken := base64.URLEncoding.EncodeToString(bytes)

		response := &models.LoginResponse{
			UserID:       1,
			FullName:     "Indra Wansyah",
			Token:        token,
			RefreshToken: refreshToken,
		}

		request := &models.LoginRequest{
			Username: "indrawansyah",
			Password: "123456",
		}

		MockLoginSvc.On("Login", mock.Anything, request).Return(response, nil)
		requestBody, _ := json.Marshal(request)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/login", bytes3.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err := loginAPI.Login(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponse := `{
			"message": "success",
			"data": {
				"user_id": 1,
				"full_name": "Indra Wansyah",
				"token": "` + token + `",
				"refresh_token": "` + refreshToken + `"
			}
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockLoginSvc.AssertExpectations(t)
	})

	t.Run("cannot login", func(t *testing.T) {
		MockLoginSvc := new(MockLoginService)
		loginAPI := api.NewLoginAPI(MockLoginSvc)
		e := echo.New()

		request := &models.LoginRequest{
			Username: "indrawansyah",
			Password: "123456",
		}

		requestBody, _ := json.Marshal(request)

		MockLoginSvc.On("Login", mock.Anything, request).Return(nil, errors.New("failed to login"))
		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/login", bytes3.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err := loginAPI.Login(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		MockLoginSvc.AssertExpectations(t)
	})
}

func TestLoginAPI_RefreshToken(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can refresh token", func(t *testing.T) {
		MockLoginSvc := new(MockLoginService)
		loginAPI := api.NewLoginAPI(MockLoginSvc)
		e := echo.New()

		bytes := make([]byte, 32)
		rand.Read(bytes)
		token := base64.URLEncoding.EncodeToString(bytes)

		bytes2 := make([]byte, 32)
		rand.Read(bytes2)
		refreshToken := base64.URLEncoding.EncodeToString(bytes)

		response := &models.RefreshTokenResponse{
			Token: token,
		}

		claimsToken := &helpers.Claims{
			UserID:   1,
			FullName: "Indra Wansyah",
			Email:    "Bk6tH@example.com",
			Role:     "admin",
		}

		MockLoginSvc.On("RefreshToken", mock.Anything, refreshToken, claimsToken).Return(response, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/refresh-token", nil)
		req.Header.Set("Authorization", refreshToken)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("token", claimsToken)
		err := loginAPI.RefreshToken(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponse := `{
			"message": "success",
			"data": {
				"token": "` + token + `"
			}
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockLoginSvc.AssertExpectations(t)
	})

	t.Run("cannot refresh token", func(t *testing.T) {
		MockLoginSvc := new(MockLoginService)
		loginAPI := api.NewLoginAPI(MockLoginSvc)
		e := echo.New()

		bytes2 := make([]byte, 32)
		rand.Read(bytes2)
		refreshToken := base64.URLEncoding.EncodeToString(bytes2)

		claimsToken := &helpers.Claims{
			UserID:   1,
			FullName: "Indra Wansyah",
			Email:    "Bk6tH@example.com",
			Role:     "admin",
		}

		MockLoginSvc.On("RefreshToken", mock.Anything, refreshToken, claimsToken).Return(nil, errors.New("unauthorized"))

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/refresh-token", nil)
		req.Header.Set("Authorization", refreshToken)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("token", claimsToken)
		err := loginAPI.RefreshToken(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

		expectedResponse := `{
			"message": "unauthorized"
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockLoginSvc.AssertExpectations(t)
	})
}
