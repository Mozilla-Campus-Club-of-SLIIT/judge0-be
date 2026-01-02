package db

import (
	"log"
	"sync"

	"github.com/Mozilla-Campus-Club-of-SLIIT/judge0-be/internal/config"
	supabase "github.com/supabase-community/supabase-go"
)

var (
	Client *supabase.Client
	once   sync.Once
)

func Init() {
	once.Do(func() {
		cfg := config.Get()

		if cfg.SupabaseURL == "" || cfg.SupabaseServiceKey == "" {
			log.Fatal("Supabase env vars missing")
		}

		c, err := supabase.NewClient(
			cfg.SupabaseURL,
			cfg.SupabaseServiceKey,
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		Client = c
	})
}
