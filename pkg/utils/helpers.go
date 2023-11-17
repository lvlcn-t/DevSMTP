package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key string) (string, error) {
	value, set := os.LookupEnv(key)
	if !set {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}

func GetEnvInt(key string) (int, error) {
	valueStr, err := GetEnv(key)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s is not set to an int number", key)
	}
	return value, nil
}
