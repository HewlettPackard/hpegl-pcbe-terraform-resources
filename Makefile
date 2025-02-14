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
	printf "${PROVIDER_TF}" > __docs.tf; \
	tfconfig=$$(mktemp); \
	sed "s@__HOME__@${HOME}@g" test/.terraformrc > $$tfconfig; \
	schemafile=$$(mktemp); \
	env TF_CLI_CONFIG_FILE=$$tfconfig terraform providers schema -json | \
	sed 's@github.com/hewlettpackard/hpegl-pcbe-terraform-resources@hpegl@g' > $$schemafile; \
	rm __docs.tf; \
	tfplugindocs generate --provider-name hpegl --providers-schema $$schemafile

test:
	go test ./...

testacc:
	go install -tags experimental ./cmd/...
	env TF_ACC=1 go test -v ./test/...

testacc-simulation:
	go install -tags experimental,simulation ./cmd/...
	tfconfig=$$(mktemp); \
	sed "s@__HOME__@${HOME}@g" test/.terraformrc > $$tfconfig; \
	env TF_ACC=1 env TF_CLI_CONFIG_FILE=$$tfconfig \
	env TF_LOG_PATH_MASK=$${tfconfig}.%s.log \
		go test -v -tags simulation ./test/...

demo:
	go install -tags experimental,simulation ./cmd/...
	tfconfig=$$(mktemp); \
	sed "s@__HOME__@${HOME}@g" test/.terraformrc > $$tfconfig; \
	env TF_LOG=INFO env TF_CLI_CONFIG_FILE=$$tfconfig \
		terraform -chdir=examples apply -auto-approve; \
	env TF_LOG=INFO env TF_CLI_CONFIG_FILE=$$tfconfig \
		terraform -chdir=examples destroy -auto-approve; \
	env TF_LOG=INFO env TF_CLI_CONFIG_FILE=$$tfconfig \
		terraform -chdir=examples import hpegl_pc_datastore.my_datastore \
		126fd201-9e6e-5e31-9ffb-a766265b1fd3,298a299e-78f5-5acb-86ce-4e9fdc290ab7,mclaren-ds19; \
	env TF_LOG=INFO env TF_CLI_CONFIG_FILE=$$tfconfig \
		terraform -chdir=examples import hpegl_pc_hypervisor_cluster.my_hypervisor_cluster \
		298a299e-78f5-5acb-86ce-4e9fdc290ab7

lint:
	@golangci-lint --version
	# Will use .golangci.yml
	golangci-lint run
