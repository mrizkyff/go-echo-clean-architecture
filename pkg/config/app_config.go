package config

import (
	"go-echo-clean-architecture/internal/models/config"

	"github.com/spf13/viper"
)

type AppConfig struct {
	PostgresConfig config.PostgresConfig
	RabbitMqConfig config.RabbitMQConfig
	JwtConfig      config.JWTConfig
	RedisConfig    config.RedisConfig
}

func NewAppConfig(v *viper.Viper) *AppConfig {
	return &AppConfig{
		PostgresConfig: config.PostgresConfig{
			Host:     v.GetString("postgres.host"),
			Username: v.GetString("postgres.username"),
			Password: v.GetString("postgres.password"),
			DBName:   v.GetString("postgres.dbname"),
			Port:     v.GetString("postgres.port"),
		},
		RabbitMqConfig: config.RabbitMQConfig{
			Username: v.GetString("rabbitMq.username"),
			Password: v.GetString("rabbitMq.password"),
			Host:     v.GetString("rabbitMq.host"),
			Port:     v.GetString("rabbitMq.port"),
			Vhost:    v.GetString("rabbitMq.vhost"),
		},
		JwtConfig: config.JWTConfig{
			SecretKey:        v.GetString("jwtConfig.secret_key"),
			ExpirationHours:  v.GetInt("jwtConfig.expirationHours"),
			RefreshTokenDays: v.GetInt("jwtConfig.refreshTokenDays"),
			Issuer:           v.GetString("jwtConfig.issuer"),
		},
		RedisConfig: config.RedisConfig{
			Host:     v.GetString("redis.host"),
			Port:     v.GetInt("redis.port"),
			Password: v.GetString("redis.password"),
			DbIndex:  v.GetInt("redis.dbIndex"),
		},
	}
}
