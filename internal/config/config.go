package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL        string
	SupabaseAnonKey    string
	SupabaseServiceKey string
	SecretKey          string
	Judge0aAPI         string
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
	if os.Getenv("VERCEL") == "" {
		_ = godotenv.Load()
		log.Println("loaded .env file")

	} else {
		log.Println("Running on Vercel, skipping .env file")
	}

	cfg = &Config{
		SupabaseURL:        must("SUPABASE_URL"),
		SupabaseAnonKey:    must("SUPABASE_ANON_KEY"),
		SupabaseServiceKey: must("SUPABASE_SERVICE_KEY"),
		SecretKey:          must("SECRET_KEY"),
		Judge0aAPI:         must("JUDGE0_API"),
	}
}

func must(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing env var: %s", key)
	}
	return val
}
