package jwt

import "time"

type Config struct {
	Issuer          *string `json:"issuer" env:"ISSUER"`
	Key             *string `json:"key" env:"KEY"`
	AccessTokenExp  *time.Duration
	RefreshTokenExp *time.Duration
}
