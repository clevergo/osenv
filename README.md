# A Helper for retrieving OS environment variables.
[![Build Status](https://img.shields.io/travis/clevergo/osenv?style=for-the-badge)](https://travis-ci.org/clevergo/osenv)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/osenv?style=for-the-badge)](https://coveralls.io/github/clevergo/osenv)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/osenv?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/osenv?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/osenv)
[![Release](https://img.shields.io/github/release/clevergo/osenv.svg?style=for-the-badge)](https://github.com/clevergo/osenv/releases)

## Usage

```shell
$ go get -u clevergo.tech/osenv
```

### Retrieve

Assume environment is that `FOO=BAR EMPTY=`.

| Method | Value
|---|---|
| `osenv.Get("FOO")` | bar |
| `osenv.Get("EMPTY")` | - |
| `osenv.Get("EMPTY", "BUZZ")` | - |
| `osenv.Get("FIZZ")` | - |
| `osenv.Get("FIZZ", "BUZZ")` | BUZZ |
