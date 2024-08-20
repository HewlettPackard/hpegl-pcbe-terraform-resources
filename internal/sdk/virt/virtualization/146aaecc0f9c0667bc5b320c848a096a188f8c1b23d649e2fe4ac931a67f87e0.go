package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Private IP address assigned by the cloud provider.
    privateIpAddress *string
    // Public IP address assigned by the cloud provider.
    publicIpAddress *string
    // Information about a subnet to which the machine instance is connected.
    subnetInfo V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable
    // Information about the VPC to which the machine instance network connection is associated.
    vpcInfo V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["subnetInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetInfo(val.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable))
        }
        return nil
    }
    res["vpcInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpcInfo(val.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable))
        }
        return nil
    }
    return res
}
// GetPrivateIpAddress gets the privateIpAddress property value. Private IP address assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetPrivateIpAddress()(*string) {
    return m.privateIpAddress
}
// GetPublicIpAddress gets the publicIpAddress property value. Public IP address assigned by the cloud provider.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetPublicIpAddress()(*string) {
    return m.publicIpAddress
}
// GetSubnetInfo gets the subnetInfo property value. Information about a subnet to which the machine instance is connected.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetSubnetInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable) {
    return m.subnetInfo
}
// GetVpcInfo gets the vpcInfo property value. Information about the VPC to which the machine instance network connection is associated.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) GetVpcInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable) {
    return m.vpcInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateIpAddress sets the privateIpAddress property value. Private IP address assigned by the cloud provider.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) SetPrivateIpAddress(value *string)() {
    m.privateIpAddress = value
}
// SetPublicIpAddress sets the publicIpAddress property value. Public IP address assigned by the cloud provider.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) SetPublicIpAddress(value *string)() {
    m.publicIpAddress = value
}
// SetSubnetInfo sets the subnetInfo property value. Information about a subnet to which the machine instance is connected.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) SetSubnetInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable)() {
    m.subnetInfo = value
}
// SetVpcInfo sets the vpcInfo property value. Information about the VPC to which the machine instance network connection is associated.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections) SetVpcInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable)() {
    m.vpcInfo = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateIpAddress()(*string)
    GetPublicIpAddress()(*string)
    GetSubnetInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable)
    GetVpcInfo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable)
    SetPrivateIpAddress(value *string)()
    SetPublicIpAddress(value *string)()
    SetSubnetInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_subnetInfoable)()
    SetVpcInfo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connections_vpcInfoable)()
}
