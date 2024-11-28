package server

import (
	"context"
	"fmt"
	"path"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/async"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/constants"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems/privatecloudbusiness"
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
			"'systemId' is nil",
		)

		return
	}

	if *(server.GetSystemId()) != systemID {
		(*diagsP).AddError(
			"error reading server",
			fmt.Sprintf("'systemId' mismatch: %s != %s",
				*(server.GetSystemId()), systemID,
			),
		)

		return
	}

	if server.GetName() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'name' is nil",
		)

		return
	}

	(*dataP).Name = types.StringValue(*(server.GetName()))

	if server.GetSerialNumber() == nil {
		(*diagsP).AddError(
			"error reading server",
			"'serial number' is nil",
		)

		return
	}

	(*dataP).SerialNumber = types.StringValue(*(server.GetSerialNumber()))
}

func doCreate(
	ctx context.Context,
	client client.PCBeClient,
	dataP *ServerModel,
	diagsP *diag.Diagnostics,
) {
	hciClusterUUID := (*dataP).HypervisorClusterId.ValueString()
	esxRootCredentialId := (*dataP).EsxRootCredentialId.ValueString()
	systemId := (*dataP).SystemId.ValueString()
	iloAdminCredentialId := (*dataP).IloAdminCredentialId.ValueString()
	sysClient, sysHeaderOpts, err := client.NewSysClient(ctx)

	if err != nil {
		(*diagsP).AddError(
			"error adding hypervisor server",
			"could not create client: "+err.Error(),
		)

		return
	}
	prc := privatecloudbusiness.
		V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration{}
	prb := privatecloudbusiness.
		NewV1beta1SystemsItemAddHypervisorServersPostRequestBody()
	prb.SetEsxRootCredentialId(&esxRootCredentialId)
	prb.SetIloAdminCredentialId(&iloAdminCredentialId)
	prb.SetHypervisorClusterId(&hciClusterUUID)

	_, err = sysClient.PrivateCloudBusiness().
		V1beta1().
		Systems().ById(systemId).
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

	serverId := path.Base(uri)
	(*dataP).Id = types.StringValue(serverId)
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

// TODO: (API) Implement delete when server delete API is implemented
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

	// For now, just delete the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, tfpath.Root("id"), req, resp)
}
