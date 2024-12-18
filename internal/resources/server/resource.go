package server

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems/privatecloudbusiness"
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

// parseNetworksToPostFormat parses the server network data from the Terraform
// input configuration and returns a slice of server network data which can be
// used in a POST request to the PCBe API to create a new server
// TODO: (API) Issue FF-31496 will prevent this from working currently.
func parseNetworksToPostFormat(dataP *ServerModel) (
	[]privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable,
	error,
) {
	var postRequestNetworks []privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable

	serverNetworks := (*dataP).ServerNetwork
	if serverNetworks.IsNull() {
		msg := "server network list is null"

		return nil, errors.New(msg)

	}

	networks := serverNetworks.Elements()
	if len(networks) != 1 {
		msg := "server network list should be of length 1"

		return nil, errors.New(msg)
	}

	serverNetwork, ok := networks[0].(ServerNetworkValue)
	if !ok {
		msg := "server network element is not a ServerNetworkValue"

		return nil, errors.New(msg)
	}

	infos := serverNetwork.DataIpInfos.Elements()
	if len(infos) != 1 {
		msg := "server network must be an array of length 1"

		return nil, errors.New(msg)
	}

	dataIPInfosValue, ok := infos[0].(DataIpInfosValue)
	if !ok {
		msg := "data ip info element is not a DataIpInfosValue"

		return nil, errors.New(msg)
	}

	ip := dataIPInfosValue.IpAddress.ValueString()
	dataIPInfos := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfos()
	dataIPInfos.SetIpAddress(&ip)
	dataIps := []privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable{
		dataIPInfos,
	}

	iloIPInfos := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfo()

	iloIP := (*dataP).IloNetworkInfo.IloIp.ValueString()
	iloIPInfos.SetIpAddress(&iloIP)

	iloSubnetMask := (*dataP).IloNetworkInfo.SubnetMask.ValueString()
	iloIPInfos.SetSubnetMask(&iloSubnetMask)

	iloGateway := (*dataP).IloNetworkInfo.Gateway.ValueString()
	iloIPInfos.SetGateway(&iloGateway)

	net := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork()
	net.SetDataIpInfos(dataIps)
	net.SetIloMgmtIpInfo(iloIPInfos)
	esxIPAddress := serverNetwork.EsxIpAddress.ValueString()
	net.SetEsxIpAddress(&esxIPAddress)

	postRequestNetworks = []privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable{net}

	return postRequestNetworks, nil
}

func (r *Resource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *Resource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = ServerResourceSchema(ctx)
}

func (r *Resource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.PCBeClient)
}

// doRead reads the server data from the PCBe API and prepares to update
// the model.
// Note: we check that returned fields match the 'required' value specified
// by the user in the Terraform configuration. For computed fields, we
// will set the value to the one returned by the API, even if the value
// has changed.
func doRead(
	ctx context.Context,
	client client.PCBeClient,
	dataP *ServerModel,
	diagsP *diag.Diagnostics,
) {
	serverID := (*dataP).Id.ValueString()
	systemID := (*dataP).SystemId.ValueString()

	sysClient, _, err := client.NewSysClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error reading server",
			"unexpected error: "+err.Error(),
		)

		return
	}

	grc := &privatecloudbusiness.
		V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration{}
	server, err := sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().
		ById(systemID).
		Servers().
		ByServerId(serverID).
		GetAsWithServerGetResponse(ctx, grc)
	if err != nil {
		(*diagsP).AddError(
			"error reading server",
			"get server by ID failed: "+err.Error(),
		)

		return
	}

	if server.GetId() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'id' is nil",
		)

		return
	}

	// If this doesn't match, something is wrong
	if *(server.GetId()) != serverID {
		(*diagsP).AddError(
			"error reading server",
			fmt.Sprintf("'id' mismatch: %s != %s",
				*(server.GetId()), serverID,
			),
		)

		return
	}

	if server.GetSystemId() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'systemID' is nil",
		)

		return
	}

	// If this doesn't match, something is wrong
	if *(server.GetSystemId()) != systemID {
		(*diagsP).AddError(
			"error reading server",
			fmt.Sprintf("'systemID' mismatch: %s != %s",
				*(server.GetSystemId()), systemID,
			),
		)

		return
	}

	if server.GetHypervisorHost() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host' is nil",
		)

		return
	}

	hypervisorClusterID := server.GetHypervisorHost().GetHypervisorClusterId()
	if hypervisorClusterID == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor cluster id' is nil",
		)

		return
	}

	// If this doesn't match, something is wrong
	if *hypervisorClusterID != (*dataP).HypervisorHost.HypervisorClusterId.ValueString() {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor cluster id' mismatch "+
				(*dataP).HypervisorHost.HypervisorClusterId.ValueString()+
				" != "+*hypervisorClusterID,
		)

		return
	}

	hypervisorHostIP := server.GetHypervisorHost().GetHypervisorHostIp()
	if hypervisorHostIP == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host ip' is nil",
		)

		return
	}

	// If this doesn't match, something is wrong
	if *hypervisorHostIP != (*dataP).HypervisorHost.HypervisorHostIp.ValueString() {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host ip' mismatch "+
				(*dataP).HypervisorHost.HypervisorHostIp.ValueString()+
				" != "+*hypervisorHostIP,
		)

		return
	}

	hypervisorClusterName := server.GetHypervisorHost().GetHypervisorClusterName()
	if hypervisorClusterName == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host cluster name' is nil",
		)

		return
	}

	hypervisorHostID := server.GetHypervisorHost().GetId()
	if hypervisorHostID == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host id' is nil",
		)

		return
	}

	hypervisorHostName := server.GetHypervisorHost().GetName()
	if hypervisorHostName == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host name' is nil",
		)

		return
	}

	value := map[string]attr.Value{
		"hypervisor_cluster_id":   types.StringValue(*hypervisorClusterID),
		"hypervisor_cluster_name": types.StringValue(*hypervisorClusterName),
		"hypervisor_host_ip":      types.StringValue(*hypervisorHostIP),
		"id":                      types.StringValue(*hypervisorHostID),
		"name":                    types.StringValue(*hypervisorHostName),
	}

	hypervisorHost, diags := NewHypervisorHostValue(
		(*dataP).HypervisorHost.AttributeTypes(ctx), value,
	)

	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	(*dataP).HypervisorHost = hypervisorHost

	if server.GetSerialNumber() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'serial number' is nil",
		)

		return
	}

	(*dataP).SerialNumber = types.StringValue(*(server.GetSerialNumber()))

	if server.GetName() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'name' is nil",
		)

		return
	}

	(*dataP).Name = types.StringValue(*(server.GetName()))

	// TODO: (API) Add esxRootCredentialId when FF-31524 is addressed
	// TODO: (API) Add iloAdminCredentialId when FF-31525 is addressed

	// Note: esxIpAddress is an outlier: it is part of an RPC-style
	// call. It only exists for the duration of the add-hypervisor-server
	// operation, and is never used again. It does not fit cleanly into
	// terraform's declarative model. We will never be able to read this
	// value back from the server.
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *ServerModel,
	diagsP *diag.Diagnostics,
) {
	hciClusterUUID := (*dataP).HypervisorHost.HypervisorClusterId.ValueString()
	esxRootCredentialID := (*dataP).EsxRootCredentialId.ValueString()
	systemID := (*dataP).SystemId.ValueString()
	iloAdminCredentialID := (*dataP).IloAdminCredentialId.ValueString()
	prc := privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration{}
	prb := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody()
	prb.SetEsxRootCredentialId(&esxRootCredentialID)
	prb.SetIloAdminCredentialId(&iloAdminCredentialID)
	prb.SetHypervisorClusterId(&hciClusterUUID)

	postRequestNetworks, err := parseNetworksToPostFormat(dataP)
	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"could not parse server networks: "+err.Error(),
		)

		return
	}

	prb.SetServerNetwork(postRequestNetworks)

	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"could not create client: "+err.Error(),
		)

		return
	}

	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().ById(systemID).
		AddHypervisorServers().
		Post(ctx, prb, &prc)
	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"unexpected post error: "+err.Error(),
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
		constants.TaskHypervisorServer,
	)

	err = asyncOperation.Poll()
	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"unexpected poll error: "+err.Error(),
		)

		return
	}

	uri, err := asyncOperation.GetAssociatedResourceURI()
	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"failed to get associated resource uri: "+err.Error(),
		)

		return
	}

	serverID := path.Base(uri)
	(*dataP).Id = types.StringValue(serverID)
}

// TODO: (API) Implement create when server create API is implemented
func (r *Resource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data ServerModel

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
	var data ServerModel

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
	tflog.Error(ctx, "update server is not implemented")
}

func (r *Resource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data ServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	serverID := data.Id.ValueString()
	sysClient, sysHeaderOpts, err := r.client.NewSysClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisor server",
			"could not create client: "+err.Error(),
		)

		return
	}

	prc := privatecloudbusiness.
		V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration{}
	prb := privatecloudbusiness.
		NewV1beta1SystemsItemRemoveHypervisorServersPostRequestBody()
	prb.SetServerIds([]string{serverID})

	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().ById(data.SystemId.ValueString()).
		RemoveHypervisorServers().
		Post(ctx, prb, &prc)
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisor server",
			"unexpected post error: "+err.Error(),
		)

		return
	}

	location := sysHeaderOpts.GetResponseHeaders().Get("Location")[0]
	sysHeaderOpts.ResponseHeaders.Clear()
	operationID := path.Base(location)
	asyncOperation := async.New(
		ctx,
		*(r.client),
		operationID,
		constants.TaskHypervisorServer,
	)

	err = asyncOperation.Poll()
	if err != nil {
		resp.Diagnostics.AddError(
			"error deleting hypervisor server",
			"unexpected poll error: "+err.Error(),
		)

		return
	}
	// delete the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, tfpath.Root("id"), req, resp)
}
