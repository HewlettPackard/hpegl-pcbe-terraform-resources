// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package datastore

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt/virtualization"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt/virtualization/v1beta1/datastores"
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

	if datastore.GetId() == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'id' is nil",
		)

		return
	}

	if *(datastore.GetId()) != datastoreID {
		(*diagsP).AddError(
			"error reading datastore",
			fmt.Sprintf("'id' mismatch: %s != %s",
				*(datastore.GetId()), datastoreID,
			),
		)

		return
	}

	if datastore.GetHciClusterUuid() == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'hciClusterUuid' is nil",
		)

		return
	}

	systemID := (*dataP).HciClusterUuid.ValueString()
	if *(datastore.GetHciClusterUuid()) != systemID {
		(*diagsP).AddError(
			"error reading datastore",
			fmt.Sprintf("'hciClusterUuid' mismatch: %s != %s",
				*(datastore.GetHciClusterUuid()), systemID,
			),
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

func createHypervisorClusterFilter(
	hciClusterUUID string,
	hypervisorClusterName string,
) string {
	return constants.HciClusterUUIDFilter + hciClusterUUID +
		constants.AndFilter +
		constants.NameFilter + hypervisorClusterName
}

func getTargetHypervisorClusterID(
	ctx context.Context,
	client client.PCBeClient,
	hciClusterUUID string, // "system" ID
	hypervisorClusterName string,
) (string, error) {
	virtClient, _, err := client.NewVirtClient(ctx)
	if err != nil {
		msg := "error getting hypervisor cluster ID for " +
			hypervisorClusterName + err.Error()
		tflog.Error(ctx, msg)

		return "", errors.New(msg)
	}
	qp := virtualization.V1beta1HypervisorClustersRequestBuilderGetQueryParameters{}
	filter := createHypervisorClusterFilter(hciClusterUUID, hypervisorClusterName)
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
			hypervisorClusterName + err.Error()
		tflog.Error(ctx, msg)

		return "", errors.New(msg)
	}

	if hypervisorClusters.GetTotal() == nil {
		msg := "error getting hypervisor cluster ID for " +
			hypervisorClusterName + " 'total' field is nil"
		tflog.Error(ctx, msg)

		return "", errors.New(msg)
	}

	total := *(hypervisorClusters.GetTotal())
	if total != 1 {
		msg := "error getting hypervisor cluster ID for " +
			hypervisorClusterName +
			fmt.Sprintf("required 1 hypervisorCluster with name %s, got %d",
				hypervisorClusterName, total)
		tflog.Error(ctx, msg)

		return "", errors.New(msg)
	}
	id := hypervisorClusters.GetItems()[0].GetId()

	return *id, nil
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *DatastoreModel,
	diagsP *diag.Diagnostics,
) {
	hciClusterUUID := (*dataP).HciClusterUuid.ValueString()
	clusterInfoName := (*dataP).ClusterInfo.Name.ValueString()

	targetHypervisorClusterID, err := getTargetHypervisorClusterID(
		ctx,
		client,
		hciClusterUUID,
		clusterInfoName,
	)
	if err != nil {
		(*diagsP).AddError(
			"error creating datastore",
			"could not get id of hypervisor cluster: "+
				clusterInfoName+" "+err.Error(),
		)

		return
	}

	virtClient, virtHeaderOpts, err := client.NewVirtClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error creating datastore",
			"unexpected error: "+err.Error(),
		)

		return
	}

	prb := virtualization.NewV1beta1DatastoresPostRequestBody()
	name := (*dataP).Name.ValueString()
	prb.SetName(&name)

	// TODO: (API) Allow setting cipher when FF-29512 is fixed
	// cipher := datastores.NONE_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER
	// enc := virtualization.NewV1beta1DatastoresPostRequestBody_volumeInfo_encryption()

	// TODO: should be able to use enum for cipher (bug in sdk or spec processing?)
	//cipherMap := map[string]any{
	//	"cipher": cipher.String(),
	//}

	// enc.SetAdditionalData(cipherMap)
	// qos := virtualization.NewV1beta1DatastoresPostRequestBody_volumeInfo_qos()
	// TODO: (API) Allow setting iopsLimit when FF-29512 is fixed
	// var iopsLimit float64 = -1 // -1 implies no limit
	// qos.SetIopsLimit(&iopsLimit)

	// TODO: (API) Allow setting mbsLimit when FF-29512 is fixed
	// var mbpsLimit float64 = -1 // -1 implies no limit
	// qos.SetMbpsLimit(&mbpsLimit)

	// volInfo := virtualization.NewV1beta1DatastoresPostRequestBody_volumeInfo()

	// TODO: (API) Allow setting duplication when FF-29512 is fixed
	// False := false
	// volInfo.SetDeduplication(&False)
	// volInfo.SetEncryption(enc)
	// volInfo.SetQos(qos)

	// prb.SetVolumeInfo(volInfo)

	sizeInBytes := (*dataP).CapacityInBytes.ValueInt64()
	prb.SetSizeInBytes(&sizeInBytes)

	prb.SetTargetHypervisorClusterId(&targetHypervisorClusterID)
	prb.SetStorageSystemId(&hciClusterUUID)

	datastoreType := datastores.VMFS_DATASTORESPOSTREQUESTBODY_DATASTORETYPE
	// TODO: should be able to use enum here (bug in sdk or spec processing?)
	datastoreTypeMap := map[string]any{
		"datastoreType": datastoreType.String(),
	}

	prb.SetAdditionalData(datastoreTypeMap)
	prc := virtualization.V1beta1DatastoresRequestBuilderPostRequestConfiguration{}
	_, err = virtClient.Virtualization().V1beta1().Datastores().Post(ctx, prb, &prc)
	if err != nil {
		(*diagsP).AddError(
			"error creating datastore",
			"unexpected error: "+err.Error(),
		)

		return
	}

	location := virtHeaderOpts.GetResponseHeaders().Get("Location")[0]
	virtHeaderOpts.ResponseHeaders.Clear()

	operationID := path.Base(location)
	asyncOperation := async.New(
		ctx,
		client,
		operationID,
		constants.TaskDatastore,
	)
	err = asyncOperation.Poll()
	if err != nil {
		(*diagsP).AddError(
			"error creating datastore",
			"unexpected poll error: "+err.Error(),
		)

		return
	}

	// TODO: Switch to  GetAssociatedResourceURI when FF-31311 is fixed
	uri, err := asyncOperation.GetSourceResourceURI()
	if err != nil {
		(*diagsP).AddError(
			"error creating datastore",
			"failed to get sourceResourceUri: "+err.Error(),
		)

		return
	}

	datastoreID := path.Base(uri)
	(*dataP).Id = types.StringValue(datastoreID)
}

func (r *Resource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data DatastoreModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	doCreate(ctx, *r.client, &data, &resp.Diagnostics)

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
	var data DatastoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	client := *r.client
	virtClient, virtHeaderOpts, err := client.NewVirtClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting datastore",
			"unexpected error: "+err.Error(),
		)

		return
	}

	rc := virtualization.
		V1beta1DatastoresDatastoresItemRequestBuilderDeleteRequestConfiguration{}

	_, err = virtClient.Virtualization().
		V1beta1().
		Datastores().
		ById(id).
		Delete(ctx, &rc)
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting datastore",
			"delete failed with: "+err.Error(),
		)

		return
	}

	location := virtHeaderOpts.GetResponseHeaders().Get("Location")[0]
	virtHeaderOpts.ResponseHeaders.Clear()
	operationID := path.Base(location)
	asyncOperation := async.New(ctx, client, operationID, constants.TaskDatastore)
	err = asyncOperation.Poll()
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting datastore",
			"delete failed with: "+err.Error(),
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
