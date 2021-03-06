package lockotron

import (
	"time"
)

const (
	NoTTL     time.Duration = -1
	NoCleanup time.Duration = -1
)

type Config struct {
	DefaultTTL      time.Duration
	CleanupInterval time.Duration
}

func NewConfig() *Config {
	return &Config{DefaultTTL: 15 * time.Minute, CleanupInterval: 10 * time.Minute}
}
