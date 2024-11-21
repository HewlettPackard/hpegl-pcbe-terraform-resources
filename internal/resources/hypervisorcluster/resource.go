// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package hypervisorcluster

import (
	"context"
	"fmt"
	"path"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems/privatecloudbusiness"
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
	resp.TypeName = req.ProviderTypeName + "_hypervisor_cluster"
}

func (r *Resource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = HypervisorclusterResourceSchema(ctx)
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
	dataP *HypervisorclusterModel,
	diagsP *diag.Diagnostics,
) {
	hypervisorClusterID := (*dataP).Id.ValueString()

	grc := virtualization.
		V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration{}
	virtClient, _, err := client.NewVirtClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"unexpected error creating client: "+err.Error(),
		)

		return
	}

	getResp, err := virtClient.Virtualization().
		V1beta1().
		HypervisorClusters().
		ByClusterId(hypervisorClusterID).
		GetAsClusterGetResponse(ctx, &grc)
	if err != nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"unexpected error: "+err.Error(),
		)

		return
	}

	if getResp.GetId() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'id' is nil",
		)

		return
	}

	if *(getResp.GetId()) != hypervisorClusterID {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			fmt.Sprintf("'id' mismatch: %s != %s",
				*(getResp.GetId()), hypervisorClusterID),
		)

		return
	}

	if getResp.GetHciClusterUuid() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'hciClusterUuid' is nil",
		)

		return
	}

	systemID := (*dataP).HciClusterUuid.ValueString()
	if *(getResp.GetHciClusterUuid()) != systemID {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			fmt.Sprintf("'hciClusterUuid' mismatch: %s != %s",
				*(getResp.GetHciClusterUuid()), systemID),
		)

		return
	}

	(*dataP).HciClusterUuid = types.StringValue(*(getResp.GetHciClusterUuid()))

	if getResp.GetName() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'name' is nil",
		)

		return
	}

	(*dataP).Name = types.StringValue(*(getResp.GetName()))

	if getResp.GetAppInfo() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'appInfo' is nil",
		)

		return
	}

	if getResp.GetAppInfo().GetVmware() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'vmware' is nil",
		)

		return
	}

	if getResp.GetAppInfo().GetVmware().GetDatacenterInfo() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'datacenterInfo' is nil",
		)

		return
	}

	if getResp.GetAppInfo().GetVmware().GetDatacenterInfo().GetName() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+hypervisorClusterID,
			"'datacenterInfo.name' is nil",
		)

		return
	}

	dsName := getResp.GetAppInfo().GetVmware().GetDatacenterInfo().GetName()

	m := map[string]attr.Value{
		"name": types.StringValue(*dsName),
	}
	attrTypes := DatacenterInfoValue{}.AttributeTypes(ctx)
	datacenterInfoValue, diags := NewDatacenterInfoValue(attrTypes, m)
	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	dataCenterInfoObj, diags := datacenterInfoValue.ToObjectValue(ctx)
	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	m = map[string]attr.Value{
		"datacenter_info": dataCenterInfoObj,
	}

	vmwareValue, diags := NewVmwareValue(VmwareValue{}.AttributeTypes(ctx), m)
	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	vmwareValueObj, diags := vmwareValue.ToObjectValue(ctx)
	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	(*dataP).AppInfo.Vmware = vmwareValueObj
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *HypervisorclusterModel,
	diagsP *diag.Diagnostics,
) {
	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error creating hypervisorcluster",
			"unexpected error: "+err.Error(),
		)

		return
	}

	prb := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorClusterPostRequestBody()
	name := (*dataP).Name.ValueString()
	prb.SetHypervisorClusterName(&name)

	// TODO: (API) Support configureVds when FF-28975 is addressed
	configureVds := false
	prb.SetConfigureVds(&configureVds)

	vmware, diags := NewVmwareValue(
		(*dataP).AppInfo.Vmware.AttributeTypes(ctx),
		(*dataP).AppInfo.Vmware.Attributes(),
	)
	(*diagsP).Append(diags...)
	if diags.HasError() {
		return
	}

	datacenterInfo, diags := NewDatacenterInfoValue(
		vmware.DatacenterInfo.AttributeTypes(ctx),
		vmware.DatacenterInfo.Attributes(),
	)

	(*diagsP).Append(diags...)
	if diags.HasError() {
		return
	}

	dataCenterName := datacenterInfo.Name.ValueString()
	prb.SetVsphereDatacenterName(&dataCenterName)
	prc := privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration{}
	systemID := (*dataP).HciClusterUuid.ValueString()
	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().
		ById(systemID).
		AddHypervisorCluster().
		Post(ctx, prb, &prc)
	if err != nil {
		(*diagsP).AddError(
			"error creating hypervisorcluster",
			"unexpected error: "+err.Error(),
		)

		return
	}

	location := sysHeaderOpts.GetResponseHeaders().Get("Location")[0]
	sysHeaderOpts.ResponseHeaders.Clear()
	operationID := path.Base(location)
	asyncOperation := async.New(
		ctx,
		client,
		operationID,
		constants.TaskHypervisorCluster,
	)
	err = asyncOperation.Poll()
	if err != nil {
		(*diagsP).AddError(
			"error creating hypervisorcluster",
			"unexpected poll error: "+err.Error(),
		)

		return
	}

	uri, err := asyncOperation.GetAssociatedResourceURI()
	if err != nil {
		(*diagsP).AddError(
			"error creating hypervisorcluster",
			"unexpected associated resource error: "+err.Error(),
		)

		return
	}

	hypervisorClusterID := path.Base(uri)
	(*dataP).Id = types.StringValue(hypervisorClusterID)
}

func (r *Resource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data HypervisorclusterModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	doCreate(ctx, *r.client, &data, &resp.Diagnostics)
	// Write state to capture the id
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	doRead(ctx, *r.client, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data HypervisorclusterModel

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
	tflog.Error(ctx, "update hypervisorcluster is not implemented")
}

func (r *Resource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data HypervisorclusterModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	client := *r.client
	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisorcluster "+id,
			"unexpected error creating client: "+err.Error(),
		)

		return
	}

	prb := privatecloudbusiness.
		NewV1beta1SystemsItemRemoveHypervisorClustersPostRequestBody()
	prb.SetHypervisorClusterIds([]string{id})
	prc := privatecloudbusiness.
		V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration{}

	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().
		ById(data.HciClusterUuid.ValueString()).
		RemoveHypervisorClusters().Post(ctx, prb, &prc)

	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisorcluster "+id,
			"unexpected error: "+err.Error(),
		)

		return
	}

	location := sysHeaderOpts.GetResponseHeaders().Get("Location")[0]
	sysHeaderOpts.ResponseHeaders.Clear()
	operationID := path.Base(location)

	// If resp.Diagnostics is not empty, then Delete is
	// considered to have failed; the hypervisor cluster entry
	// in the tfstate file will not be removed
	asyncOperation := async.New(
		ctx,
		client,
		operationID,
		constants.TaskHypervisorCluster,
	)
	err = asyncOperation.Poll()
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisorcluster "+id,
			"unexpected error: "+err.Error(),
		)

		return
	}
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, tfpath.Root("id"), req, resp)
}
