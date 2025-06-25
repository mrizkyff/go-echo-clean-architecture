package config

type PostgresConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}
