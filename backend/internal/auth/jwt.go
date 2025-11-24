package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/models"
)

const (
	accessTokenTTL   = 24 * time.Hour
	exchangeTokenTTL = 5 * time.Minute
)

// JWTManager encapsulates token operations bound to a config.
type JWTManager struct {
	cfg *config.Config
}

func NewJWT(cfg *config.Config) *JWTManager {
	return &JWTManager{cfg: cfg}
}

// Custom claims structure
type Claims struct {
	models.User
	jwt.RegisteredClaims
}

type ExchangeClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// EncodeAccessToken creates a signed access token with user claims.

func (m *JWTManager) EncodeAccessToken(userModel models.User) (string, error) {
	claims := Claims{
		User:             userModel,
		RegisteredClaims: defaultRegisteredClaims(accessTokenTTL),
	}

	return m.signToken(claims)
}

// DecodeAccessToken validates and decodes an access token into Claims.
func (m *JWTManager) DecodeAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, m.signingKeyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid access token")
}

// EncodeExchangeToken creates a short-lived token used for login exchanges.
func (m *JWTManager) EncodeExchangeToken(userID string) (string, error) {
	claims := ExchangeClaims{
		UserID:           userID,
		RegisteredClaims: defaultRegisteredClaims(exchangeTokenTTL),
	}

	return m.signToken(claims)
}

// DecodeExchangeToken validates and decodes an exchange token into ExchangeClaims.
func (m *JWTManager) DecodeExchangeToken(tokenString string) (*ExchangeClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ExchangeClaims{}, m.signingKeyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ExchangeClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid exchange token")
}

// signToken signs any jwt.Claims with the configured key.
func (m *JWTManager) signToken(claims jwt.Claims) (string, error) {
	if err := m.ensureConfig(); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := m.secretForClaims(claims)
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

// signingKeyFunc centralizes validation of the signing algorithm/key.
func (m *JWTManager) signingKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	if err := m.ensureConfig(); err != nil {
		return nil, err
	}

	return m.secretForClaims(token.Claims)
}

func (m *JWTManager) ensureConfig() error {
	if m == nil || m.cfg == nil {
		return errors.New("jwt manager is not configured")
	}
	return nil
}

func (m *JWTManager) secretForClaims(claims interface{}) ([]byte, error) {
	switch claims.(type) {
	case Claims, *Claims:
		if m.cfg.JwtAccessSecret == "" {
			return nil, errors.New("jwt access secret is not configured")
		}
		return []byte(m.cfg.JwtAccessSecret), nil
	case ExchangeClaims, *ExchangeClaims:
		if m.cfg.JwtExchangeSecret == "" {
			return nil, errors.New("jwt exchange secret is not configured")
		}
		return []byte(m.cfg.JwtExchangeSecret), nil
	default:
		return nil, errors.New("unknown claims type")
	}
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
