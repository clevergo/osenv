// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package osenv

import (
	"fmt"
	"os"
	"strconv"
)

// Get returns the environment variable associated with the given key.
// Returns the fallback value if not exists.
func Get(key string, fallback ...string) string {
	if len(fallback) == 0 {
		return os.Getenv(key)
	}
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	return fallback[0]
}

// MustGet returns the environment variable associated with the given key,
// and panics if the key is not exists.
func MustGet(key string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}
	panic(fmt.Errorf("getting an unknown environment variable %q", key))
}

// SetNX sets the value of the environment variable named by the key if not exist.
func SetNX(key, value string) error {
	if _, exists := os.LookupEnv(key); !exists {
		return os.Setenv(key, value)
	}

	return nil
}

// GetInt converts the environment variable associated with the given key to int
// and returns it.
// Returns the fallback value if the key is not exists or the value is invalid.
func GetInt(key string, fallback ...int) (int, error) {
	if s, exists := os.LookupEnv(key); exists {
		return strconv.Atoi(s)
	}

	if len(fallback) > 0 {
		return fallback[0], nil
	}

	return 0, noEnvError(key)
}

// MustGetInt converts the environment variable associated with the given key to int
// and returns it.
// Panics if the key is not exists or the value is invalid.
func MustGetInt(key string) int {
	v, err := GetInt(key)
	if err != nil {
		panic(err)
	}

	return v
}

func noEnvError(key string) error {
	return fmt.Errorf("no such environment variable: %s", key)
}
