My Golang playground

# Notes

### General

1. Libraries are called modules
1. Directories are called packages
1. A package is actually a file
1. The root directory is called package main
1. Each file should specify which package it belongs to at the top
1. All files in the same package/directory are compiled together as if they were concatenated
1. All the files in a directory should have the same package name
1. Compilation happens per package (per directory) so if multiple directories have same package name (eg: main) they are actually different packages and they will not merge in compilation and different binaries will be produced

# Commands

### Init new Project/Module

run `go mod init github.com/easazade/goback`

The convention for the unique path used for the module:

1. Personal projects: `github.com/you/project`
1. Companies: `company.tld/team/service`
1. Throwaway code: `foo/bar`

### Dependencies

To install dependencies copy the dependency path from https://pkg.go.dev/ then do like below:

- `go get com.foo/module@latest` // for latest version
- `go get com.foo/module/v4@v4.0.0` // for specific version

Add, Remove, Upgrade, Downgrade of dependencies is done by `go get` command.

`go mod` handles tasks regarding modules/dependencies.

- `go mod tidy` adds missing and removes unused modules

### Run

To run code simply use:

- `go run .` // finds the main function inside package (current dir)
- `go run file.go` // finds the main function in file.go
- `go run ./cmd/server`

**NOTE:** Only files that belong to package main will run.
