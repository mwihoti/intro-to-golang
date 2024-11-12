# Intro to Golang

This repository contains a simple Go project that demonstrates basic Go functionalities, including package management, vendoring, and basic functions.


## Files

- **go.mod**: The Go module file that defines the module path and dependencies.
- **main.go**: The main entry point of the Go application.
- **goFun/goFun.go**: A package that contains utility functions used in the main application.
- **vendor/**: Directory containing vendored dependencies.


## SetUp


1. **Initialize the module** (if not already initialized):
    ```sh
    go mod init stackup.dev/intro-to-golang
    ```

2. **Enable vendoring**:
    ```sh
    go mod vendor
    ```

3. **Run the application**:
    ```sh
    go run -mod=vendor main.go
    ```

## Functions

### main.go

- **is_legal_age(age int) (msg string, legal bool)**: Determines if the given age is of legal age (18 or older).
- **is_rich(rich bool) string**: Returns a message based on the boolean value indicating if the person is rich.
- **forLoop(c string) string**: Iterates over the characters of the string "Daniel" and prints each character.

### goFun/goFun.go

- **hello(name string) string**: Returns a greeting message for the given name.
- **Add(a, b int) int**: Returns the sum of two integers.
