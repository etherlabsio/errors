# errors 

[![Travis-CI](https://travis-ci.org/etherlabsio/errors.svg)](https://travis-ci.org/etherlabsio/errors) [![Go Report Card](https://goreportcard.com/badge/github.com/etherlabsio/errors)](https://goreportcard.com/report/github.com/etherlabsio/errors) [![GoDoc](https://godoc.org/github.com/etherlabsio/errors?status.svg)](https://godoc.org/github.com/etherlabsio/errors) [![codecov](https://codecov.io/gh/etherlabsio/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/etherlabsio/errors)

An errors package built on top of existing battle tested packages with structured JSON serialization support.

## Requirements

Go v1.10+ since `strings.Builder` API is used for generating the output for `Error()` method.

## Features

* A very primitive support for monads
* Maintains backwards compatibility with Dave Cheney's errors package.
* Additional functions for wrapping errors inspired by Upspin's error package.
* Additional methods for error inspection.
* JSON serialization support.

## Prior Art

* [Dave Cheney's error package](https://github.com/pkg/errors).
* [Upspin Error package](https://github.com/upspin/upspin/blob/master/errors)
