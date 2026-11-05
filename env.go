package env

import (
	"os"
	"strconv"
	"strings"
)

func Get(key string, defaultValue string) string {
	value := GetAsString(key, defaultValue)
	return value
}

func GetAsString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetAsInt64(key string, defaultValue int64) int64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func GetAsBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value == "true"
}

func GetAsInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func GetAsFloat(key string, defaultValue float64) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

func GetAsSlice(key string, defaultValue []string, sep string) []string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	valueTrimmed := []string{}
	for _, v := range strings.Split(value, sep) {
		v = strings.TrimSpace(v)
		if v != "" {
			valueTrimmed = append(valueTrimmed, v)
		}
	}
	return valueTrimmed
}
