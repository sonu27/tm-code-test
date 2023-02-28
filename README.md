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

* Use money value objects instead of float64 for amounts/price if calculations are needed in the future
* Add a logger, so can log errors or results elsewhere if needed
* Mapping for the fields
