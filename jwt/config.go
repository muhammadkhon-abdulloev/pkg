package jwt

import "time"

type Config struct {
	Issuer          *string        `json:"issuer" mapstructure:"ISSUER"`
	Key             *string        `json:"key" mapstructure:"KEY"`
	AccessTokenExp  *time.Duration `json:"accessTokenExp" mapstructure:"ACCESS_TOKEN_EXP"`
	RefreshTokenExp *time.Duration `json:"refreshTokenExp" mapstructure:"REFRESH_TOKEN_EXP"`
}
