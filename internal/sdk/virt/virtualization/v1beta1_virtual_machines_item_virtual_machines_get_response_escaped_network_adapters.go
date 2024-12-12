package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // MAC address of the network adapter.
    macAddress *string
    // Name of the network adapter.
    name *string
    // The networkDetails property
    networkDetails V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
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
    res["networkDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDetails(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable))
        }
        return nil
    }
    return res
}
// GetMacAddress gets the macAddress property value. MAC address of the network adapter.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) GetMacAddress()(*string) {
    return m.macAddress
}
// GetName gets the name property value. Name of the network adapter.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) GetName()(*string) {
    return m.name
}
// GetNetworkDetails gets the networkDetails property value. The networkDetails property
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) GetNetworkDetails()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable) {
    return m.networkDetails
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkDetails", m.GetNetworkDetails())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMacAddress sets the macAddress property value. MAC address of the network adapter.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetName sets the name property value. Name of the network adapter.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) SetName(value *string)() {
    m.name = value
}
// SetNetworkDetails sets the networkDetails property value. The networkDetails property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters) SetNetworkDetails(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable)() {
    m.networkDetails = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMacAddress()(*string)
    GetName()(*string)
    GetNetworkDetails()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable)
    SetMacAddress(value *string)()
    SetName(value *string)()
    SetNetworkDetails(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable)()
}
