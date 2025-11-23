package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define your secret key (in production, use environment variables)
var secretKey = []byte("your-secret-key-here")

const (
	accessTokenTTL   = 24 * time.Hour
	exchangeTokenTTL = 5 * time.Minute
)

// Custom claims structure
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ExchangeClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// EncodeJWT creates a new JWT token. Deprecated: use EncodeAccessToken.
func EncodeJWT(userID, username string) (string, error) {
	return EncodeAccessToken(userID, username)
}

// DecodeJWT validates and decodes a JWT token. Deprecated: use DecodeAccessToken.
func DecodeJWT(tokenString string) (*Claims, error) {
	return DecodeAccessToken(tokenString)
}

// EncodeAccessToken creates a signed access token with user claims.
func EncodeAccessToken(userID, username string) (string, error) {
	claims := Claims{
		UserID:           userID,
		Username:         username,
		RegisteredClaims: defaultRegisteredClaims(accessTokenTTL),
	}

	return signToken(claims)
}

// DecodeAccessToken validates and decodes an access token into Claims.
func DecodeAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, signingKeyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid access token")
}

// EncodeExchangeToken creates a short-lived token used for login exchanges.
func EncodeExchangeToken(email string) (string, error) {
	claims := ExchangeClaims{
		Email:            email,
		RegisteredClaims: defaultRegisteredClaims(exchangeTokenTTL),
	}

	return signToken(claims)
}

// DecodeExchangeToken validates and decodes an exchange token into ExchangeClaims.
func DecodeExchangeToken(tokenString string) (*ExchangeClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ExchangeClaims{}, signingKeyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ExchangeClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid exchange token")
}

// signToken signs any jwt.Claims with the configured key.
func signToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// signingKeyFunc centralizes validation of the signing algorithm/key.
func signingKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return secretKey, nil
}

// defaultRegisteredClaims builds RegisteredClaims with shared metadata.
func defaultRegisteredClaims(ttl time.Duration) jwt.RegisteredClaims {
	now := time.Now()
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    "click-in",
	}
}
