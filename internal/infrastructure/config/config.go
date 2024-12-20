package config

import (
	"log"
	"os"
	"strconv"
)

var (
	VictoriaLogsURL     string
	LogTimeframeMinutes int
	ErrorThreshold      float32
)

func Init() {
	VictoriaLogsURL = loadVictoriaLogsURL()
	LogTimeframeMinutes = loadLogTimeframeMinutes()
	ErrorThreshold = loadErrorThreshold()
}

func loadVictoriaLogsURL() string {
	victoriaLogsURL := os.Getenv("VICTORIA_LOGS_URL")
	if victoriaLogsURL == "" {
		log.Fatal("Environment variable VICTORIA_LOGS_URL is required but not set")
	}
	return victoriaLogsURL
}

func loadLogTimeframeMinutes() int {
	logTimeframeMinutesStr := os.Getenv("LOG_TIMEFRAME_MINUTES")
	if logTimeframeMinutesStr == "" {
		log.Fatal("Environment variable LOG_TIMEFRAME_MINUTES is required but not set")
	}

	logTimeframeMinutes, err := strconv.Atoi(logTimeframeMinutesStr)
	if err != nil || logTimeframeMinutes <= 0 {
		log.Fatal("Environment variable LOG_TIMEFRAME_MINUTES must be a positive integer")
	}
	return logTimeframeMinutes
}

func loadErrorThreshold() float32 {
	errorThresholdStr := os.Getenv("ERROR_THRESHOLD")
	if errorThresholdStr == "" {
		log.Fatal("Environment variable ERROR_THRESHOLD is required but not set")
	}

	errorThreshold, err := strconv.ParseFloat(errorThresholdStr, 32)
	if err != nil {
		log.Fatal("Environment variable ERROR_THRESHOLD must be a float")
	}
	return float32(errorThreshold)
}
