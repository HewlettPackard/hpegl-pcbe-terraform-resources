// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package acceptance

import (
	"fmt"
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

func TestAccHypervisorclusterResource(t *testing.T) {
	config1 := providerConfig + `
	resource "hpegl_pc_hypervisor_cluster" "test" {
	  name = "mclaren01"
	  hci_cluster_uuid = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	  app_info = {
	    vmware = {
	      datacenter_info = {
	        name = "5305-DC"
	      }
	    }
	  }
	}
	`

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"hpegl_pc_hypervisor_cluster.test",
			"name",
			"mclaren01",
		),
		resource.TestCheckResourceAttr(
			"hpegl_pc_hypervisor_cluster.test",
			"app_info.vmware.datacenter_info.name",
			"5305-DC",
		),
		checkUUIDAttr("hpegl_pc_hypervisor_cluster.test", "id"),
	}

	if simulation {
		// In simulation mode the ID value is known in advance
		checks = append(checks,
			resource.TestCheckResourceAttr(
				"hpegl_pc_hypervisor_cluster.test",
				"id",
				"298a299e-78f5-5acb-86ce-4e9fdc290ab7",
			),
		)
	}

	checkFn := resource.ComposeAggregateTestCheckFunc(checks...)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check:  checkFn,
			},
			{
				// Remove hypervisor cluster from config to test delete
				Config: providerConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceDestroyed("hpegl_pc_hypervisor_cluster.test"),
				),
			},
		},
	})
}
