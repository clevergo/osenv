// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package osenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	key := "GET"
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

func TestMustGet(t *testing.T) {
	key := "MUST_GET"
	value := "MUST_GET VALUE"
	os.Setenv(key, value)
	assert.Equal(t, value, MustGet(key))

	assert.Panics(t, func() {
		MustGet("NIL")
	})
}

func TestSetNX(t *testing.T) {
	key := "SEX_NX"
	value := "SET_NX VALUE"
	assert.Equal(t, "", os.Getenv(key))

	assert.Nil(t, SetNX(key, value))
	assert.Equal(t, value, os.Getenv(key))

	assert.Nil(t, SetNX(key, "new value"))
	assert.Equal(t, value, os.Getenv(key))
}

func TestGetInt(t *testing.T) {
	key := "GET_INT"
	_, err := GetInt(key)
	assert.NotNil(t, err)

	fallback := 1
	v, _ := GetInt(key, fallback)
	assert.Equal(t, fallback, v)

	os.Setenv(key, "2")
	v, _ = GetInt(key)
	assert.Equal(t, 2, v)
}

func TestMustGetInt(t *testing.T) {
	key := "MUST_GET_INT"
	assert.Panics(t, func() {
		MustGetInt(key)
	})

	os.Setenv(key, "3")
	assert.Equal(t, 3, MustGetInt(key))
}
