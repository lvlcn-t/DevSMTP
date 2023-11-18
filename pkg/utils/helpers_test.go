package utils

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	// test environment variable not set
	_, err := GetEnv("NON_EXISTENT_ENV", ParseString)
	if err == nil || err.Error() != "environment variable NON_EXISTENT_ENV is not set" {
		t.Errorf("GetEnv did not return expected error for non-existent environment variable")
	}

	// set an environment variable for testing
	t.Setenv("TEST_ENV", "TestValue")

	// test environment variable is set
	value, err := GetEnv("TEST_ENV", ParseString)
	if err != nil {
		t.Errorf("GetEnv returned an error for an existing environment variable: %v", err)
	} else if value != "TestValue" {
		t.Errorf("GetEnv returned wrong value: got %v want %v", value, "TestValue")
	}
}

func TestGetEnvInt(t *testing.T) {
	// test environment variable not set
	_, err := GetEnv("NON_EXISTENT_ENV", ParseInt)
	if err == nil || err.Error() != "environment variable NON_EXISTENT_ENV is not set" {
		t.Errorf("GetEnv did not return expected error for non-existent environment variable")
	}

	// set an environment variable for testing
	t.Setenv("TEST_ENV_INT", "123")

	// test environment variable is set and contains an integer
	value, err := GetEnv("TEST_ENV_INT", ParseInt)
	if err != nil {
		t.Errorf("GetEnv returned an error for an existing environment variable: %v", err)
	} else if value != 123 {
		t.Errorf("GetEnv returned wrong value: got %v want %v", value, 123)
	}

	// set an environment variable with non-integer value
	t.Setenv("TEST_ENV_NON_INT", "NonIntValue")

	// test environment variable is set and contains a non-integer
	_, err = GetEnv("TEST_ENV_NON_INT", ParseInt)
	if err == nil {
		t.Errorf("GetEnv did not return expected error for non-integer environment variable")
	}
}

func TestGetEnvBool(t *testing.T) {
	// test environment variable not set
	_, err := GetEnv("NON_EXISTENT_ENV", ParseBool)
	if err == nil || err.Error() != "environment variable NON_EXISTENT_ENV is not set" {
		t.Errorf("GetEnv did not return expected error for non-existent environment variable")
	}

	// set an environment variable for testing
	t.Setenv("TEST_ENV_BOOL", "false")

	// test environment variable is set and contains a boolean
	value, err := GetEnv("TEST_ENV_BOOL", ParseBool)
	if err != nil {
		t.Errorf("GetEnv returned an error for an existing environment variable: %v", err)
	} else if value != false {
		t.Errorf("GetEnv returned wrong value: got %v want %v", value, 123)
	}

	// set an environment variable with non-boolean value
	t.Setenv("TEST_ENV_NON_BOOL", "NonBoolValue")

	// test environment variable is set and contains a non-boolean
	_, err = GetEnv("TEST_ENV_NON_BOOL", ParseBool)
	if err == nil {
		t.Errorf("GetEnv did not return expected error for non-integer environment variable")
	}
}
