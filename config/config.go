package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}
type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}
type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

func (c Config) readConfig() Config {
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
		ApiHost: os.Getenv("API_HOST"),
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("JWT_APPNAME"),
		JwtSignatureKey:     os.Getenv("JWT_KEY"),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: 60 * time.Second,
	}
	return c
}

func NewConfig() Config {
	conf := Config{}
	return conf.readConfig()
}
