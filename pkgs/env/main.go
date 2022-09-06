package env

import "os"

// getEnv gets the ENV key with fallback if it's not found
func Get(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
