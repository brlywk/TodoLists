package utils

import (
	"log"
	"os"
	"time"
)

// Helper function that tries to fetch an  env var, and if it fails, returns fallback
func GetEnv(key string, fallback string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return fallback
	}
	return value
}


// Helper measuring the time a function (or request) takes. Needs to be called deferred at
// the top of a function
func Measure(name string, method string) func() {
	start := time.Now()

	if method == "" {
		method = "GET"
	}

	return func() {
		log.Printf("\tPath: %s\t\tMethod: %s\tTime: %v", name, method, time.Since(start))
	}
}
