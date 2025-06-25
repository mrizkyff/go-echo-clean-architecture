package config

type JWTConfig struct {
	SecretKey        string
	ExpirationHours  int
	RefreshTokenDays int
	Issuer           string
}
