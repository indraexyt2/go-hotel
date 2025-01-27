package helpers

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauth2Config = &oauth2.Config{
		ClientID:     "63186506952-4uh4cg88mh7lkd8vh5fi9sdgm4hgsauq.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/api/ums/v1/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	State = "state"
)
