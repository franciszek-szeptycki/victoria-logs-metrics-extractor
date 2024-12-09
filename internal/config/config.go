package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	VictoriaLogsURL     string
	LogTimeframeMinutes int
}

func LoadConfig() Config {
	victoriaLogsURL := os.Getenv("VICTORIA_LOGS_URL")
	if victoriaLogsURL == "" {
		log.Fatal("Environment variable VICTORIA_LOGS_URL is required but not set")
	}

	logTimeframeMinutesStr := os.Getenv("LOG_TIMEFRAME_MINUTES")
	if logTimeframeMinutesStr == "" {
		log.Fatal("Environment variable LOG_TIMEFRAME_MINUTES is required but not set")
	}

	logTimeframeMinutes, err := strconv.Atoi(logTimeframeMinutesStr)
	if err != nil || logTimeframeMinutes <= 0 {
		log.Fatal("Environment variable LOG_TIMEFRAME_MINUTES must be a positive integer")
	}

	return Config{
		VictoriaLogsURL:     victoriaLogsURL,
		LogTimeframeMinutes: logTimeframeMinutes,
	}
}
