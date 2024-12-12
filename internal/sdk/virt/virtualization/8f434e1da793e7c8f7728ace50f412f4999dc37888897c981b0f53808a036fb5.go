package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Private IP address assigned by the cloud provider.
    privateIpAddress *string
    // Public IP address assigned by the cloud provider.
    publicIpAddress *string
    // Indicates whether the public IP address is a floating address assignment.
    publicIpAddressIsFloating *bool
    // The cloud provider security groups for the machine instance.
    securityGroups []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable
    // Information about the subnet to which the machine instance is connected.
    subnetInfo V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable
    // Information about the VPC to which the machine instance is associated.
    vpcInfo V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable
}
// NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo instantiates a new V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo and sets the default values.
func NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo()(*V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) {
    m := &V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["privateIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateIpAddress(val)
        }
        return nil
    }
    res["publicIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIpAddress(val)
        }
        return nil
    }
    res["publicIpAddressIsFloating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIpAddressIsFloating(val)
        }
        return nil
    }
    res["securityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable)
                }
            }
            m.SetSecurityGroups(res)
        }
        return nil
    }
    res["subnetInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetInfo(val.(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable))
        }
        return nil
    }
    res["vpcInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpcInfo(val.(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable))
        }
        return nil
    }
    return res
}
// GetPrivateIpAddress gets the privateIpAddress property value. Private IP address assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetPrivateIpAddress()(*string) {
    return m.privateIpAddress
}
// GetPublicIpAddress gets the publicIpAddress property value. Public IP address assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetPublicIpAddress()(*string) {
    return m.publicIpAddress
}
// GetPublicIpAddressIsFloating gets the publicIpAddressIsFloating property value. Indicates whether the public IP address is a floating address assignment.
// returns a *bool when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetPublicIpAddressIsFloating()(*bool) {
    return m.publicIpAddressIsFloating
}
// GetSecurityGroups gets the securityGroups property value. The cloud provider security groups for the machine instance.
// returns a []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetSecurityGroups()([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable) {
    return m.securityGroups
}
// GetSubnetInfo gets the subnetInfo property value. Information about the subnet to which the machine instance is connected.
// returns a V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetSubnetInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable) {
    return m.subnetInfo
}
// GetVpcInfo gets the vpcInfo property value. Information about the VPC to which the machine instance is associated.
// returns a V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable when successful
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) GetVpcInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable) {
    return m.vpcInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateIpAddress sets the privateIpAddress property value. Private IP address assigned by the cloud provider.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetPrivateIpAddress(value *string)() {
    m.privateIpAddress = value
}
// SetPublicIpAddress sets the publicIpAddress property value. Public IP address assigned by the cloud provider.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetPublicIpAddress(value *string)() {
    m.publicIpAddress = value
}
// SetPublicIpAddressIsFloating sets the publicIpAddressIsFloating property value. Indicates whether the public IP address is a floating address assignment.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetPublicIpAddressIsFloating(value *bool)() {
    m.publicIpAddressIsFloating = value
}
// SetSecurityGroups sets the securityGroups property value. The cloud provider security groups for the machine instance.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetSecurityGroups(value []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable)() {
    m.securityGroups = value
}
// SetSubnetInfo sets the subnetInfo property value. Information about the subnet to which the machine instance is connected.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetSubnetInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable)() {
    m.subnetInfo = value
}
// SetVpcInfo sets the vpcInfo property value. Information about the VPC to which the machine instance is associated.
func (m *V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo) SetVpcInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable)() {
    m.vpcInfo = value
}
type V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateIpAddress()(*string)
    GetPublicIpAddress()(*string)
    GetPublicIpAddressIsFloating()(*bool)
    GetSecurityGroups()([]V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable)
    GetSubnetInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable)
    GetVpcInfo()(V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable)
    SetPrivateIpAddress(value *string)()
    SetPublicIpAddress(value *string)()
    SetPublicIpAddressIsFloating(value *bool)()
    SetSecurityGroups(value []V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_securityGroupsable)()
    SetSubnetInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_subnetInfoable)()
    SetVpcInfo(value V1beta1CspMachineInstancesGetResponse_items_cspInfoMember1_networkInfo_vpcInfoable)()
}
