// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package hypervisorcluster

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

func HypervisorclusterResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_info": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vmware": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"datacenter_info": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required:            true,
										Description:         "VMware provided name for the datacenter.",
										MarkdownDescription: "VMware provided name for the datacenter.",
									},
								},
								CustomType: DatacenterInfoType{
									ObjectType: types.ObjectType{
										AttrTypes: DatacenterInfoValue{}.AttributeTypes(ctx),
									},
								},
								Required:            true,
								Description:         "References to the datacenter that house this virtual machine.",
								MarkdownDescription: "References to the datacenter that house this virtual machine.",
							},
						},
						CustomType: VmwareType{
							ObjectType: types.ObjectType{
								AttrTypes: VmwareValue{}.AttributeTypes(ctx),
							},
						},
						Required: true,
					},
				},
				CustomType: AppInfoType{
					ObjectType: types.ObjectType{
						AttrTypes: AppInfoValue{}.AttributeTypes(ctx),
					},
				},
				Required:            true,
				Description:         "Application specific information for this cluster.",
				MarkdownDescription: "Application specific information for this cluster.",
			},
			"hci_cluster_uuid": schema.StringAttribute{
				Required:            true,
				Description:         "UUID string uniquely identifying the HCI cluster.",
				MarkdownDescription: "UUID string uniquely identifying the HCI cluster.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "UUID string uniquely identifying the hypervisor cluster.",
				MarkdownDescription: "UUID string uniquely identifying the hypervisor cluster.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the cluster as reported by the hypervisor manager.",
				MarkdownDescription: "Name of the cluster as reported by the hypervisor manager.",
			},
		},
	}
}

type HypervisorclusterModel struct {
	AppInfo        AppInfoValue `tfsdk:"app_info"`
	HciClusterUuid types.String `tfsdk:"hci_cluster_uuid"`
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
}

var _ basetypes.ObjectTypable = AppInfoType{}

type AppInfoType struct {
	basetypes.ObjectType
}

func (t AppInfoType) Equal(o attr.Type) bool {
	other, ok := o.(AppInfoType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t AppInfoType) String() string {
	return "AppInfoType"
}

func (t AppInfoType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	vmwareAttribute, ok := attributes["vmware"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vmware is missing from object`)

		return nil, diags
	}

	vmwareVal, ok := vmwareAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vmware expected to be basetypes.ObjectValue, was: %T`, vmwareAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return AppInfoValue{
		Vmware: vmwareVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewAppInfoValueNull() AppInfoValue {
	return AppInfoValue{
		state: attr.ValueStateNull,
	}
}

func NewAppInfoValueUnknown() AppInfoValue {
	return AppInfoValue{
		state: attr.ValueStateUnknown,
	}
}

func NewAppInfoValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (AppInfoValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing AppInfoValue Attribute Value",
				"While creating a AppInfoValue value, a missing attribute value was detected. "+
					"A AppInfoValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AppInfoValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid AppInfoValue Attribute Type",
				"While creating a AppInfoValue value, an invalid attribute value was detected. "+
					"A AppInfoValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AppInfoValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("AppInfoValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra AppInfoValue Attribute Value",
				"While creating a AppInfoValue value, an extra attribute value was detected. "+
					"A AppInfoValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra AppInfoValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewAppInfoValueUnknown(), diags
	}

	vmwareAttribute, ok := attributes["vmware"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vmware is missing from object`)

		return NewAppInfoValueUnknown(), diags
	}

	vmwareVal, ok := vmwareAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vmware expected to be basetypes.ObjectValue, was: %T`, vmwareAttribute))
	}

	if diags.HasError() {
		return NewAppInfoValueUnknown(), diags
	}

	return AppInfoValue{
		Vmware: vmwareVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewAppInfoValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) AppInfoValue {
	object, diags := NewAppInfoValue(attributeTypes, attributes)

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

		panic("NewAppInfoValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t AppInfoType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewAppInfoValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewAppInfoValueUnknown(), nil
	}

	if in.IsNull() {
		return NewAppInfoValueNull(), nil
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

	return NewAppInfoValueMust(AppInfoValue{}.AttributeTypes(ctx), attributes), nil
}

func (t AppInfoType) ValueType(ctx context.Context) attr.Value {
	return AppInfoValue{}
}

var _ basetypes.ObjectValuable = AppInfoValue{}

type AppInfoValue struct {
	Vmware basetypes.ObjectValue `tfsdk:"vmware"`
	state  attr.ValueState
}

func (v AppInfoValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["vmware"] = basetypes.ObjectType{
		AttrTypes: VmwareValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.Vmware.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["vmware"] = val

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

func (v AppInfoValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v AppInfoValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v AppInfoValue) String() string {
	return "AppInfoValue"
}

func (v AppInfoValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var vmware basetypes.ObjectValue

	if v.Vmware.IsNull() {
		vmware = types.ObjectNull(
			VmwareValue{}.AttributeTypes(ctx),
		)
	}

	if v.Vmware.IsUnknown() {
		vmware = types.ObjectUnknown(
			VmwareValue{}.AttributeTypes(ctx),
		)
	}

	if !v.Vmware.IsNull() && !v.Vmware.IsUnknown() {
		vmware = types.ObjectValueMust(
			VmwareValue{}.AttributeTypes(ctx),
			v.Vmware.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"vmware": basetypes.ObjectType{
			AttrTypes: VmwareValue{}.AttributeTypes(ctx),
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
			"vmware": vmware,
		})

	return objVal, diags
}

func (v AppInfoValue) Equal(o attr.Value) bool {
	other, ok := o.(AppInfoValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Vmware.Equal(other.Vmware) {
		return false
	}

	return true
}

func (v AppInfoValue) Type(ctx context.Context) attr.Type {
	return AppInfoType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v AppInfoValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"vmware": basetypes.ObjectType{
			AttrTypes: VmwareValue{}.AttributeTypes(ctx),
		},
	}
}

var _ basetypes.ObjectTypable = VmwareType{}

type VmwareType struct {
	basetypes.ObjectType
}

func (t VmwareType) Equal(o attr.Type) bool {
	other, ok := o.(VmwareType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t VmwareType) String() string {
	return "VmwareType"
}

func (t VmwareType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	datacenterInfoAttribute, ok := attributes["datacenter_info"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`datacenter_info is missing from object`)

		return nil, diags
	}

	datacenterInfoVal, ok := datacenterInfoAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`datacenter_info expected to be basetypes.ObjectValue, was: %T`, datacenterInfoAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return VmwareValue{
		DatacenterInfo: datacenterInfoVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewVmwareValueNull() VmwareValue {
	return VmwareValue{
		state: attr.ValueStateNull,
	}
}

func NewVmwareValueUnknown() VmwareValue {
	return VmwareValue{
		state: attr.ValueStateUnknown,
	}
}

func NewVmwareValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (VmwareValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing VmwareValue Attribute Value",
				"While creating a VmwareValue value, a missing attribute value was detected. "+
					"A VmwareValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("VmwareValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid VmwareValue Attribute Type",
				"While creating a VmwareValue value, an invalid attribute value was detected. "+
					"A VmwareValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("VmwareValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("VmwareValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra VmwareValue Attribute Value",
				"While creating a VmwareValue value, an extra attribute value was detected. "+
					"A VmwareValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra VmwareValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewVmwareValueUnknown(), diags
	}

	datacenterInfoAttribute, ok := attributes["datacenter_info"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`datacenter_info is missing from object`)

		return NewVmwareValueUnknown(), diags
	}

	datacenterInfoVal, ok := datacenterInfoAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`datacenter_info expected to be basetypes.ObjectValue, was: %T`, datacenterInfoAttribute))
	}

	if diags.HasError() {
		return NewVmwareValueUnknown(), diags
	}

	return VmwareValue{
		DatacenterInfo: datacenterInfoVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewVmwareValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) VmwareValue {
	object, diags := NewVmwareValue(attributeTypes, attributes)

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

		panic("NewVmwareValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t VmwareType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewVmwareValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewVmwareValueUnknown(), nil
	}

	if in.IsNull() {
		return NewVmwareValueNull(), nil
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

	return NewVmwareValueMust(VmwareValue{}.AttributeTypes(ctx), attributes), nil
}

func (t VmwareType) ValueType(ctx context.Context) attr.Value {
	return VmwareValue{}
}

var _ basetypes.ObjectValuable = VmwareValue{}

type VmwareValue struct {
	DatacenterInfo basetypes.ObjectValue `tfsdk:"datacenter_info"`
	state          attr.ValueState
}

func (v VmwareValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["datacenter_info"] = basetypes.ObjectType{
		AttrTypes: DatacenterInfoValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.DatacenterInfo.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["datacenter_info"] = val

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

func (v VmwareValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v VmwareValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v VmwareValue) String() string {
	return "VmwareValue"
}

func (v VmwareValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var datacenterInfo basetypes.ObjectValue

	if v.DatacenterInfo.IsNull() {
		datacenterInfo = types.ObjectNull(
			DatacenterInfoValue{}.AttributeTypes(ctx),
		)
	}

	if v.DatacenterInfo.IsUnknown() {
		datacenterInfo = types.ObjectUnknown(
			DatacenterInfoValue{}.AttributeTypes(ctx),
		)
	}

	if !v.DatacenterInfo.IsNull() && !v.DatacenterInfo.IsUnknown() {
		datacenterInfo = types.ObjectValueMust(
			DatacenterInfoValue{}.AttributeTypes(ctx),
			v.DatacenterInfo.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"datacenter_info": basetypes.ObjectType{
			AttrTypes: DatacenterInfoValue{}.AttributeTypes(ctx),
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
			"datacenter_info": datacenterInfo,
		})

	return objVal, diags
}

func (v VmwareValue) Equal(o attr.Value) bool {
	other, ok := o.(VmwareValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.DatacenterInfo.Equal(other.DatacenterInfo) {
		return false
	}

	return true
}

func (v VmwareValue) Type(ctx context.Context) attr.Type {
	return VmwareType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v VmwareValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"datacenter_info": basetypes.ObjectType{
			AttrTypes: DatacenterInfoValue{}.AttributeTypes(ctx),
		},
	}
}

var _ basetypes.ObjectTypable = DatacenterInfoType{}

type DatacenterInfoType struct {
	basetypes.ObjectType
}

func (t DatacenterInfoType) Equal(o attr.Type) bool {
	other, ok := o.(DatacenterInfoType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t DatacenterInfoType) String() string {
	return "DatacenterInfoType"
}

func (t DatacenterInfoType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return DatacenterInfoValue{
		Name:  nameVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewDatacenterInfoValueNull() DatacenterInfoValue {
	return DatacenterInfoValue{
		state: attr.ValueStateNull,
	}
}

func NewDatacenterInfoValueUnknown() DatacenterInfoValue {
	return DatacenterInfoValue{
		state: attr.ValueStateUnknown,
	}
}

func NewDatacenterInfoValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (DatacenterInfoValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing DatacenterInfoValue Attribute Value",
				"While creating a DatacenterInfoValue value, a missing attribute value was detected. "+
					"A DatacenterInfoValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DatacenterInfoValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid DatacenterInfoValue Attribute Type",
				"While creating a DatacenterInfoValue value, an invalid attribute value was detected. "+
					"A DatacenterInfoValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DatacenterInfoValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("DatacenterInfoValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra DatacenterInfoValue Attribute Value",
				"While creating a DatacenterInfoValue value, an extra attribute value was detected. "+
					"A DatacenterInfoValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra DatacenterInfoValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewDatacenterInfoValueUnknown(), diags
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewDatacenterInfoValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return NewDatacenterInfoValueUnknown(), diags
	}

	return DatacenterInfoValue{
		Name:  nameVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewDatacenterInfoValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) DatacenterInfoValue {
	object, diags := NewDatacenterInfoValue(attributeTypes, attributes)

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

		panic("NewDatacenterInfoValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t DatacenterInfoType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewDatacenterInfoValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewDatacenterInfoValueUnknown(), nil
	}

	if in.IsNull() {
		return NewDatacenterInfoValueNull(), nil
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

	return NewDatacenterInfoValueMust(DatacenterInfoValue{}.AttributeTypes(ctx), attributes), nil
}

func (t DatacenterInfoType) ValueType(ctx context.Context) attr.Value {
	return DatacenterInfoValue{}
}

var _ basetypes.ObjectValuable = DatacenterInfoValue{}

type DatacenterInfoValue struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	state attr.ValueState
}

func (v DatacenterInfoValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

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

func (v DatacenterInfoValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v DatacenterInfoValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v DatacenterInfoValue) String() string {
	return "DatacenterInfoValue"
}

func (v DatacenterInfoValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"name": basetypes.StringType{},
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
			"name": v.Name,
		})

	return objVal, diags
}

func (v DatacenterInfoValue) Equal(o attr.Value) bool {
	other, ok := o.(DatacenterInfoValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	return true
}

func (v DatacenterInfoValue) Type(ctx context.Context) attr.Type {
	return DatacenterInfoType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v DatacenterInfoValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}
