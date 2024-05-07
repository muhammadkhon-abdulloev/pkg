package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	issuer          = "TEMP_ISSUER"
	key             = "random_key_12345"
	accessTokenExp  = time.Minute * 15
	refreshTokenExp = time.Hour * 24
)

func Init(c *Config) {
	if c == nil {
		return
	}

	if c.Issuer != nil && *c.Issuer != "" {
		issuer = *c.Issuer
	}

	if c.Key != nil && *c.Key != "" {
		key = *c.Key
	}

	if c.AccessTokenExp != nil {
		accessTokenExp = *c.AccessTokenExp
	}

	if c.RefreshTokenExp != nil {
		refreshTokenExp = *c.RefreshTokenExp
	}
}

func GeneratePairs(userID string) (Tokens, error) {
	accessToken, err := generate(userID, accessTokenExp).SignedString([]byte(key))
	if err != nil {
		return Tokens{}, err
	}

	refreshToken, err := generate(userID, refreshTokenExp).SignedString([]byte(key))
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generate(userID string, exp time.Duration) *jwt.Token {
	issuedAt := time.Now()
	expiration := issuedAt.Add(exp)

	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		ID:        uuid.New().String(),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiration),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func Verify(token string) (jwt.Claims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, err
	}

	return jwtToken.Claims, nil
}
