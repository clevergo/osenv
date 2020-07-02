// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package osenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	key := "FOO"
	assert.Equal(t, "", Get(key))

	for _, fallback := range []string{"", "FIZZ", "BUZZ"} {
		assert.Equal(t, fallback, Get(key, fallback))
	}

	value := "BAR"
	os.Setenv(key, value)
	assert.Equal(t, value, Get(key))
	assert.Equal(t, value, Get(key, "FALLBACK"))

	os.Setenv("EMPTY", "")
	assert.Equal(t, "", Get("EMPTY", "NOT EMPTY"))
}