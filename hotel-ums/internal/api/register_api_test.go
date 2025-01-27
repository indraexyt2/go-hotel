package api_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hotel-ums/helpers"
	"hotel-ums/internal/api"
	"hotel-ums/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockRegisterService struct {
	mock.Mock
}

func (m *MockRegisterService) RegisterNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), nil
}

func (m *MockRegisterService) EmailVerification(ctx context.Context, tokenVerify string) error {
	args := m.Called(ctx, tokenVerify)
	return args.Error(0)
}

func (m *MockRegisterService) ResendEmailVerification(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Error(args ...interface{}) {
}

func (m *MockLogger) Info(args ...interface{}) {
}

func (m *MockLogger) Warn(args ...interface{}) {
}

func (m *MockLogger) WithField(key string, value interface{}) helpers.LoggerInterface {
	return nil
}

func TestRegisterAPI_RegisterNewUser(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can register new user", func(t *testing.T) {
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		requestBody := `{
			"username": "indra",
			"password": "password123",
			"email": "indra@example.com",
			"full_name": "Indra Wansyah",
			"phone": "1234567890",
			"address": "Karawang, Indonesia"
		}`

		mockUser := &models.User{
			Username: "indra",
			Email:    "indra@example.com",
			FullName: "Indra Wansyah",
			Phone:    "1234567890",
			Address:  "Karawang, Indonesia",
		}

		MockRegisterSvc.On("RegisterNewUser", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
			return u.Username == mockUser.Username && u.Email == mockUser.Email
		})).Return(mockUser, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := api.RegisterNewUser(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponse := `{
			"message": "success",
			"data": {
				"id": 0,
				"username": "indra",
				"email": "indra@example.com",
				"full_name": "Indra Wansyah",
				"phone": "1234567890",
				"address": "Karawang, Indonesia",
				"is_verified": false,
				"photo_path": "",
				"role": "",
				"source": ""
			}
		}`
		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})

	t.Run("cannot register new user", func(t *testing.T) {
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		requestBody := `{
        	"username": "indra",
        	"password": "password123",
        	"email": "indra@example.com",
        	"full_name": "Indra Wansyah",
        	"phone": "1234567890",
        	"address": "Karawang, Indonesia"
    	}`

		mockUser := &models.User{
			Username: "indra",
			Email:    "indra@example.com",
			FullName: "Indra Wansyah",
			Phone:    "1234567890",
			Address:  "Karawang, Indonesia",
		}

		MockRegisterSvc.On("RegisterNewUser", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
			return u.Username == mockUser.Username && u.Email == mockUser.Email
		})).Return(nil, errors.New("username already exists"))

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := api.RegisterNewUser(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expectedResponse := `{
        	"message": "username already exists"
    	}`
		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})
}

func TestRegisterAPI_EmailVerification(t *testing.T) {
	t.Run("can verify email", func(t *testing.T) {
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		MockRegisterSvc.On("EmailVerification", mock.Anything, "valid-token").Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register/verify/:token", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("token")
		c.SetParamValues("valid-token")

		err := api.EmailVerification(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponse := `{
			"message": "success"
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})

	t.Run("cannot verify email", func(t *testing.T) {
		helpers.Logger = &MockLogger{}
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		MockRegisterSvc.On("EmailVerification", mock.Anything, "invalid-token").Return(errors.New("invalid token"))

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register/verify/:token", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("token")
		c.SetParamValues("invalid-token")

		err := api.EmailVerification(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expectedResponse := `{
			"message": "invalid token"
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})
}

func TestRegisterAPI_ResendEmailVerification(t *testing.T) {
	helpers.Logger = &MockLogger{}

	t.Run("can resend email verification", func(t *testing.T) {
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		MockRegisterSvc.On("ResendEmailVerification", mock.Anything, "indra@example.com").Return(nil)

		requestBody := []byte(`{"email": "indra@example.com"}`)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register/email-verification", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err := api.ResendEmailVerification(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponse := `{
			"message": "success"
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})

	t.Run("cannot resend email verification", func(t *testing.T) {
		MockRegisterSvc := new(MockRegisterService)
		api := api.NewRegisterAPI(MockRegisterSvc)
		e := echo.New()

		MockRegisterSvc.On("ResendEmailVerification", mock.Anything, "invalid-email").Return(errors.New("failed to bind request"))

		requestBody := []byte(`{"email": "invalid-email"}`)

		req := httptest.NewRequest(http.MethodPost, "/api/ums/v1/register/email-verification", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err := api.ResendEmailVerification(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expectedResponse := `{
			"message": "failed to bind request"
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
		MockRegisterSvc.AssertExpectations(t)
	})
}
