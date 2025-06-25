VERSION=$(shell git describe --tags --always)

# build
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -gcflags "all=-N -l" -o ./bin/ ./cmd/ksl/...

# run all tests
.PHONY: test
test:
	@echo ""
	@echo "Running tests."
	go test ./... -count=1

# regenerate ANTLR code
.PHONY: antlr
antlr:
	java -jar external/antlr-4.13.1-complete.jar  -o "$(realpath pkg/ksl/gen)" -visitor -lib "$(realpath pkg/ksl)" -Dlanguage=Go "$(realpath pkg/ksl/ksl.g4)"
