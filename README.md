# ErrorsIs Linter

[![Go Reference](https://pkg.go.dev/badge/github.com/swills/errorsis.svg)](https://pkg.go.dev/github.com/swills/errorsis)
[![Go Report Card](https://goreportcard.com/badge/github.com/swills/errorsis)](https://goreportcard.com/report/github.com/swills/errorsis)

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
go build -trimpath -o errorsis ./cmd/errorsis
```

## Build with golangci-lint

Note, `golangci-lint` supports two different ways of building and using plugins.

One is via a [go plugin](https://golangci-lint.run/plugins/go-plugins/). This will only work
with compatible versions of golangci-lint (we must use the same version of `golang.org/x/tools` as
golangci-lint). This version has been built and tested against golangci-lint v2.1.6 (built from
source, not sure about pre-built binaries). See the [example linter](https://github.com/golangci/example-plugin-linter/tree/1d4f00fda884c1928a9dbbfea865e7dc01e16477?tab=readme-ov-file#create-the-plugin-from-this-linter) for more details. To
build for usage as a Go plugin:

```sh
go build -trimpath -buildmode=plugin -o errorsis.so plugin/errorsis.go
```

The other is via its own [module plugin system](https://golangci-lint.run/plugins/module-plugins/). See
the [example module linter](https://github.com/golangci/example-plugin-module-linter) for more details.

This package supports both, but note that the latter is preferred and recommended by `golangci-lint`.

## Usage
To run the linter standalone directly on Go code:

```sh
./errorsis ./...
```

## Usage with golangci-lint as Go plugin:

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

## Usage with golangci-lint as Module plugin:

Create `.custom-gcl.yml` similar to this if using a local working copy:

```yaml
version: v2.1.6
plugins:
  - module: "github.com/swills/errorsis"
    path: "/path/to/git/errorsis"
```

or like this if using Go proxy:

```yaml
version: v2.1.6
plugins:
  - module: 'github.com/swills/errorsis'
    import: 'github.com/swills/errorsis'
    version: v0.0.5
```

And put something similar to this in your `.golangci.yml`:

```yaml
version: "2"
linters:
  default: none
  enable:
    - errorsis
  settings:
    custom:
      errorsis:
        type: module
        description: Detects incorrect usage of errors.Is
```

Then run:

```shell
golangci-lint -v custom
```

to build the custom `golangci-lint` and finally, run it:

```shell
./custom-gcl run
```
