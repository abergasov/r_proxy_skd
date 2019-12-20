package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerUrl    string
	BotToken     string
	Chat         string
	DelayTimeout int
	Version      string
}

var conf *Config

func NewAppConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	var delay, _ = strconv.Atoi(getEnv("REQUEST_DELAY", "5"))
	conf = &Config{
		ServerUrl:    getEnv("SERVER_URL", ""),
		BotToken:     getEnv("BOT_TOKEN", ""),
		Chat:         getEnv("CHAT_ID", ""),
		DelayTimeout: delay,
		Version:      "0.1",
	}
	return conf
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
