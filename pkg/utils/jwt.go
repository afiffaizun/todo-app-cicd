package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

var jwtSecret []byte

// InitJWT initializes JWT secret
func InitJWT(secret string) {
    jwtSecret = []byte(secret)
}

// GenerateToken generates a new JWT token
func GenerateToken(userID uint, email string) (string, error) {
    claims := &JWTClaim{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// GenerateRefreshToken generates a refresh token with longer expiry
func GenerateRefreshToken(userID uint, email string) (string, error) {
    claims := &JWTClaim{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateToken validates JWT token and returns claims
func ValidateToken(tokenString string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &JWTClaim{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        },
    )

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaim)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}