package internal

import "os"

// TODO: functions should be moved to meaningful packages.
// Adding functions in this package should temporarly only
// for functions that can't be grouped as a separate package.

func GetEnv(name string, defaultValue ...string) string {
	value := os.Getenv(name)

	if value == "" {
		return defaultValue[0]
	}

	return value
}
