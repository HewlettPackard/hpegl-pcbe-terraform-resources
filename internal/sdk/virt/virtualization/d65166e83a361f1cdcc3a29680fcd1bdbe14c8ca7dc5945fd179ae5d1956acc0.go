package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP Address of the Network Interface.
    ipAddress *string
    // Prefix length of the IP Address.
    prefixLength *int64
    // State of the IP Address.
    state *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["prefixLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefixLength(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. IP Address of the Network Interface.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetPrefixLength gets the prefixLength property value. Prefix length of the IP Address.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) GetPrefixLength()(*int64) {
    return m.prefixLength
}
// GetState gets the state property value. State of the IP Address.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("prefixLength", m.GetPrefixLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIpAddress sets the ipAddress property value. IP Address of the Network Interface.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetPrefixLength sets the prefixLength property value. Prefix length of the IP Address.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) SetPrefixLength(value *int64)() {
    m.prefixLength = value
}
// SetState sets the state property value. State of the IP Address.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddresses) SetState(value *string)() {
    m.state = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfo_networkInterfacesInfo_ipAddressesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddress()(*string)
    GetPrefixLength()(*int64)
    GetState()(*string)
    SetIpAddress(value *string)()
    SetPrefixLength(value *int64)()
    SetState(value *string)()
}
