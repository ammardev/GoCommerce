package app

import "os"

func GetEnv(name string, defaultValue ...string) string {
	value := os.Getenv(name)

	if value == "" {
		return defaultValue[0]
	}

	return value
}
