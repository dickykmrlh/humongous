pkgs = $(shell go list ./...)

.PHONY: build

# go build command
build:
	@go build -v -o humongous cmd/*.go

# go run command
run:
	make build
	@./humongous

test:
	@echo "RUN TESTING..."
	@go test -v -cover -race $(pkgs)