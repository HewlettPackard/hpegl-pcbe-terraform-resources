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

func TestAccSecretDataSource(t *testing.T) {
	config := providerConfig + `
	data "hpegl_pc_secret" "test" {
		name = "mysecret1"
	}
	`

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(
			"data.hpegl_pc_secret.test",
			"name",
			"mysecret1",
		),
		checkUUIDAttr("data.hpegl_pc_secret.test", "id"),
	}

	if simulation {
		// In simulation mode the ID value is known in advance
		checks = append(checks,
			resource.TestCheckResourceAttr(
				"data.hpegl_pc_secret.test",
				"id",
				"cccfcad1-85b7-4162-b16e-f7cadc2c46b5",
			),
			resource.TestCheckResourceAttr(
				"data.hpegl_pc_secret.test",
				"name",
				"mysecret1",
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
		},
	})
}

func TestAccSecretDataSourceMissingName(t *testing.T) {
	config := providerConfig + `
	data "hpegl_pc_secret" "test" {
	}
	`
	expected := `The argument "name" is required, but no definition was found.`
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

func TestAccSystemDataSourceBadAuth(t *testing.T) {
	providerConfigBadAuth := `
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
                token           = "expired-token"

                http_dump       = true
                poll_interval   = 0.001
                max_polls       = 10
        }
}
`
	config := providerConfigBadAuth + `
	data "hpegl_pc_secret" "test" {
		name = "mysecret1"
	}
	`
	// TODO: return more informative error message - including
	// http response code (requires change to request handler)
	expected := `text does not support structured data`
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
