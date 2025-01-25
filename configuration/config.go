package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

type New struct {
	AppHost          string
	AppPort          string
	APIPath          string
	AppEnv           string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	AllowedOrigins   []string
}

func Load() (*New, error) {
	if APP_ENV := viper.GetString("APP_ENV"); APP_ENV != "production" {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()

	allowedOrigins := viper.GetString("ALLOWED_ORIGINS")
	origins := strings.Split(allowedOrigins, ",")

	return &New{
		AppHost:          viper.GetString("APP_HOST"),
		AppPort:          viper.GetString("APP_PORT"),
		APIPath:          viper.GetString("API_PATH"),
		AppEnv:           viper.GetString("APP_ENV"),
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
		PostgresDB:       viper.GetString("POSTGRES_DB"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		AllowedOrigins:   origins,
	}, nil
}
