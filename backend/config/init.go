package config

import (
	"os"
	"strconv"
)

// Config holds all the application configuration
type Config struct {
	Port                string
	BasePath            string
	MaxSegments         int
	SegmentDuration     float32
	LastSegmentDuration float32
	TotalSegments       int
}

// Global configuration instance
var AppConfig Config

func init() {
	// Set port
	AppConfig.Port = ":" + os.Getenv("PORT")
	if AppConfig.Port == ":" {
		AppConfig.Port = ":8080"
	}

	// Set base path
	AppConfig.BasePath = os.Getenv("BASE_PATH")
	if AppConfig.BasePath == "" {
		AppConfig.BasePath = "./hls test/"
	}

	// Set max segments
	maxSegmentsStr := os.Getenv("MAX_SEGMENTS")
	if maxSegmentsStr == "" {
		AppConfig.MaxSegments = 3
	} else {
		maxSegments, err := strconv.Atoi(maxSegmentsStr)
		if err != nil {
			AppConfig.MaxSegments = 3
		} else {
			AppConfig.MaxSegments = maxSegments
		}
	}

	// Set segment duration
	segmentDurationStr := os.Getenv("SEGMENT_DURATION")
	if segmentDurationStr == "" {
		AppConfig.SegmentDuration = 10.0
	} else {
		duration, err := strconv.ParseFloat(segmentDurationStr, 32)
		if err != nil {
			AppConfig.SegmentDuration = 10.0
		} else {
			AppConfig.SegmentDuration = float32(duration)
		}
	}

	// Set last segment duration
	lastSegmentDurationStr := os.Getenv("LAST_SEGMENT_DURATION")
	if lastSegmentDurationStr == "" {
		AppConfig.LastSegmentDuration = 4.566667
	} else {
		duration, err := strconv.ParseFloat(lastSegmentDurationStr, 32)
		if err != nil {
			AppConfig.LastSegmentDuration = 4.566667
		} else {
			AppConfig.LastSegmentDuration = float32(duration)
		}
	}

	// Set total segments
	totalSegmentsStr := os.Getenv("TOTAL_SEGMENTS")
	if totalSegmentsStr == "" {
		AppConfig.TotalSegments = 64
	} else {
		totalSegments, err := strconv.Atoi(totalSegmentsStr)
		if err != nil {
			AppConfig.TotalSegments = 64
		} else {
			AppConfig.TotalSegments = totalSegments
		}
	}
}
