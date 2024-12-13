// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

//go:build experimental

package secret

import (
	"context"
	"errors"
	"fmt"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices/dataservices"
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
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (s *DataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = SecretDataSourceSchema(ctx)
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

func getSecretByName(
	ctx context.Context,
	client client.PCBeClient,
	name string,
) (dataservices.V1beta1SecretsGetResponse_itemsable, error) {
	dataServicesClient, _, err := client.NewDataServicesClient(ctx)
	if err != nil {
		tflog.Error(ctx, "failed to create client")

		return nil, err
	}

	filter := createNameFilter(name)
	qp := dataservices.V1beta1SecretsRequestBuilderGetQueryParameters{}
	qp.Filter = &filter
	grc := &dataservices.V1beta1SecretsRequestBuilderGetRequestConfiguration{}
	grc.QueryParameters = &qp
	secrets, err := dataServicesClient.
		DataServices().
		V1beta1().
		Secrets().
		GetAsSecretsGetResponse(ctx, grc)

	if err != nil {
		tflog.Error(ctx, "failed to get secret by name: "+name)

		return nil, err
	}

	if secrets.GetTotal() == nil {
		msg := "total is nil"
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	total := *(secrets.GetTotal())
	if total != 1 {
		msg := fmt.Sprintf(
			"required 1 secret with name %s, got %d secrets", name, total,
		)
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	return secrets.GetItems()[0], err
}

func (s *DataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data SecretModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := *s.client
	secret, err := getSecretByName(ctx, client, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"error reading secret",
			"unexpected error: "+err.Error(),
		)

		return
	}
	name := secret.GetName()
	if name == nil {
		resp.Diagnostics.AddError(
			"error reading secret",
			"secret name is nil",
		)

		return
	}

	id := secret.GetId()
	if id == nil {
		resp.Diagnostics.AddError(
			"error reading system",
			"secret id is nil",
		)

		return
	}
	data.Id = types.StringValue(id.String())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
