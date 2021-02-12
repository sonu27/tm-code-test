# TM Code Test

## Running the app

Make sure you have golang installed. This code was tested with `go 1.15`.

First build the app: `go build ./cmd/app/main.go`

Then `./main [input file path]` e.g. `./main input-file.txt`

## Running the tests

Run `go test ./...`

## Check test coverage

Run `go test ./... -coverpkg=./... -coverprofile=coverage.out`

Then `go tool cover -html=coverage.out`

This should open a html page with the lines covered by tests and percentages per file.

## Future improvements

* Use interfaces
* Use mocks instead of concrete implementations when testing a separate domain  
* Add a logger, so can log errors or results elsewhere if needed
* Would consider using uint/uint64 for amounts/price for less memory usage, and slightly improved performance
* UserID could be a struct if more details are needed on each user
* More testing, especially of bootstrap and moving to its own package
* Mapping for the fields
* Stream data in from STDIN
