package utils

import (
	"net/http"
	"strconv"
)

func ParseQueryInt(r *http.Request, key string, defaultVal int) int {
	// Get string value
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultVal
	}

	// Convert string to int
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return valInt
}

func ParseQueryString(r *http.Request, key string, defaultVal string) string {
	// Get string value
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	return val
}
