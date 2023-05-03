# --- Required ----------------------------------------------------------------
export PATH   := $(PWD)/bin:$(PATH)                    # ./bin to $PATH
export SHELL  := bash                                  # Default Shell

GOPKGS := $(shell go list ./... | grep -vE "(cmd|examples|plugin|testdata|backup)" | tr -s '\n' ',' | sed 's/.\{1\}$$//' )


build:
	@ go build -trimpath -ldflags="-w -s" \
		-o bin/mirror ./cmd/mirror/

tests:
	go test -v -count=1 -race \
		-failfast \
		-parallel=2 \
		-timeout=1m \
		-covermode=atomic \
		-coverpkg=$(GOPKGS) -coverprofile=coverage.cov ./...

tests-summary:
	go test -v -count=1 -race \
		-failfast \
		-parallel=2 \
		-timeout=1m \
		-covermode=atomic \
		-coverpkg=$(GOPKGS) -coverprofile=coverage.cov --json ./... | tparse -all

generate:
	go run ./cmd/internal/generate-tests/ "$(PWD)/testdata"

lints:
	golangci-lint run --no-config ./... -D deadcode --skip-dirs "^(cmd|plugin|examples|testdata)"


cover:
	go tool cover -html=coverage.cov

install:
	go install -trimpath -v -ldflags="-w -s" \
		./cmd/mirror

