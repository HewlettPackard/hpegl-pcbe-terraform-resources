package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the network interface.
    name *string
    // IP address associated with the network interface.
    networkAddress *string
    // Subnet mask.
    subnetMask *string
}
// NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic instantiates a new V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic and sets the default values.
func NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic()(*V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) {
    m := &V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["networkAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkAddress(val)
        }
        return nil
    }
    res["subnetMask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetMask(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the network interface.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) GetName()(*string) {
    return m.name
}
// GetNetworkAddress gets the networkAddress property value. IP address associated with the network interface.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetSubnetMask gets the subnetMask property value. Subnet mask.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) GetSubnetMask()(*string) {
    return m.subnetMask
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkAddress", m.GetNetworkAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetMask", m.GetSubnetMask())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the network interface.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) SetName(value *string)() {
    m.name = value
}
// SetNetworkAddress sets the networkAddress property value. IP address associated with the network interface.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetSubnetMask sets the subnetMask property value. Subnet mask.
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nic) SetSubnetMask(value *string)() {
    m.subnetMask = value
}
type V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_nicable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNetworkAddress()(*string)
    GetSubnetMask()(*string)
    SetName(value *string)()
    SetNetworkAddress(value *string)()
    SetSubnetMask(value *string)()
}
