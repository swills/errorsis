# ErrorsIs Linter

## Overview
ErrorsIs is a custom Go linter that detects incorrect usage of the
`errors.Is()` function, specifically targeting cases where the second argument
is a pointer to a struct that does not properly implement the `error`
interface.

## Features
- Detects improper usage of `errors.Is()` where the second argument is a
  pointer to a struct that fails to implement `error` directly.
- Supports type-safe matching of `errors.Is()` function calls, ensuring
  accurate analysis.

## Build and Installation
Ensure you have Go installed (version 1.23+).

```sh
go build -o errorsis ./cmd/errorsis
```

## Build as plugin for golangci-lint

Usage as a plugin will only work with compatible versions of golangci-lint (we
must use the same version of golang.org/x/tools as golangci-lint). This version
has been built and tested against golangci-lint v2.1.6 (built from source, not
sure about pre-built binaries). See the [example linter](https://github.com/golangci/example-plugin-linter/tree/1d4f00fda884c1928a9dbbfea865e7dc01e16477?tab=readme-ov-file#create-the-plugin-from-this-linter) for more details.

```sh
go build -buildmode=plugin plugin/errorsis.go
```

## Usage
To run the linter directly on Go code:

```sh
./errorsis ./...
```

## Usage with golangci-lint

Add something like this to your .golangci.yml

```yaml
    version: "2"
    linters:
      default: none
      enable:
        - errorsis
      settings:
        custom:
          errorsis:
            path: /path/to/git/errorsis/errorsis.so
            description: Detects incorrect usage of errors.Is
```
