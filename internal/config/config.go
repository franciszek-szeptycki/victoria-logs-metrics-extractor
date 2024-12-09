package config

import (
	"log"
	"main/internal/constants"
	"os"
	"strconv"
)

type Config struct {
	VictoriaLogsURL     string
	LogTimeframeMinutes int
}

func LoadConfig() Config {
	victoriaLogsURL := os.Getenv(constants.VictoriaLogsURLEnvVar)
	if victoriaLogsURL == "" {
		log.Fatal("Environment variable VICTORIA_LOGS_URL is required but not set")
	}

	logTimeframeMinutesStr := os.Getenv(constants.LogTimeframeMinutesEnvVar)
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
