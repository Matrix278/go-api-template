package configuration

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Env struct {
	AppHost          string
	AppPort          string
	APIPath          string
	AppEnv           string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string
	AllowedOrigins   []string
}

func Load() (*Env, error) {
	if appName := viper.GetString("APP_ENV"); appName != "production" {
		viper.SetConfigFile(".env")
	}

	if err := viper.ReadInConfig(); err == nil {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	viper.SetDefault("POSTGRES_SSL_MODE", "disable")

	viper.AutomaticEnv()

	allowedOrigins := viper.GetString("ALLOWED_ORIGINS")
	origins := strings.Split(allowedOrigins, ",")

	return &Env{
		AppHost:          viper.GetString("APP_HOST"),
		AppPort:          viper.GetString("APP_PORT"),
		APIPath:          viper.GetString("API_PATH"),
		AppEnv:           viper.GetString("APP_ENV"),
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
		PostgresDB:       viper.GetString("POSTGRES_DB"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresSSLMode:  viper.GetString("POSTGRES_SSL_MODE"),
		AllowedOrigins:   origins,
	}, nil
}
