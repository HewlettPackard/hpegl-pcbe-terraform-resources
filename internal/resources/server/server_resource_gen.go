// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package server

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"esx_root_credential_id": schema.StringAttribute{
				Required:            true,
				Description:         "ID corresponding to the ESXi password",
				MarkdownDescription: "ID corresponding to the ESXi password",
			},
			"hypervisor_cluster_id": schema.StringAttribute{
				Required:            true,
				Description:         "hypervisor cluster UUID",
				MarkdownDescription: "hypervisor cluster UUID",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "An identifier for the resource, usually a UUID.",
				MarkdownDescription: "An identifier for the resource, usually a UUID.",
			},
			"ilo_admin_credential_id": schema.StringAttribute{
				Required:            true,
				Description:         "ID corresponding to the ilo admin credential",
				MarkdownDescription: "ID corresponding to the ilo admin credential",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Description:         "A system specified name for the resource.",
				MarkdownDescription: "A system specified name for the resource.",
			},
			"serial_number": schema.StringAttribute{
				Computed:            true,
				Description:         "Server's serial number.",
				MarkdownDescription: "Server's serial number.",
			},
			"server_network": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data_ip_infos": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Required:            true,
										Description:         "IP address",
										MarkdownDescription: "IP address",
									},
								},
								CustomType: DataIpInfosType{
									ObjectType: types.ObjectType{
										AttrTypes: DataIpInfosValue{}.AttributeTypes(ctx),
									},
								},
							},
							Required:            true,
							Description:         "Details of data network",
							MarkdownDescription: "Details of data network",
						},
					},
					CustomType: ServerNetworkType{
						ObjectType: types.ObjectType{
							AttrTypes: ServerNetworkValue{}.AttributeTypes(ctx),
						},
					},
				},
				Required:            true,
				Description:         "Details of network (and other info) for a server.",
				MarkdownDescription: "Details of network (and other info) for a server.",
			},
			"system_id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique identifier of the system, usually a UUID.",
				MarkdownDescription: "Unique identifier of the system, usually a UUID.",
			},
		},
	}
}

type ServerModel struct {
	EsxRootCredentialId  types.String `tfsdk:"esx_root_credential_id"`
	HypervisorClusterId  types.String `tfsdk:"hypervisor_cluster_id"`
	Id                   types.String `tfsdk:"id"`
	IloAdminCredentialId types.String `tfsdk:"ilo_admin_credential_id"`
	Name                 types.String `tfsdk:"name"`
	SerialNumber         types.String `tfsdk:"serial_number"`
	ServerNetwork        types.List   `tfsdk:"server_network"`
	SystemId             types.String `tfsdk:"system_id"`
}

var _ basetypes.ObjectTypable = ServerNetworkType{}

type ServerNetworkType struct {
	basetypes.ObjectType
}

func (t ServerNetworkType) Equal(o attr.Type) bool {
	other, ok := o.(ServerNetworkType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ServerNetworkType) String() string {
	return "ServerNetworkType"
}

func (t ServerNetworkType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	dataIpInfosAttribute, ok := attributes["data_ip_infos"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`data_ip_infos is missing from object`)

		return nil, diags
	}

	dataIpInfosVal, ok := dataIpInfosAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`data_ip_infos expected to be basetypes.ListValue, was: %T`, dataIpInfosAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ServerNetworkValue{
		DataIpInfos: dataIpInfosVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewServerNetworkValueNull() ServerNetworkValue {
	return ServerNetworkValue{
		state: attr.ValueStateNull,
	}
}

func NewServerNetworkValueUnknown() ServerNetworkValue {
	return ServerNetworkValue{
		state: attr.ValueStateUnknown,
	}
}

func NewServerNetworkValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ServerNetworkValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ServerNetworkValue Attribute Value",
				"While creating a ServerNetworkValue value, a missing attribute value was detected. "+
					"A ServerNetworkValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServerNetworkValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ServerNetworkValue Attribute Type",
				"While creating a ServerNetworkValue value, an invalid attribute value was detected. "+
					"A ServerNetworkValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServerNetworkValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ServerNetworkValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ServerNetworkValue Attribute Value",
				"While creating a ServerNetworkValue value, an extra attribute value was detected. "+
					"A ServerNetworkValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ServerNetworkValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewServerNetworkValueUnknown(), diags
	}

	dataIpInfosAttribute, ok := attributes["data_ip_infos"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`data_ip_infos is missing from object`)

		return NewServerNetworkValueUnknown(), diags
	}

	dataIpInfosVal, ok := dataIpInfosAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`data_ip_infos expected to be basetypes.ListValue, was: %T`, dataIpInfosAttribute))
	}

	if diags.HasError() {
		return NewServerNetworkValueUnknown(), diags
	}

	return ServerNetworkValue{
		DataIpInfos: dataIpInfosVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewServerNetworkValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ServerNetworkValue {
	object, diags := NewServerNetworkValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewServerNetworkValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ServerNetworkType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewServerNetworkValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewServerNetworkValueUnknown(), nil
	}

	if in.IsNull() {
		return NewServerNetworkValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewServerNetworkValueMust(ServerNetworkValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ServerNetworkType) ValueType(ctx context.Context) attr.Value {
	return ServerNetworkValue{}
}

var _ basetypes.ObjectValuable = ServerNetworkValue{}

type ServerNetworkValue struct {
	DataIpInfos basetypes.ListValue `tfsdk:"data_ip_infos"`
	state       attr.ValueState
}

func (v ServerNetworkValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["data_ip_infos"] = basetypes.ListType{
		ElemType: DataIpInfosValue{}.Type(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.DataIpInfos.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["data_ip_infos"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ServerNetworkValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ServerNetworkValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ServerNetworkValue) String() string {
	return "ServerNetworkValue"
}

func (v ServerNetworkValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	dataIpInfos := types.ListValueMust(
		DataIpInfosType{
			basetypes.ObjectType{
				AttrTypes: DataIpInfosValue{}.AttributeTypes(ctx),
			},
		},
		v.DataIpInfos.Elements(),
	)

	if v.DataIpInfos.IsNull() {
		dataIpInfos = types.ListNull(
			DataIpInfosType{
				basetypes.ObjectType{
					AttrTypes: DataIpInfosValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	if v.DataIpInfos.IsUnknown() {
		dataIpInfos = types.ListUnknown(
			DataIpInfosType{
				basetypes.ObjectType{
					AttrTypes: DataIpInfosValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	attributeTypes := map[string]attr.Type{
		"data_ip_infos": basetypes.ListType{
			ElemType: DataIpInfosValue{}.Type(ctx),
		},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"data_ip_infos": dataIpInfos,
		})

	return objVal, diags
}

func (v ServerNetworkValue) Equal(o attr.Value) bool {
	other, ok := o.(ServerNetworkValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.DataIpInfos.Equal(other.DataIpInfos) {
		return false
	}

	return true
}

func (v ServerNetworkValue) Type(ctx context.Context) attr.Type {
	return ServerNetworkType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ServerNetworkValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"data_ip_infos": basetypes.ListType{
			ElemType: DataIpInfosValue{}.Type(ctx),
		},
	}
}

var _ basetypes.ObjectTypable = DataIpInfosType{}

type DataIpInfosType struct {
	basetypes.ObjectType
}

func (t DataIpInfosType) Equal(o attr.Type) bool {
	other, ok := o.(DataIpInfosType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t DataIpInfosType) String() string {
	return "DataIpInfosType"
}

func (t DataIpInfosType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	ipAddressAttribute, ok := attributes["ip_address"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip_address is missing from object`)

		return nil, diags
	}

	ipAddressVal, ok := ipAddressAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip_address expected to be basetypes.StringValue, was: %T`, ipAddressAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return DataIpInfosValue{
		IpAddress: ipAddressVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewDataIpInfosValueNull() DataIpInfosValue {
	return DataIpInfosValue{
		state: attr.ValueStateNull,
	}
}

func NewDataIpInfosValueUnknown() DataIpInfosValue {
	return DataIpInfosValue{
		state: attr.ValueStateUnknown,
	}
}

func NewDataIpInfosValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (DataIpInfosValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing DataIpInfosValue Attribute Value",
				"While creating a DataIpInfosValue value, a missing attribute value was detected. "+
					"A DataIpInfosValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DataIpInfosValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid DataIpInfosValue Attribute Type",
				"While creating a DataIpInfosValue value, an invalid attribute value was detected. "+
					"A DataIpInfosValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DataIpInfosValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("DataIpInfosValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra DataIpInfosValue Attribute Value",
				"While creating a DataIpInfosValue value, an extra attribute value was detected. "+
					"A DataIpInfosValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra DataIpInfosValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewDataIpInfosValueUnknown(), diags
	}

	ipAddressAttribute, ok := attributes["ip_address"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip_address is missing from object`)

		return NewDataIpInfosValueUnknown(), diags
	}

	ipAddressVal, ok := ipAddressAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip_address expected to be basetypes.StringValue, was: %T`, ipAddressAttribute))
	}

	if diags.HasError() {
		return NewDataIpInfosValueUnknown(), diags
	}

	return DataIpInfosValue{
		IpAddress: ipAddressVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewDataIpInfosValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) DataIpInfosValue {
	object, diags := NewDataIpInfosValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewDataIpInfosValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t DataIpInfosType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewDataIpInfosValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewDataIpInfosValueUnknown(), nil
	}

	if in.IsNull() {
		return NewDataIpInfosValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewDataIpInfosValueMust(DataIpInfosValue{}.AttributeTypes(ctx), attributes), nil
}

func (t DataIpInfosType) ValueType(ctx context.Context) attr.Value {
	return DataIpInfosValue{}
}

var _ basetypes.ObjectValuable = DataIpInfosValue{}

type DataIpInfosValue struct {
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
	state     attr.ValueState
}

func (v DataIpInfosValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["ip_address"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.IpAddress.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ip_address"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v DataIpInfosValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v DataIpInfosValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v DataIpInfosValue) String() string {
	return "DataIpInfosValue"
}

func (v DataIpInfosValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"ip_address": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"ip_address": v.IpAddress,
		})

	return objVal, diags
}

func (v DataIpInfosValue) Equal(o attr.Value) bool {
	other, ok := o.(DataIpInfosValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.IpAddress.Equal(other.IpAddress) {
		return false
	}

	return true
}

func (v DataIpInfosValue) Type(ctx context.Context) attr.Type {
	return DataIpInfosType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v DataIpInfosValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"ip_address": basetypes.StringType{},
	}
}
