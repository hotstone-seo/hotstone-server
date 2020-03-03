package repository

import (
	"time"
)

// MetricsMismatchedCount represented  metrics_unmatched entity
type MetricsMismatchedCount struct {
	URL      string    `json:"url"`
	Count    int64     `json:"count"`
	LastSeen time.Time `json:"last_seen"`
}

type MetricsCountHitPerDay struct {
	Date  time.Time `json:"date"`
	Count int64     `json:"count"`
}
