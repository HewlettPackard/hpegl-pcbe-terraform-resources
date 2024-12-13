// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

//go:build experimental

package provider

import (
	"context"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/datasources/secret"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/datasources/system"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// DataSources defines the data sources implemented in the provider.
func (p *PCBeProvider) DataSources(
	_ context.Context,
) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		system.NewDataSource,
		secret.NewDataSource,
	}
}
