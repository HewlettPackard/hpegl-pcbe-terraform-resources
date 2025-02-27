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

func TestAccServerResource(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_server" "test" {
	        name   = "16.182.105.217"
	        system_id   = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	        esx_root_credential_id  = "cccfcad1-85b7-4162-b16e-f7cadc2c46b5"
	        ilo_admin_credential_id = "dddfcad1-85b7-4162-b16e-f7cadc2c46b5"
	        hypervisor_host = {
	                hypervisor_cluster_id = "acd4daea-e5e3-5f35-8be3-ce4a4b6d946c"
	                hypervisor_host_ip = "16.182.105.217"
	        }
	        ilo_network_info = {
	                ilo_ip = "16.182.105.216"
	                gateway = "16.182.104.1"
	                subnet_mask = "255.255.248.0"
	        }
	        server_network = [
	                {
	                        esx_ip_address = "10.0.0.88"
	                        data_ip_infos = [
	                                {
	                                        ip_address = "16.182.105.217"
	                                }
	                        ]
	                }
	        ]
	}
	`

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"hpegl_pc_server.test",
			"name",
			"16.182.105.217",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"name",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"hypervisor_host.name",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"hypervisor_host.hypervisor_cluster_name",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"hypervisor_host.hypervisor_host_ip",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"ilo_network_info.ilo_ip",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"ilo_network_info.gateway",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"ilo_network_info.subnet_mask",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"hypervisor_host.hypervisor_host_ip",
		),
		resource.TestCheckResourceAttrSet(
			"hpegl_pc_server.test",
			"server_network.0.esx_ip_address",
		),
		checkUUIDAttr("hpegl_pc_server.test", "id"),
		checkUUIDAttr("hpegl_pc_server.test", "system_id"),
		checkUUIDAttr("hpegl_pc_server.test", "esx_root_credential_id"),
		checkUUIDAttr("hpegl_pc_server.test", "ilo_admin_credential_id"),
		checkUUIDAttr("hpegl_pc_server.test", "hypervisor_host.hypervisor_cluster_id"),
		checkUUIDAttr("hpegl_pc_server.test", "hypervisor_host.id"),
	}

	if simulation {
		checks = append(checks,
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"id",
				"697e8cbf-df7e-570c-a3c7-912d4ce8375a",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"system_id",
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"esx_root_credential_id",
				"cccfcad1-85b7-4162-b16e-f7cadc2c46b5",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"ilo_admin_credential_id",
				"dddfcad1-85b7-4162-b16e-f7cadc2c46b5",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"name",
				"16.182.105.217",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"hypervisor_host.hypervisor_cluster_id",
				"acd4daea-e5e3-5f35-8be3-ce4a4b6d946c",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"hypervisor_host.hypervisor_cluster_name",
				"5305-CL",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"hypervisor_host.name",
				"16.182.105.217",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"hypervisor_host.id",
				"530b1894-9bd0-5627-9362-565aff1e5cbd",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"ilo_network_info.ilo_ip",
				"16.182.105.216",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"ilo_network_info.gateway",
				"16.182.104.1",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"ilo_network_info.subnet_mask",
				"255.255.248.0",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"hypervisor_host.hypervisor_host_ip",
				"16.182.105.217",
			),
			resource.TestCheckResourceAttr(
				"hpegl_pc_server.test",
				"server_network.0.esx_ip_address",
				"10.0.0.88",
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
				// Remove server from config to test delete
				Config: providerConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceDestroyed("hpegl_pc_server.test"),
				),
			},
			{
				// Import
				Check:        checkFn,
				Config:       config,
				ImportState:  true,
				ResourceName: "hpegl_pc_server.test",
				ImportStateId: "126fd201-9e6e-5e31-9ffb-a766265b1fd3," +
					"acd4daea-e5e3-5f35-8be3-ce4a4b6d946c," +
					"16.182.105.217",
				ImportStatePersist: true,
			},
			/*
				TODO: (API) This test cannot currently pass.
				See FF-31582 FF-31587 FF-31581 PC-6095 etc
				{
					// Check post import state matches the resource config
					// e.g. verfies 'name' in state matches 'name' in config
					Config:             config,
					Check:              checkFn,
					ExpectNonEmptyPlan: false,
				},
			*/
		},
	})
}

func TestAccServerImportBadId(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_server" "test" {
	        name   = "16.182.105.217"
	        system_id   = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	        esx_root_credential_id  = "cccfcad1-85b7-4162-b16e-f7cadc2c46b5"
	        ilo_admin_credential_id = "dddfcad1-85b7-4162-b16e-f7cadc2c46b5"
	        hypervisor_host = {
	                hypervisor_cluster_id = "acd4daea-e5e3-5f35-8be3-ce4a4b6d946c"
	                hypervisor_host_ip = "16.182.105.217"
	        }
	        ilo_network_info = {
	                ilo_ip = "16.182.105.216"
	                gateway = "16.182.104.1"
	                subnet_mask = "255.255.248.0"
	        }
	        server_network = [
	                {
	                        esx_ip_address = "10.0.0.88"
	                        data_ip_infos = [
	                                {
	                                        ip_address = "16.182.105.217"
	                                }
	                        ]
	                }
	        ]
	}
	`

	expected := `import has invalid server id format(.|\n)*698de955-87b5-5fe6-b683-78c3948beede`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Import
				ExpectError:  regexp.MustCompile(expected),
				Config:       config,
				ImportState:  true,
				ResourceName: "hpegl_pc_server.test",
				// Invalid id (not "<cluster id>,<datastore id>")
				ImportStateId:      "698de955-87b5-5fe6-b683-78c3948beede",
				ImportStatePersist: true,
			},
		},
	})
}
