package config

import (
	"os"
	"sync"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/app/logger"
)

type Config struct {
	Judge0API         string
	PGURL             string
	AUTH_API          string
	Judge0CallbackURL string
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	once.Do(load)
	return cfg
}

func load() {
	cfg = &Config{
		Judge0API:         must("JUDGE0_API"),
		PGURL:             must("PG_URL"),
		AUTH_API:          must("AUTH_API"),
		Judge0CallbackURL: must("JUDGE0_CALLBACK_URL"),
	}
}

func must(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logger.Log.Fatal("missing env var", "key", key)
	}
	return val
}
