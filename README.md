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
1. Types, functions, fields, structs that are capitalized are public and lowercased ones are private
1. Methods are just functions with a receiver. No inheritance. No `this` keyword
1. Interfaces describe behavior. Structs happen to satisfy it. No `implements` keyword
1. Types don’t declare what they are. They prove it by behavior
1. exported/unexported are used instead of public/private because it's actually package-level visibility not object-level

### Run/Build/Compilation

1. Same package name in different directories = different packages. Only files in the same directory get compiled together.
1. Compilation happens per package (per directory) so if multiple directories have same package name (eg: main) they are actually different packages and they will not merge in compilation and different binaries will be produced

### Versioning

Go uses git or other version control tools to determine the version of the module. In case of the git it is git tags. the tag names should be like `v1.2.0` or `v1.2.0-beta.1`
To bump the version on a module:

run: `git tag v1.1.0`
After bumping to version 2 or higher the path of module must change appending the `/v2` at the end. Do this only if you are bumping the major version
`eg: github.com/easazade/mygomodule/v2`
Then run: `git push --tags`

If you forget to change the path for major version bump before pushing the tags you need to update path and create a new tag like `v2.0.1`
and also to avoid users making a mistake getting the wrong version
Add a retract directive in `go.mod`

```
retract v2.0.0 // published with incorrect module path
```

After you push tags to your **public** repository it will be picked up by golang proxies which are servers watching a version control systems like github for example and downloads modules caches them immutably by version serves them back to Go tools in a standardized way to eg `go` cli fetches its dependencies from the official go proxy server at `proxy.golang.org` the `pkg.go.dev` website also uses that server as the source of data as well.

**NOTE:** It is possible to define our own golang proxy for teams and orgs
**NOTE:** Go also runs a checksum public service at `sum.golang.org` which contains all checksums for every public module version

##### Github Reference

However it is possible to use modules from private repositories that we have access to in our projects just tell Go that those module paths are private so it won’t use the public proxy/checksum DB for them. run command (one time only):
`go env -w GOPRIVATE="github.com/easazade/*"`

# Commands

### Init new Project/Module

run `go mod init github.com/easazade/goback`

The convention for the unique path used for the module:

1. Personal projects: `github.com/you/project`
1. Companies: `company.tld/team/service`
1. Throwaway code: `foo/bar`

You can always manually change the module path in the `go.mod` file. Go is very tolerant of manually changing the go.mod file. Even though for somethings like managing dependencies you should use the `go mod` command.

`god mod init` actually enables dependency tracking and everything module related. It is possible to just run a go files without a `go.mod` file
But you won't get very far. You need to only use standard library (stdlib) and everything needs to be in the main package

### Dependencies

##### Install

To install dependencies copy the dependency path from https://pkg.go.dev/ then do like below:

- `go get com.foo/module@latest` // for latest version
- `go get com.foo/module/v4@v4.0.0` // for specific version

Add, Remove, Upgrade, Downgrade of dependencies is done by `go get` command.

`go mod` handles tasks regarding modules/dependencies.

`go mod tidy` adds missing and removes unused modules. It is possible to just add an import into a file and just call `tidy` command to handle the rest.

##### Local dependencies

To install use local modules, it is required to replace the module path name with the actual local directory path of the module by running:
`go mod edit -replace example.com/module=local/path/to/module`
This will add a replace directive alongside the require directive to go.mod file

### Run

To run code simply use:

- `go run .` // finds the main function inside package (current dir)
- `go run file.go` // finds the main function in file.go
- `go run ./cmd/server`

**NOTE:** Only files that belong to package main will run.

### ENV

Go has its own environment values separated from shell env
To list all go env values
`go env`

To check the value of one variable:
`go env GOPRIVATE`
`go env GOPROXY`
