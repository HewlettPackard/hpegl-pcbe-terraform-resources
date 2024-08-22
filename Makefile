#(C) Copyright 2024 Hewlett Packard Enterprise Development LP
#
# Note: this Makefile works with GNUMake and BSDMake
#
build:
	go build ./...

build-experimental:
	go build -tags experimental ./...

test:
	go test ./...

lint:
	@golangci-lint --version
	# Will use .golangci.yml
	golangci-lint run
