package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_jwt_secret")

type Claims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a given user ID
func GenerateJWT(userID string) (string, error) {
    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(tokenStr string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}