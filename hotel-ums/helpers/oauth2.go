package helpers

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"hotel-ums/internal/models"
	"os"
)

var (
	GoogleOauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/api/ums/v1/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	State = "state"
)

func RefreshAccessToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: refreshToken}
	source := GoogleOauth2Config.TokenSource(ctx, token)

	newToken, err := source.Token()
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func ValidateTokenGoogle(ctx context.Context, token *oauth2.Token) (*models.GoogleUserInfo, error) {
	resp := &models.GoogleUserInfo{}

	client := GoogleOauth2Config.Client(ctx, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
