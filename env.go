// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package osenv

import "os"

// Get returns the environment variable associated with the given key.
// Returns the fallback value if not exist.
func Get(key string, fallback ...string) string {
	if len(fallback) == 0 {
		return os.Getenv(key)
	}
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	return fallback[0]
}

// SetNX sets the value of the environment variable named by the key if not exist.
func SetNX(key, value string) error {
	if _, exists := os.LookupEnv(key); !exists {
		return os.Setenv(key, value)
	}

	return nil
}
