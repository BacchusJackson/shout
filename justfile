INSTALL_DIR := env('INSTALL_DIR', '/usr/local/bin')

# build gatecheck binary
build:
    mkdir -p bin
    go build -ldflags="-X 'main.cliVersion=$(git describe --tags)' -X 'main.gitCommit=$(git rev-parse HEAD)' -X 'main.buildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)' -X 'main.gitDescription=$(git log -1 --pretty=%B)'" -o ./bin .

# build and install binary
install: build
    cp ./bin/shout {{ INSTALL_DIR }}/shout

# unit testing with coverage
test:
    go test -cover ./...

# golangci-lint view only
lint:
    golangci-lint run --fast

# golangci-lint fix linting errors and format if possible
fix:
    golangci-lint run --fast --fix

release-snapshot:
    goreleaser release --snapshot --clean

release:
    goreleaser release --clean

# Locally serve documentation
serve-docs:
    mdbook serve
