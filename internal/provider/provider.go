// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package provider

import (
	"context"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &pcbeProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &pcbeProvider{
			version: version,
		}
	}
}

type pcbeProvider struct {
	version string
}

func (p *pcbeProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = constants.ProviderType + "_" + constants.ProviderBlock
	resp.Version = p.version
}

func (p *pcbeProvider) Schema(
	_ context.Context,
	_ provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
}

func (p *pcbeProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
}

// DataSources defines the data sources implemented in the provider.
func (p *pcbeProvider) DataSources(
	_ context.Context,
) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *pcbeProvider) Resources(
	_ context.Context,
) []func() resource.Resource {
	return []func() resource.Resource{}
}
