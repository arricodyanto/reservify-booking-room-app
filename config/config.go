package config

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) ConfigConfiguration() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{ApiPort: os.Getenv("API_PORT")}
	// tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	// c.TokenConfig = TokenConfig{
	// 	IssuerName:       os.Getenv("TOKEN_ISSUE"),
	// 	JwtSignatureKey:  []byte(os.Getenv("TOKEN_SECRET")),
	// 	JwtSigningMethod: jwt.SigningMethodHS256,
	// 	JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	// }

	if c.Host == "" || c.Port == "" || c.User == "" || c.Name == "" || c.Driver == "" || c.ApiPort == "" {
		return fmt.Errorf("missing required environment")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ConfigConfiguration(); err != nil {
		return nil, err
	}
	return cfg, nil
}

