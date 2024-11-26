package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		MongoURI: getEnv("MONGO_URI", "mongodb+srv://admin:admin@serverlessinstance0.wxsf1y7.mongodb.net/?retryWrites=true&w=majority&appName=ServerlessInstance0"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
