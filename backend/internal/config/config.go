package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DatabaseURL        string
	WebURL             string
	JWTSecret          string
	SecretKey          string
	GoogleClientID     string
	GoogleClientSecret string
	IsProd             bool
	BaseUrl            string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Port:               getEnv("PORT", "9000"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		WebURL:             getEnv("WEB_URL", "http://localhost:3000"),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		SecretKey:          getEnv("SECRET_KEY", ""),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		IsProd:             getEnvBool("IS_PROD", false),
		BaseUrl:            getEnv("BASE_URL", "http://localhost:"+getEnv("PORT", "9000")),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			log.Printf("invalid boolean value for %s: %v, using fallback", key, err)
			return fallback
		}
		return parsed
	}
	return fallback
}
