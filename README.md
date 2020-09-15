# Extensional OS environment variables helper for Go.
[![Build Status](https://img.shields.io/travis/clevergo/osenv?style=flat)](https://travis-ci.org/clevergo/osenv)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/osenv?style=flat)](https://coveralls.io/github/clevergo/osenv)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/clevergo.tech/osenv?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/osenv?style=flat)](https://goreportcard.com/report/github.com/clevergo/osenv)
[![Release](https://img.shields.io/github/release/clevergo/osenv.svg?style=flat)](https://github.com/clevergo/osenv/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/osenv&style=flat)](https://pkg.clevergo.tech/clevergo.tech/osenv)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat)](https://forum.clevergo.tech)

## Usage

```shell
$ go get -u clevergo.tech/osenv
```

### Get

Gets the environment variable associated with the given key, a fallback(default) value will be returned if not exist.

Assume environment is that `FOO=BAR EMPTY=`.

| Method | Value
|---|:---:|
| `osenv.Get("FOO")` | BAR |
| `osenv.Get("EMPTY")` | - |
| `osenv.Get("EMPTY", "BUZZ")` | - |
| `osenv.Get("FIZZ")` | - |
| `osenv.Get("FIZZ", "BUZZ")` | BUZZ |

Similar functions:

- `GetInt`

### MustGet

Gets the environment variable associated with the given key, and panics if not exist.

```go
osenv.MustGet("FOO") // BAR
osenv.MustGet("NIL") // panics
```

Similar functions:

- `MustGetInt`

### SetNX

Sets an environment variable if not exist.

```go
osenv.SetNX("NONEXISTENT", "VALUE")
```
