VERSION=$(shell git describe --tags --always)

# build
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# run all tests
.PHONY: test
test:
	@echo ""
	@echo "Running tests."
	go test ./... -count=1