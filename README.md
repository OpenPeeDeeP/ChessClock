# Chess Clock

## Installing

- `brew tap openpeedeep/openpdp`
- `brew install openpeedeep/openpdp/chessclock`
- `brew services start chessclockd`
- `clock --version`

## Building locally

- `go get -u github.com/golang/dep/cmd/dep`
- `dep ensure`
- `make build`
