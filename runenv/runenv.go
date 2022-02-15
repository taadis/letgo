package runenv

import (
	"os"
	"strings"
)

var runEnvKey = "LET_RUN_ENV"

type RunEnv = string

const (
	Development RunEnv = "development"
	Testing     RunEnv = "testing"
	Gray        RunEnv = "gray"
	Production  RunEnv = "production"
)

// GetKey the runtime environment key
func GetKey() string {
	return runEnvKey
}

// Get the current runtime environment
func Get() RunEnv {
	env := os.Getenv(runEnvKey)
	if env == "" {
		// return default env if env empty
		env = Development
	}

	return strings.ToLower(env)
}

// Set the runtime environment
func Set(env RunEnv) error {
	value := strings.ToLower(env)
	return os.Setenv(runEnvKey, value)
}

func Is(env RunEnv) bool {
	s := Get()
	suffix := strings.ToLower(env)
	return strings.HasSuffix(s, suffix)
}

func Not(env RunEnv) bool {
	return !Is(env)
}
