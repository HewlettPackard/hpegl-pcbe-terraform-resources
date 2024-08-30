#(C) Copyright 2024 Hewlett Packard Enterprise Development LP
#
# Note: this Makefile works with GNUMake and BSDMake
#

build:
	go build ./...

build-experimental:
	go build -tags experimental ./...

install:
	go install ./...

install-experimental:
	go install -tags experimental  ./...

test:
	go test ./...

testacc:
	go install -tags experimental ./cmd/...
	env TF_ACC=1 go test -v ./test/...

testacc-simulation:
	go install -tags experimental,simulation ./cmd/...
	tempfile=$$(mktemp); \
	sed "s@__HOME__@${HOME}@g" test/.terraformrc > $$tempfile; \
	env TF_ACC=1 env TF_CLI_CONFIG_FILE=$$tempfile \
		go test -v -tags simulation ./test/...

lint:
	@golangci-lint --version
	# Will use .golangci.yml
	golangci-lint run
