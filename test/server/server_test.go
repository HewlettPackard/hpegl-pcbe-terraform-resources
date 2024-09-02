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

func TestAccServerResource(t *testing.T) {
	config := providerConfig + `
	resource "hpegl_pc_server" "test" {
		name = "16.182.105.217"
		system_id = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	}
	`

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"hpegl_pc_server.test",
			"name",
			"16.182.105.217",
		),
		checkUUIDAttr("hpegl_pc_server.test", "id"),
		checkUUIDAttr("hpegl_pc_server.test", "system_id"),
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
