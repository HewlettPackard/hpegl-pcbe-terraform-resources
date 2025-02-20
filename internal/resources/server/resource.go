package server

import (
	"context"
	"errors"
	"path"
	"strings"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/errordefs"
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
// used in a POST request to the PCBe API to create a new server TODO: (API)
// Issue FF-31496/FF-31582/etc will prevent this from working currently,
// specifically we don't get enough info back from the API to populate
// server_network fields.
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

func createServerFilter(
	hypervisorClusterID string,
	name string,
) string {
	return constants.NameFilter + "'" + name + "'" +
		constants.AndFilter +
		constants.ServerHypervisorClusterIDFilter +
		"'" + hypervisorClusterID + "'"
}

func getServer(
	ctx context.Context,
	client client.PCBeClient,
	hciClusterUUID string, // "system" ID
	hypervisorClusterID string, // "cluster" ID
	name string, // server name
) (privatecloudbusiness.V1beta1SystemsItemServersGetResponse_itemsable, error) {
	sysClient, _, err := client.NewSysClient(ctx)
	if err != nil {
		return nil, err
	}

	qp := privatecloudbusiness.
		V1beta1SystemsItemServersRequestBuilderGetQueryParameters{}
	filter := createServerFilter(hypervisorClusterID, name)
	qp.Filter = &filter

	grc := privatecloudbusiness.
		V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration{}
	grc.QueryParameters = &qp

	servers, err := sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().
		ById(hciClusterUUID).
		Servers().
		GetAsServersGetResponse(ctx, &grc)
	if err != nil {
		return nil, err
	}

	if servers.GetTotal() == nil {
		return nil, errors.New("servers 'total' field is nil")
	}

	total := *(servers.GetTotal())
	if total == 0 {
		// server does not exist
		return nil, errordefs.NewNotFoundError(name)
	}

	if total != 1 {
		msg := "error getting server ID for " + name

		return nil, errors.New(msg)
	}

	return servers.GetItems()[0], nil
}

// doRead reads the server data from the PCBe API and prepares to update
// the model.
// Note: we check that returned fields match the 'required' value specified
// by the user in the Terraform configuration. For computed fields, we
// will set the value to the one returned by the API, even if the value
// has changed.
func doRead(
	ctx context.Context,
	server privatecloudbusiness.V1beta1SystemsItemServersGetResponse_itemsable,
	dataP *ServerModel,
	diagsP *diag.Diagnostics,
) {
	if server.GetId() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'id' is nil",
		)

		return
	}
	(*dataP).Id = types.StringValue(*(server.GetId()))

	if server.GetSystemId() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'systemID' is nil",
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

	hypervisorClusterID := server.
		GetHypervisorHost().
		GetHypervisorClusterId()
	if hypervisorClusterID == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor cluster id' is nil",
		)

		return
	}

	hypervisorHostIP := server.
		GetHypervisorHost().
		GetHypervisorHostIp()
	if hypervisorHostIP == nil {
		(*diagsP).AddError(
			"error reading server",
			"'hypervisor host ip' is nil",
		)

		return
	}

	hypervisorClusterName := server.
		GetHypervisorHost().
		GetHypervisorClusterName()
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

	iloNetInfo := server.GetIloNetworkInfo()
	if iloNetInfo == nil {
		(*diagsP).AddError(
			"error reading server",
			"'ilo network info' is nil",
		)

		return
	}
	iloGateway := server.GetIloNetworkInfo().GetGateway()
	if iloGateway == nil {
		(*diagsP).AddError(
			"error reading server",
			"'ilo network info gateway' is nil",
		)

		return
	}

	iloIP := server.GetIloNetworkInfo().GetIloIp()
	if iloIP == nil {
		(*diagsP).AddError(
			"error reading server",
			"'ilo network info ip address' is nil",
		)

		return
	}

	iloSubnetMask := server.GetIloNetworkInfo().GetSubnetMask()
	if iloSubnetMask == nil {
		(*diagsP).AddError(
			"error reading server",
			"'ilo network info subnet mask' is nil",
		)

		return
	}

	value = map[string]attr.Value{
		"gateway":     types.StringValue(*iloGateway),
		"ilo_ip":      types.StringValue(*iloIP),
		"subnet_mask": types.StringValue(*iloSubnetMask),
	}

	iloNetworkInfo, diags := NewIloNetworkInfoValue(
		(*dataP).IloNetworkInfo.AttributeTypes(ctx), value,
	)

	(*diagsP).Append(diags...)
	if (*diagsP).HasError() {
		return
	}

	(*dataP).IloNetworkInfo = iloNetworkInfo

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
) error {
	name := (*dataP).Name.ValueString()
	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)
	if err != nil {
		tflog.Error(ctx, "failed to create client "+err.Error())

		return errordefs.NewClientError(name)

	}

	hciClusterUUID := (*dataP).HypervisorHost.HypervisorClusterId.ValueString()
	esxRootCredentialID := (*dataP).EsxRootCredentialId.ValueString()
	systemID := (*dataP).SystemId.ValueString()
	iloAdminCredentialID := (*dataP).IloAdminCredentialId.ValueString()
	prb := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody()
	prb.SetEsxRootCredentialId(&esxRootCredentialID)
	prb.SetIloAdminCredentialId(&iloAdminCredentialID)
	prb.SetHypervisorClusterId(&hciClusterUUID)

	postRequestNetworks, err := parseNetworksToPostFormat(dataP)
	if err != nil {
		tflog.Error(ctx,
			"error adding hypervisor server "+
				"could not parse server networks: "+
				err.Error(),
		)

		return errordefs.NewNetParseError(name)
	}

	prb.SetServerNetwork(postRequestNetworks)
	prc := privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration{}

	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().ById(systemID).
		AddHypervisorServers().
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
		constants.TaskHypervisorServer,
	)
	err = asyncOperation.Poll()
	if err != nil {
		tflog.Error(ctx, "failed to poll resource "+err.Error())

		return errordefs.NewPollError(name)
	}

	uri, err := asyncOperation.GetAssociatedResourceURI()
	if err != nil {
		tflog.Error(
			ctx, "failed to get associatedResourceUri: "+err.Error(),
		)

		return errordefs.NewNoURIError(name)
	}

	serverID := path.Base(uri)
	(*dataP).Id = types.StringValue(serverID)

	return nil
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

	err := doCreate(ctx, *r.client, &data)
	if err != nil {
		resp.Diagnostics.AddError(
			"create server error",
			err.Error(),
		)
		if errors.As(err, &errordefs.Create) ||
			errors.As(err, &errordefs.Client) ||
			errors.As(err, &errordefs.NetParse) {
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

	server, err := getServer(
		ctx,
		*r.client,
		data.SystemId.ValueString(),
		data.HypervisorHost.HypervisorClusterId.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.Diagnostics.AddError(
				"error creating server "+data.Name.ValueString(),
				"server missing: "+err.Error(),
			)
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error creating server "+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, server, &data, &resp.Diagnostics)
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

	server, err := getServer(
		ctx,
		*r.client,
		data.SystemId.ValueString(),
		data.HypervisorHost.HypervisorClusterId.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		if errors.As(err, &errordefs.NotFound) {
			// gone missing, purge state
			resp.State.RemoveResource(ctx)

			return
		}

		resp.Diagnostics.AddError(
			"error reading server "+
				data.Name.ValueString(),
			"unexpected error: "+err.Error(),
		)

		return
	}

	doRead(ctx, server, &data, &resp.Diagnostics)

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

// Import only grants access to a single "ID" parameter. Therefore, we have to
// combine the "hci_cluster_uuid" and datastore "id" values into the single
// req.ID string
func parseImportID(
	id string,
) (systemID string, clusterID string, serverName string, error error) {
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
	systemID, hypervisorClusterID, serverName, err := parseImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"import has invalid server id format",
			"Provided import ID \""+req.ID+"\" is invalid. "+
				"Format must be \"<system_id>,<hypervisor_host.hypervisor_cluster_id>,<server_name>. "+
				"For example: "+
				"126fd201-9e6e-5e31-9ffb-a766265b1fd3,298a299e-78f5-5acb-86ce-4e9fdc290ab7,server1",
		)

		return
	}

	diags := resp.State.SetAttribute(
		ctx, tfpath.Root("hypervisor_host").AtName("hypervisor_cluster_id"),
		hypervisorClusterID,
	)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.SetAttribute(
		ctx, tfpath.Root("system_id"), systemID,
	)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.SetAttribute(
		ctx, tfpath.Root("name"), serverName,
	)
	resp.Diagnostics.Append(diags...)
}
