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

func testAccCheckResourceDestroyed(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Check if the resource is in the state
		if rs, ok := s.RootModule().Resources[resourceName]; ok {
			return fmt.Errorf("Resource %s still exists: %v", resourceName, rs.Primary.ID)
		}

		return nil
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
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
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
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"cluster_info.id",
				"298a299e-78f5-5acb-86ce-4e9fdc290ab7",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"hci_cluster_uuid",
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3",
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
			{
				// Destroy by specifying empty config
				Config: providerConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceDestroyed("hpegl_pc_datastore.test"),
				),
			},
			{
				// Import
				Check:        checkFn,
				Config:       config,
				ImportState:  true,
				ResourceName: "hpegl_pc_datastore.test",
				ImportStateId: "126fd201-9e6e-5e31-9ffb-a766265b1fd3," +
					"298a299e-78f5-5acb-86ce-4e9fdc290ab7," +
					"mclaren-ds19",
				ImportStatePersist: true,
			},
			{
				// Check post import state matches the resource config
				// e.g. verfies 'name' in state matches 'name' in config
				Config:             config,
				Check:              checkFn,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccDatastoreResourceMissing ensures that if the datastore
// resource goes missing just after POST, the state is purged
func TestAccDatastoreResourceMissing(t *testing.T) {
	resourceConfig := `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "VMFS"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
		}
	}
	`
	config := providerConfig + resourceConfig
	// token triggers a missing datastore (NotFound) during create
	badProviderConfig := `
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
			token           = "missing-datastore"

			http_dump       = true
			poll_interval   = 0.001
			max_polls       = 10
		}
	}
	`

	badConfig := badProviderConfig + resourceConfig

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
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"cluster_info.id",
				"298a299e-78f5-5acb-86ce-4e9fdc290ab7",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"hci_cluster_uuid",
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3",
			),
		)
	}

	checkFn := resource.ComposeAggregateTestCheckFunc(checks...)

	expected := `datastore missing: resource mclaren-ds19, not found`

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
				Config:      badConfig,
				Check:       checkFn,
				ExpectError: regexp.MustCompile(expected),
			},
		},
	})
}

func TestAccDatastoreDestroyMissing(t *testing.T) {
	resourceConfig := `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		datastore_type = "VMFS"
		capacity_in_bytes = 17179869184
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		cluster_info = {
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
		}
	}
	`
	config := providerConfig + resourceConfig
	// token triggers a missing datastore (NotFound) during Delete
	missingDatastoreProviderConfig := `
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
			token           = "missing-datastore"
			http_dump       = true
			poll_interval   = 0.001
			max_polls       = 10
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
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"cluster_info.id",
				"298a299e-78f5-5acb-86ce-4e9fdc290ab7",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_datastore.test",
				"hci_cluster_uuid",
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3",
			),
		)
	}

	checkFn := resource.ComposeAggregateTestCheckFunc(checks...)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  checkFn,
			},
			{
				// Case where non-existent datastore is destroyed
				// Empty resource config triggers destroy
				Config: missingDatastoreProviderConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceDestroyed("hpegl_pc_datastore.test"),
				),
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
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
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
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
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

func TestAccDatastoreImportBadId(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_datastore" "test" {
		name = "mclaren-ds19"
		hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
		datastore_type = "VVOL"
		capacity_in_bytes = 17179869184
		cluster_info = {
			"name": "5305-CL"
		}
	}
	`

	expected := `import has invalid datastore id format(.|\n)*698de955-87b5-5fe6-b683-78c3948beede`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Import
				ExpectError:  regexp.MustCompile(expected),
				Config:       config,
				ImportState:  true,
				ResourceName: "hpegl_pc_datastore.test",
				// Invalid id (not "<hci_cluster_uuid>,<cluster_info.id>,<datastore_name>")
				ImportStateId:      "698de955-87b5-5fe6-b683-78c3948beede",
				ImportStatePersist: true,
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
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
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
			"id": "298a299e-78f5-5acb-86ce-4e9fdc290ab7"
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

func TestAccDatastoreResourceMissingClusterInfoId(t *testing.T) {
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

	expected := `Inappropriate value for attribute \"cluster_info\": attribute \"id\"`
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
