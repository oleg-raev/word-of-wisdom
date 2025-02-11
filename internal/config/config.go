package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerAddr          string
	ChallengeComplexity int
	QuotesFile          string
}

func LoadConfig() *Config {
	complexity, err := strconv.Atoi(os.Getenv("CHALLENGE_COMPLEXITY"))
	if err != nil || complexity < 1 || complexity > 65 {
		log.Fatal("Invalid complexity value")
	}

	return &Config{
		ServerAddr:          os.Getenv("SERVER_ADDR"),
		ChallengeComplexity: complexity,
		QuotesFile:          os.Getenv("QUOTES_FILE"),
	}
}
