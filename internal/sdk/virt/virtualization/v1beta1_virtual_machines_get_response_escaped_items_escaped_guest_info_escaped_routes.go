package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesGetResponse_items_guestInfo_routes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Gateway address of the guest operating system on this virtual machine.
    gatewayAddress *string
    // Interface index of the guest operating system on this virtual machine.
    interfaceIndex *int64
    // Network of the guest operating system on this virtual machine.
    network *string
    // Prefix length of the guest operating system on this virtual machine.
    prefixLength *int64
}
// NewV1beta1VirtualMachinesGetResponse_items_guestInfo_routes instantiates a new V1beta1VirtualMachinesGetResponse_items_guestInfo_routes and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_guestInfo_routes()(*V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) {
    m := &V1beta1VirtualMachinesGetResponse_items_guestInfo_routes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_routesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_routesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_guestInfo_routes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gatewayAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGatewayAddress(val)
        }
        return nil
    }
    res["interfaceIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterfaceIndex(val)
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
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
    return res
}
// GetGatewayAddress gets the gatewayAddress property value. Gateway address of the guest operating system on this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetGatewayAddress()(*string) {
    return m.gatewayAddress
}
// GetInterfaceIndex gets the interfaceIndex property value. Interface index of the guest operating system on this virtual machine.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetInterfaceIndex()(*int64) {
    return m.interfaceIndex
}
// GetNetwork gets the network property value. Network of the guest operating system on this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetNetwork()(*string) {
    return m.network
}
// GetPrefixLength gets the prefixLength property value. Prefix length of the guest operating system on this virtual machine.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) GetPrefixLength()(*int64) {
    return m.prefixLength
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("gatewayAddress", m.GetGatewayAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("interfaceIndex", m.GetInterfaceIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGatewayAddress sets the gatewayAddress property value. Gateway address of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) SetGatewayAddress(value *string)() {
    m.gatewayAddress = value
}
// SetInterfaceIndex sets the interfaceIndex property value. Interface index of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) SetInterfaceIndex(value *int64)() {
    m.interfaceIndex = value
}
// SetNetwork sets the network property value. Network of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) SetNetwork(value *string)() {
    m.network = value
}
// SetPrefixLength sets the prefixLength property value. Prefix length of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_routes) SetPrefixLength(value *int64)() {
    m.prefixLength = value
}
type V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGatewayAddress()(*string)
    GetInterfaceIndex()(*int64)
    GetNetwork()(*string)
    GetPrefixLength()(*int64)
    SetGatewayAddress(value *string)()
    SetInterfaceIndex(value *int64)()
    SetNetwork(value *string)()
    SetPrefixLength(value *int64)()
}
