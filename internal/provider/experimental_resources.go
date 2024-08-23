// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

//go:build experimental

package provider

import (
	"context"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/resources/datastore"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/resources/hypervisorcluster"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Resources defines the resources implemented in the provider.
func (p *PCBeProvider) Resources(
	_ context.Context,
) []func() resource.Resource {
	return []func() resource.Resource{
		datastore.NewResource,
		hypervisorcluster.NewResource,
	}
}
