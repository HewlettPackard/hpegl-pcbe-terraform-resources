// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

//go:build experimental

package system

import (
	"context"
	"errors"
	"fmt"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems/privatecloudbusiness"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DataSource{}

func NewDataSource() datasource.DataSource {
	return &DataSource{}
}

// DataSource defines the data source implementation.
type DataSource struct {
	client *client.PCBeClient
}

func (s *DataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_system"
}

func (s *DataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = SystemDataSourceSchema(ctx)
}

func (s *DataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		tflog.Warn(ctx, "provider has not been configured.")

		return
	}

	s.client = req.ProviderData.(*client.PCBeClient)
}

func createNameFilter(name string) string {
	return constants.NameFilter + name
}

func getSystemByName(
	ctx context.Context,
	client client.PCBeClient,
	name string,
) (privatecloudbusiness.V1beta1SystemsGetResponse_itemsable, error) {
	sysClient, _, err := client.NewSysClient(ctx)
	if err != nil {
		tflog.Error(ctx, "failed to create client")

		return nil, err
	}

	filter := createNameFilter(name)
	qp := privatecloudbusiness.V1beta1SystemsRequestBuilderGetQueryParameters{}
	qp.Filter = &filter
	grc := &privatecloudbusiness.V1beta1SystemsRequestBuilderGetRequestConfiguration{}
	grc.QueryParameters = &qp
	systems, err := sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().
		GetAsSystemsGetResponse(ctx, grc)

	if systems.GetTotal() == nil {
		msg := "total is nil"
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	total := *(systems.GetTotal())
	if total != 1 {
		msg := fmt.Sprintf(
			"required 1 system with name %s, got %d systems", name, total,
		)
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	return systems.GetItems()[0], err
}

func (s *DataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data SystemModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := *s.client
	system, err := getSystemByName(ctx, client, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"error reading system",
			"unexpected error: "+err.Error(),
		)

		return
	}

	name := system.GetName()
	if name == nil {
		resp.Diagnostics.AddError(
			"error reading system",
			"system name is nil",
		)

		return
	}

	id := system.GetId()
	if id == nil {
		resp.Diagnostics.AddError(
			"error reading system",
			"system id is nil",
		)

		return
	}

	data.Name = types.StringValue(*name)
	data.Id = types.StringValue(*id)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
