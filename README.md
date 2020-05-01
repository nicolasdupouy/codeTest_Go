# README

The goal of this repository is to store all the codes I want to test in Go

# Go
## Run

    $ go run ./hello/main.go

## Import dependencies
Get the dependencies imported in the code but missing in the path (it will update the *go.mod* file).

    $ go mod tidy


## Build
    $ go build -o bin/ ./hello/
    
or

    $ go build -o bin/ ./...
    
### Cross-compilation
Get the available OS/architecture combinations:

    $ go tool dist list

Compile for MacOS when on Debian Linux
    
    $ GOOS=darwin GOARCH=amd64 go build -o bin/ ./...
    
Or compile for Linux amd64:

    $ GOOS=linux GOARCH=amd64 go build -o bin/ ./...
    
## Test
Test a package

    $ go test ./gopher
    
Test all packages

    $ go test ./...
    
