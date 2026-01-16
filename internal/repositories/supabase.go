package repositories

import (
	"errors"
	"os"

	supabase "github.com/supabase-community/supabase-go"
)

type SupabaseClient struct {
	Client *supabase.Client
}

func NewSupabaseClient() (*SupabaseClient, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		return nil, errors.New("SUPABASE_URL is required")
	}

	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	if supabaseKey == "" {
		supabaseKey = os.Getenv("SUPABASE_ANON_KEY")
	}
	if supabaseKey == "" {
		return nil, errors.New("SUPABASE_SERVICE_ROLE_KEY or SUPABASE_ANON_KEY is required")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, err
	}

	return &SupabaseClient{Client: client}, nil
}
