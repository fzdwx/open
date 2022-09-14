package env

import "os"

// Or get or default value
func Or(key, def string) string {
	return OrWithFunc(key, func() string {
		return def
	})
}

// OrWithFunc get env by key, if is empty use default val
func OrWithFunc(key string, def func() string) string {
	v := os.Getenv(key)

	if v == "" {
		return def()
	}

	return v
}

// Set key val to env
func Set(key, val string) error {
	return os.Setenv(key, val)
}
