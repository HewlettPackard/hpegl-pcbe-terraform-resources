// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package hypervisorcluster

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/errordefs"
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

func createHypervisorClusterFilter(
	hciClusterUUID string,
	hypervisorClusterName string,
) string {
	return constants.HciClusterUUIDFilter + hciClusterUUID +
		constants.AndFilter +
		constants.NameFilter + hypervisorClusterName
}

func getHypervisorCluster(
	ctx context.Context,
	client client.PCBeClient,
	hciClusterUUID string, // "system" ID
	name string,
) (virtualization.V1beta1HypervisorClustersGetResponse_itemsable, error) {
	virtClient, _, err := client.NewVirtClient(ctx)
	if err != nil {
		msg := "error getting hypervisor cluster ID for " +
			name + err.Error()
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}
	qp := virtualization.V1beta1HypervisorClustersRequestBuilderGetQueryParameters{}
	filter := createHypervisorClusterFilter(hciClusterUUID, name)
	qp.Filter = &filter
	grc := virtualization.
		V1beta1HypervisorClustersRequestBuilderGetRequestConfiguration{}
	grc.QueryParameters = &qp

	hypervisorClusters, err := virtClient.Virtualization().
		V1beta1().
		HypervisorClusters().
		GetAsHypervisorClustersGetResponse(ctx, &grc)
	if err != nil {
		msg := "error getting hypervisor cluster ID for " +
			name + err.Error()
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	if hypervisorClusters.GetTotal() == nil {
		msg := "error getting hypervisor cluster ID for " +
			name + " 'total' field is nil"
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	total := *(hypervisorClusters.GetTotal())
	if total == 0 {
		// We don't have a hypervisor cluster with the given name
		return nil, errordefs.NewNotFoundError(name)
	}

	if total != 1 {
		msg := "error getting hypervisor cluster ID for " +
			name +
			fmt.Sprintf("required 1 hypervisorCluster with name %s, got %d",
				name, total)
		tflog.Error(ctx, msg)

		return nil, errors.New(msg)
	}

	return hypervisorClusters.GetItems()[0], nil
}

func doRead(
	ctx context.Context,
	cluster virtualization.V1beta1HypervisorClustersGetResponse_itemsable,
	dataP *HypervisorclusterModel,
	diagsP *diag.Diagnostics,
) {
	if cluster.GetId() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster ",
			"'id' is nil",
		)

		return
	}

	(*dataP).Id = types.StringValue(*(cluster.GetId()))

	if cluster.GetHciClusterUuid() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'hciClusterUuid' is nil",
		)

		return
	}

	(*dataP).HciClusterUuid = types.StringValue(*(cluster.GetHciClusterUuid()))

	if cluster.GetName() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'name' is nil",
		)

		return
	}

	(*dataP).Name = types.StringValue(*(cluster.GetName()))

	if cluster.GetAppInfo() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'appInfo' is nil",
		)

		return
	}

	if cluster.GetAppInfo().GetVmware() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'vmware' is nil",
		)

		return
	}

	if cluster.GetAppInfo().GetVmware().GetDatacenterInfo() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'datacenterInfo' is nil",
		)

		return
	}

	if cluster.GetAppInfo().GetVmware().GetDatacenterInfo().GetName() == nil {
		(*diagsP).AddError(
			"error reading hypervisor cluster "+
				(*dataP).Id.ValueString(),
			"'datacenterInfo.name' is nil",
		)

		return
	}

	dsName := cluster.GetAppInfo().GetVmware().GetDatacenterInfo().GetName()

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

	m = map[string]attr.Value{
		"vmware": vmwareValueObj,
	}

	appInfoValue, diags := NewAppInfoValue(AppInfoValue{}.AttributeTypes(ctx), m)
	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}
	(*dataP).AppInfo = appInfoValue
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *HypervisorclusterModel,
) error {
	name := (*dataP).Name.ValueString()
	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)
	if err != nil {
		tflog.Error(ctx, "failed to create client "+err.Error())

		return errordefs.NewClientError(name)
	}

	prb := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorClusterPostRequestBody()
	prb.SetHypervisorClusterName(&name)

	// TODO: (API) Support configureVds when FF-28975 is addressed
	configureVds := false
	prb.SetConfigureVds(&configureVds)

	vmware, diags := NewVmwareValue(
		(*dataP).AppInfo.Vmware.AttributeTypes(ctx),
		(*dataP).AppInfo.Vmware.Attributes(),
	)
	if diags.HasError() {
		return errordefs.NewValueError("vmware")
	}

	datacenterInfo, diags := NewDatacenterInfoValue(
		vmware.DatacenterInfo.AttributeTypes(ctx),
		vmware.DatacenterInfo.Attributes(),
	)

	if diags.HasError() {
		return errordefs.NewValueError("datacenterinfo")
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
		tflog.Error(ctx, "failed to create resource "+err.Error())

		return errordefs.NewCreateError(name)
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
		tflog.Error(ctx, "failed to poll resource "+err.Error())

		return errordefs.NewPollError(name)
	}

	uri, err := asyncOperation.GetAssociatedResourceURI()
	if err != nil {
		tflog.Error(ctx, "failed to get associatedResourceUri: "+
			err.Error(),
		)

		return errordefs.NewNoURIError(name)
	}

	hypervisorClusterID := path.Base(uri)
	(*dataP).Id = types.StringValue(hypervisorClusterID)

	return nil
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

	err := doCreate(ctx, *r.client, &data)
	if err != nil {
		resp.Diagnostics.AddError(
			"create hypervisor cluster error",
			err.Error(),
		)
		if errors.As(err, &errordefs.Create) ||
			errors.As(err, &errordefs.Client) ||
			errors.As(err, &errordefs.Value) {
			return
		}

		// tainted state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

		return
	}
	// Write state to capture the id
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	hypervisorCluster, err := getHypervisorCluster(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.Name.ValueString())
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.Diagnostics.AddError(
				"error creating hypervisor cluster "+
					data.Name.ValueString(),
				"hypervisor cluster missing: "+err.Error(),
			)
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error creating hypervisor cluster"+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, hypervisorCluster, &data, &resp.Diagnostics)

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

	hypervisorCluster, err := getHypervisorCluster(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error reading hypervisor cluster "+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, hypervisorCluster, &data, &resp.Diagnostics)

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

	hypervisorCluster, err := getHypervisorCluster(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.Name.ValueString())
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			tflog.Debug(
				ctx, "hypervisor cluster not found during delete",
			)

			return
		}
		resp.Diagnostics.AddError(
			"error deleting hypervisor cluster"+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	if hypervisorCluster.GetId() == nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisor cluster",
			"hypervisor cluster has no id",
		)

		return
	}

	id := *(hypervisorCluster.GetId())
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
			"delete failed with: "+err.Error(),
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
			"delete failed with: "+err.Error(),
		)

		return
	}
}

// Import only grants access to a single "ID" parameter. Therefore, we have to
// combine the "hci_cluster_uuid" and hypervisor cluster "name" values into the
// single req.ID string
func parseImportID(
	id string,
) (systemID string, clusterName string, error error) {
	params := strings.Split(id, ",")
	if len(params) != 2 || params[0] == "" || params[1] == "" {
		return "", "", errors.New("invalid import ID format")
	}

	return params[0], params[1], nil
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	systemID, name, err := parseImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"import has invalid hypervisor cluster id format",
			"Provided import ID \""+req.ID+"\" is invalid. "+
				"Format must be \"<hci_cluster_uuid>,<hypervisor_cluster_name>\". For example: "+
				"f8d3e2fd-a0e0-41a3-83b3-a8f92b21a9f3,cluster1",
		)

		return
	}

	diags := resp.State.SetAttribute(ctx, tfpath.Root("name"), name)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, tfpath.Root("hci_cluster_uuid"), systemID)
	resp.Diagnostics.Append(diags...)
}
