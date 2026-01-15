package config

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBPort              string
	ServerPort          string
	JWTSecret           string
	JWTExpirationHours  int
}

func LoadConfig() *Config {
	godotenv.Load()
	
	jwtExpHours, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if jwtExpHours == 0 {
		jwtExpHours = 24 // default 24 hours
	}
	
	return &Config{
		DBHost:             os.Getenv("DB_HOST"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		DBPort:             os.Getenv("DB_PORT"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTExpirationHours: jwtExpHours,
	}
}