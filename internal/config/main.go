package config

import (
	"golang.org/x/oauth2/google"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
	//AuthGoogle GoogleOAuthConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
}

type AuthConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

type ServerConfig struct {
	Environment    string
	Port           string
	AllowedOrigins []string
	BaseURL        string
}

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_ID_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
			return
		}

		cfg = &Config{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Name:     os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				SSLMode:  "", // this field is set depending on the server environment
			},
			Auth: AuthConfig{
				AccessTokenSecret:  os.Getenv("JWT_ACCESS_SECRET"),
				RefreshTokenSecret: os.Getenv("JWT_REFRESH_SECRET"),
			},
			Server: ServerConfig{
				Environment: os.Getenv("SERVER_ENVIRONMENT"),
				Port:        os.Getenv("PORT"),
				AllowedOrigins: []string{
					os.Getenv("MOBILE_CLIENT_ORIGIN"),
					os.Getenv("WEBSITE_CLIENT_ORIGIN"),
					os.Getenv("ADMIN_PORTAL_ORIGIN"),
				},
				BaseURL: os.Getenv("SERVER_BASE_URL"),
			},
		}
	})

	if err != nil {
		return nil, err
	}

	return cfg, nil
}