package auth

import (
    "context"
    "encoding/json"
    // "errors"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
    ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
    ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
    RedirectURL:  "http://localhost:8080/v1/auth/google/callback",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}

func GoogleLoginURL() string {
    return googleOAuthConfig.AuthCodeURL("state-token")
}

func HandleGoogleCallback(code string) (string, error) {
    token, err := googleOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        return "", err
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var userInfo struct {
        Email string `json:"email"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        return "", err
    }

    return userInfo.Email, nil
}