// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package acceptance

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/provider"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	providerConfig = `
terraform {
  required_providers {
    hpegl = {
      source = "github.com/HewlettPackard/hpegl-pcbe-terraform-resources"
    }
  }
}

provider "hpegl" {
        pc {
                host            = "http://localhost:8080"
                token           = "abcdefghijklmnopqrstuvwxyz-0123456789"

                http_dump       = true
                poll_interval   = 0.001
                max_polls       = 10
        }
}
`
)

var simulation = false

var testAccProtoV6ProviderFactories = map[string]func() (
	tfprotov6.ProviderServer, error,
){
	"scaffolding": providerserver.NewProtocol6WithError(
		provider.New("test")(),
	),
}

func checkUUIDAttr(resource string, attr string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("resource not found: %s", resource)
		}

		attrValue := rs.Primary.Attributes[attr]
		_, err := uuid.Parse(attrValue)

		return err
	}
}

func TestAccDatastoreResourceOk(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "VMFS"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"hpegl_pc_datastore.test",
			"name",
			"mclaren-ds19",
		),
		resource.TestCheckResourceAttr(
			"hpegl_pc_datastore.test",
			"datastore_type",
			"VMFS",
		),
		resource.TestCheckResourceAttr(
			"hpegl_pc_datastore.test",
			"capacity_in_bytes",
			"17179869184",
		),
		checkUUIDAttr("hpegl_pc_datastore.test", "id"),
	}

	if simulation {
		// In simulation mode the ID value is known in advance
		checks = append(checks,
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"id",
				"698de955-87b5-5fe6-b683-78c3948beede",
			),
		)
	}

	checkFn := resource.ComposeAggregateTestCheckFunc(checks...)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             config,
				Check:              checkFn,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: config,
				Check:  checkFn,
			},
			{
				Config:             config,
				Check:              checkFn,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccDatastoreResourceBadDatastoreType(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "XXX"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	expected := `Attribute datastore_type value must be one of: \["VMFS" "VVOL"\], got: "XXX"`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}

func TestAccDatastoreResourceMissingClusterInfo(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "XXX"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	}
	`

	expected := `The argument \"cluster_info\" is required, but no definition was found.`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}

func TestAccDatastoreResourceMissingClusterId(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "VVOL"
		capacity_in_bytes = 17179869184
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	expected := `The argument \"hci_cluster_uuid\" is required, but no definition was found.`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}

func TestAccDatastoreResourceMissingName(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		datastore_type = "VVOL"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	expected := `The argument \"name\" is required, but no definition was found.`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}

func TestAccDatastoreResourceMissingCapacity(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "VVOL"
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	expected := `The argument \"capacity_in_bytes\" is required, but no definition was found.`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}

func TestAccDatastoreResourceMissingClusterInfoName(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		capacity_in_bytes = 17179869184
		datastore_type = "VVOL"
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
		}
	}
	`

	expected := `Inappropriate value for attribute \"cluster_info\": attribute \"name\"`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(expected),
				PlanOnly:    true,
			},
		},
	})
}
