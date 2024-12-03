package utils

import "os"

// GetEnv retrieves and return environment variables. If none are configured, it returns a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
