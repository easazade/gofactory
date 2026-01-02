My Golang playground
# Notes

### General

1. Libraries are called modules
1. Directories are called packages
1. The root directory is called package main

# Commands

### Dependencies

To install dependencies copy the dependency path from https://pkg.go.dev/ then do like below:

- `go get com.foo/module@latest` // for latest version
- `go get com.foo/module/v4@v4.0.0` // for specific version

Add, Remove, Upgrade, Downgrade of dependencies is done by `go get` command. `go mod` handles tasks regarding modules/dependencies.


