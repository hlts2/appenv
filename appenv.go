package appenv

import (
	"os"
	"strings"
)

// AppEnv represents an application runtime environment.
type AppEnv uint8

const (
	// Unknown represents the unknown environment.
	Unknown AppEnv = iota
	// Test represents the test environment.
	Test
	// Dev represents the development environment.
	Dev
	// Stg represents the staging environment.
	Stg
	// Prod represents the production environment.
	Prod
)

// String returns string environment.
func (e AppEnv) String() string {
	switch e {
	case Test:
		return "test"
	case Dev:
		return "development"
	case Stg:
		return "staging"
	case Prod:
		return "production"
	default:
		return "unknown"
	}
}

// AtoE converts from string to AppEnv.
func AtoE(str string) AppEnv {
	str = strings.ToLower(str)
	switch str {
	case Test.String():
		return Test
	case Dev.String():
		return Dev
	case Stg.String():
		return Stg
	case Prod.String():
		return Prod
	default:
		return Unknown
	}
}

// Env gets the environment variable with the given key and returns the corresponding AppEnv.
func Env(key string) AppEnv {
	str := os.Getenv(key)
	return AtoE(str)
}
