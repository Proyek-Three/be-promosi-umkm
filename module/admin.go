package module

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("jdsfsdfsd54353454354235234") 

// GenerateToken membuat token JWT
func GenerateToken(username string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   username,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidateToken memvalidasi token JWT
func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, err
}
