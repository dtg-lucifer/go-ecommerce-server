package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
  JWT_SECRET string
  JWT_EXP    int64
}

var Env Config = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("MYSQL_USER", "root"),
		DBPassword: getEnv("MYSQL_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("MYSQL_DATABASE", "mydb"),
    JWT_SECRET: getEnv("JWT_SECRET", "SECRET_TEXT"),
    JWT_EXP: getEnvInt("JWT_EXP", 3600 * 24 * 7),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
    i, err := strconv.Atoi(value)
    if err != nil {
      return fallback
    }
		return int64(i)
	}
	return fallback
}
