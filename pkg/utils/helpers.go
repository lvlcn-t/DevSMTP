package utils

import (
	"fmt"
	"os"
	"strconv"
)

// GetEnv retrieves an environment variable as a specified type.
// It requires the key of the environment variable and a parseFunc,
// which is a function to convert the string value to the desired type.
// If the environment variable is not set or cannot be converted,
// an appropriate error is returned.
//
// Parameters:
// - key: The key of the environment variable.
// - parseFunc: A function that takes a string and returns a value of the desired type and an error.
//
// Returns:
// - The converted value of the environment variable.
// - An error if the variable is not set or cannot be converted.
//
// Example usage:
//
//	intValue, err := utils.GetEnv("MY_INT_ENV", utils.ParseInt)
//	boolValue, err := utils.GetEnv("MY_BOOL_ENV", utils.ParseBool)
//	stringValue, err := utils.GetEnv("MY_STRING_ENV", utils.ParseString)
func GetEnv[T any](key string, parseFunc func(string) (T, error)) (T, error) {
	var zeroValue T

	value, set := os.LookupEnv(key)
	if !set {
		return zeroValue, fmt.Errorf("environment variable %s is not set", key)
	}

	parsedValue, err := parseFunc(value)
	if err != nil {
		return zeroValue, fmt.Errorf("error parsing environment variable %s: %v", key, err)
	}

	return parsedValue, nil
}

// ParseString is a helper function for GetEnv to handle string values.
// It returns the input string as is without any conversion.
//
// Parameters:
// - s: The string to parse.
//
// Returns:
// - The input string.
// - nil as there is no error in this operation.
func ParseString(s string) (string, error) {
	return s, nil
}

// ParseInt is a helper function for GetEnv to handle integer values.
// It converts a string to an integer.
//
// Parameters:
// - s: The string to convert to an integer.
//
// Returns:
// - The converted integer value.
// - An error if the conversion fails.
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseBool is a helper function for GetEnv to handle boolean values.
// It converts a string to a boolean.
//
// Parameters:
// - s: The string to convert to a boolean.
//
// Returns:
// - The converted boolean value.
// - An error if the conversion fails.
func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
