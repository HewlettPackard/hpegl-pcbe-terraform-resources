// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package datastore

import (
	"context"
	"errors"
	"path"
	"strings"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/errordefs"
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
	datastore virtualization.V1beta1DatastoresGetResponse_itemsable,
	dataP *DatastoreModel,
	diagsP *diag.Diagnostics,
) {
	if datastore.GetId() == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'id' is nil",
		)

		return
	}

	(*dataP).Id = types.StringValue(*(datastore.GetId()))

	if datastore.GetHciClusterUuid() == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'hciClusterUuid' is nil",
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

	datastoreType := datastore.GetAdditionalData()["datastoreType"]
	if datastoreType == nil {
		(*diagsP).AddError(
			"error reading datastore",
			"'datastoreType' is nil",
		)

		return
	}

	datastoreTypeString, ok := datastoreType.(*string)
	if !ok {
		(*diagsP).AddError(
			"error reading datastore",
			"'datastoreType' is invalid",
		)

		return
	}
	(*dataP).DatastoreType = types.StringValue(*datastoreTypeString)

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

func createDatastoreFilter(
	hciClusterUUID string,
	hypervisorClusterID string,
	DatastoreName string,
) string {
	return constants.HciClusterUUIDFilter + hciClusterUUID +
		constants.AndFilter +
		constants.NameFilter + DatastoreName +
		constants.AndFilter +
		constants.HypervisorClusterIDFilter + hypervisorClusterID
}

func getDatastore(
	ctx context.Context,
	client client.PCBeClient,
	hciClusterUUID string, // "system" ID
	hypervisorClusterID string, // "cluster" ID
	name string, // datastore name
) (virtualization.V1beta1DatastoresGetResponse_itemsable, error) {
	virtClient, _, err := client.NewVirtClient(ctx)
	if err != nil {
		return nil, err
	}

	filter := createDatastoreFilter(
		hciClusterUUID,
		hypervisorClusterID,
		name,
	)

	qp := virtualization.V1beta1DatastoresRequestBuilderGetQueryParameters{}
	qp.Filter = &filter

	grc := virtualization.
		V1beta1DatastoresRequestBuilderGetRequestConfiguration{}
	grc.QueryParameters = &qp

	datastores, err := virtClient.Virtualization().
		V1beta1().
		Datastores().
		GetAsDatastoresGetResponse(ctx, &grc)
	if err != nil {
		return nil, err
	}

	if datastores.GetTotal() == nil {
		return nil, errors.New("datastores 'total' field is nil")
	}

	total := *(datastores.GetTotal())
	if total == 0 {
		// datastore does not exist
		return nil, errordefs.NewNotFoundError(name)
	}

	if total != 1 {
		msg := "error getting datastore ID for " + name

		return nil, errors.New(msg)
	}

	return datastores.GetItems()[0], nil
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *DatastoreModel,
) error {
	hciClusterUUID := (*dataP).HciClusterUuid.ValueString()
	targetHypervisorClusterID := (*dataP).ClusterInfo.Id.ValueString()
	name := (*dataP).Name.ValueString()

	virtClient, virtHeaderOpts, err := client.NewVirtClient(ctx)
	if err != nil {
		tflog.Error(ctx, "failed to create client "+err.Error())

		return errordefs.NewClientError(name)
	}

	prb := virtualization.NewV1beta1DatastoresPostRequestBody()
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
		tflog.Error(ctx, "failed to create resource "+err.Error())

		return errordefs.NewCreateError(name)
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
		tflog.Error(ctx, "failed to poll resource "+err.Error())

		return errordefs.NewPollError(name)
	}

	// TODO: Switch to  GetAssociatedResourceURI when FF-31311 is fixed
	uri, err := asyncOperation.GetSourceResourceURI()
	if err != nil {
		tflog.Error(ctx, "failed to get sourceResourceUri: "+err.Error())

		return errordefs.NewNoURIError(name)
	}

	datastoreID := path.Base(uri)
	(*dataP).Id = types.StringValue(datastoreID)

	return nil
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

	err := doCreate(ctx, *r.client, &data)
	if err != nil {
		resp.Diagnostics.AddError(
			"create datastore error",
			err.Error(),
		)
		if errors.As(err, &errordefs.Create) || errors.As(err, &errordefs.Client) {
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

	datastore, err := getDatastore(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.ClusterInfo.Id.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.Diagnostics.AddError(
				"error creating datastore "+data.Name.ValueString(),
				"datastore missing: "+err.Error(),
			)
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error creating datastore"+data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, datastore, &data, &resp.Diagnostics)

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

	datastore, err := getDatastore(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.ClusterInfo.Id.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error reading datastore "+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, datastore, &data, &resp.Diagnostics)

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

	datastore, err := getDatastore(
		ctx,
		*r.client,
		data.HciClusterUuid.ValueString(),
		data.ClusterInfo.Id.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			tflog.Debug(ctx, "datastore not found during delete")

			return
		}
		resp.Diagnostics.AddError(
			"error deleting datastore "+data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	if datastore.GetId() == nil {
		resp.Diagnostics.AddError(
			"error deleting datastore",
			"datastore has no id",
		)

		return
	}
	id := *(datastore.GetId())

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

// Import only grants access to a single "ID" parameter. Therefore, we have to
// combine the "hci_cluster_uuid" and datastore "id" values into the single
// req.ID string
func parseImportID(
	id string,
) (systemID string, hypervisorClusterID string, datastoreName string, error error) {
	params := strings.Split(id, ",")
	if len(params) != 3 || params[0] == "" || params[1] == "" || params[2] == "" {
		return "", "", "", errors.New("invalid import ID format")
	}

	return params[0], params[1], params[2], nil
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	systemID, hypervisorClusterID, datastoreName, err := parseImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"import has invalid datastore id format",
			"Provided import ID \""+req.ID+"\" is invalid. "+
				"Format must be \"<hci_cluster_uuid>,<cluster_info.id>,<datastore_name>\". For example: "+
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3,298a299e-78f5-5acb-86ce-4e9fdc290ab7,my-datastore",
		)

		return
	}

	diags := resp.State.SetAttribute(
		ctx, tfpath.Root("hci_cluster_uuid"), systemID,
	)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(
		ctx, tfpath.Root("cluster_info").AtName("id"), hypervisorClusterID,
	)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(
		ctx, tfpath.Root("name"), datastoreName,
	)
	resp.Diagnostics.Append(diags...)
}
