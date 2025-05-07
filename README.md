# ErrorsIs Linter

## Overview
ErrorsIs is a custom Go linter that detects incorrect usage of the `errors.Is()` function, specifically targeting cases where the second argument is a pointer to a struct that does not properly implement the `error` interface.

## Features
- Detects improper usage of `errors.Is()` where the second argument is a pointer to a struct that fails to implement `error` directly.
- Supports type-safe matching of `errors.Is()` function calls, ensuring accurate analysis.

## Installation
Ensure you have Go installed (version 1.19+).

```sh
# Clone the repository
git clone <repo-url>

# Build the linter
cd errorsis
go build -o errorsis ./cmd/errorsis
```

## Usage
To run the linter directly on Go code:

```sh
./errorsis ./...
```

## Development
This project is structured using the following layout:

- `cmd/errorsis/`: Command-line entry point for the linter.
- `pkg/analyzer/`: Contains the core linter logic.
- `testdata/`: Contains test cases used for automated testing.
