// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package provider

import (
	"context"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/defaults"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &PCBeProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PCBeProvider{
			version: version,
		}
	}
}

type PCBeCfg struct {
	Host         types.String  `tfsdk:"host"`
	Token        types.String  `tfsdk:"token"`
	HTTPDump     types.Bool    `tfsdk:"http_dump"`
	MaxPolls     types.Int32   `tfsdk:"max_polls"`
	PollInterval types.Float32 `tfsdk:"poll_interval"`
}

type PCBeProviderModel struct {
	PCBeCfg PCBeCfg `tfsdk:"pc"`
}

type PCBeProvider struct {
	version string
}

func (p *PCBeProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = constants.ProviderType + "_" + constants.ProviderBlock
	resp.Version = p.version
}

func (p *PCBeProvider) Schema(
	_ context.Context,
	_ provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Blocks: map[string]schema.Block{
			"pc": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Required: true,
					},
					"token": schema.StringAttribute{
						Required: true,
					},
					"http_dump": schema.BoolAttribute{
						Optional: true,
					},
					"max_polls": schema.Int32Attribute{
						Optional: true,
					},
					"poll_interval": schema.Float32Attribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (p *PCBeProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var config PCBeProviderModel
	var httpDump bool
	var host string
	var token string
	var maxPolls int32
	var pollInterval float32

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.PCBeCfg.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"unknown 'host' value in provider configuration block",
			"the provider cannot create an API client "+
				"as there is an unknown configuration value for "+
				"the API host.",
		)
	}

	if config.PCBeCfg.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"unknown 'token' value in provider configuration block",
			"the provider cannot create an API client "+
				"as there is an unknown configuration value for "+
				"the API token",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if !config.PCBeCfg.Host.IsNull() {
		host = config.PCBeCfg.Host.ValueString()
	}

	if !config.PCBeCfg.Token.IsNull() {
		token = config.PCBeCfg.Token.ValueString()
	}

	if !config.PCBeCfg.HTTPDump.IsNull() {
		httpDump = config.PCBeCfg.HTTPDump.ValueBool()
	}

	if !config.PCBeCfg.MaxPolls.IsNull() {
		maxPolls = config.PCBeCfg.MaxPolls.ValueInt32()
	}

	if !config.PCBeCfg.PollInterval.IsNull() {
		pollInterval = config.PCBeCfg.PollInterval.ValueFloat32()
	}

	if resp.Diagnostics.HasError() {
		return
	}

	cfg := client.Config{
		Token:        token,
		Host:         host,
		HTTPDump:     httpDump,
		MaxPolls:     maxPolls,
		PollInterval: pollInterval,
		HTTPTimeout:  defaults.HTTPTimeout,
	}
	client, err := client.NewPCBeClient(context.Background(), cfg)
	if err != nil {
		resp.Diagnostics.AddError(
			"unable to create API client",
			err.Error(),
		)

		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}
