#(C) Copyright 2024 Hewlett Packard Enterprise Development LP
#
# Note: this Makefile works with GNUMake and BSDMake
#
.PHONY: test docs

PROVIDER_TF = \
terraform { \\n \
 required_providers { \\n \
  hpegl = { \\n \
   source = \"github.com/HewlettPackard/hpegl-pcbe-terraform-resources\" \\n \
  } \\n \
 } \\n \
}

build:
	go build ./...

build-experimental:
	go build -tags experimental ./...

install:
	go install ./...

install-experimental:
	go install -tags experimental  ./...

docs:
	echo -e ${PROVIDER_TF} > __docs.tf; \
	tempfile=$$(mktemp); \
	terraform providers schema -json | \
	sed 's@github.com/hewlettpackard/hpegl-pcbe-terraform-resources@hpegl@g' > $$tempfile; \
	tfplugindocs generate --providers-schema $$tempfile -provider-name hpegl;
	rm __docs.tf

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
