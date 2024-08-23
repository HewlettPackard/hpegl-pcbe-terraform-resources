// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package datastore

import (
	"context"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt/virtualization"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	tfpath "github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &Resource{}
	_ resource.ResourceWithImportState = &Resource{}
)

func NewResource() resource.Resource {
	return &Resource{}
}

// Resource defines the resource implementation.
type Resource struct {
	client *client.PCBeClient
}

func (r *Resource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_datastore"
}

func (r *Resource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = DatastoreResourceSchema(ctx)
}

func (r *Resource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.PCBeClient)
}

func doRead(
	ctx context.Context,
	client client.PCBeClient,
	dataP *DatastoreModel,
	diagsP *diag.Diagnostics,
) {
	virtClient, _, err := client.NewVirtClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error reading datastore",
			"unexpected error: "+err.Error(),
		)

		return
	}

	datastoreID := (*dataP).Id.ValueString()

	grc := virtualization.
		V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration{}

	datastore, err := virtClient.Virtualization().
		V1beta1().
		Datastores().
		ById(datastoreID).
		GetAsDatastoresGetResponse(ctx, &grc)
	if err != nil {
		(*diagsP).AddError(
			"error reading datastore",
			"get datastore by ID returned: "+err.Error(),
		)

		return
	}

	datastoreName := datastore.GetName()
	if datastoreName == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'name' is nil",
		)

		return
	}

	(*dataP).Name = types.StringValue(*datastoreName)

	clusterInfo := datastore.GetClusterInfo()
	if clusterInfo == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'clusterInfo' is nil",
		)

		return
	}

	clusterInfoName := clusterInfo.GetName()
	if clusterInfoName == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'clusterInfo.name' is nil",
		)

		return
	}

	clusterInfoID := clusterInfo.GetId()
	if clusterInfoID == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'clusterInfo.id' is nil",
		)

		return
	}

	value := map[string]attr.Value{
		"id":   types.StringValue(*clusterInfoID),
		"name": types.StringValue(*clusterInfoName),
	}

	ci, diags := NewClusterInfoValue(
		(*dataP).ClusterInfo.AttributeTypes(ctx), value,
	)

	(*diagsP).Append((diags)...)
	if (*diagsP).HasError() {
		return
	}

	(*dataP).ClusterInfo = ci

	capacityInBytes := datastore.GetCapacityInBytes()
	if capacityInBytes == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'capacityInBytes' is nil",
		)

		return
	}

	(*dataP).CapacityInBytes = types.Int64Value(*capacityInBytes)

	datacentersInfo := datastore.GetDatacentersInfo()
	if datacentersInfo == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'datacentersInfo' is nil",
		)

		return
	}

	numDatacenters := len(datacentersInfo)
	if numDatacenters == 0 {
		(*diagsP).AddError(
			"error reading datastore",
			"'datacentersInfo' has zero length",
		)

		return
	}

	datacentersInfoValues := []attr.Value{}

	for _, d := range datacentersInfo {
		if d == nil {
			(*diagsP).AddError(
				"error reading datastore",
				"'datacentersInfo' entry is nil",
			)

			return
		}

		if d.GetName() == nil {
			(*diagsP).AddError(
				"error reading datastore",
				"'datacentersInfo' entry 'name' is nil",
			)

			return
		}

		if d.GetId() == nil {
			(*diagsP).AddError(
				"error reading datastore",
				"'datacentersInfo' entry 'id' is nil",
			)

			return
		}

		m := map[string]attr.Value{
			"name": types.StringValue(*(d.GetName())),
			"id":   types.StringValue(*(d.GetId())),
		}

		v, diags := NewDatacentersInfoValue(
			DatacentersInfoValue{}.AttributeTypes(ctx),
			m,
		)
		(*diagsP).Append((diags)...)
		if (*diagsP).HasError() {
			return
		}

		datacentersInfoValues = append(datacentersInfoValues, v)
	}

	(*dataP).DatacentersInfo = types.ListValueMust((*dataP).
		DatacentersInfo.ElementType(ctx), datacentersInfoValues)
}

func (r *Resource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	tflog.Error(ctx, "update datastore is not implemented")
}

func (r *Resource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data DatastoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	doRead(ctx, *r.client, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	tflog.Error(ctx, "update datastore is not implemented")
}

func (r *Resource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	tflog.Error(ctx, "delete datastore is not implemented")
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, tfpath.Root("id"), req, resp)
}
