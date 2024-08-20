package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4 iPV4 configurations for a network adapter.
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of gateways for the IPV4 IP address. Only to be assigned type STATIC IP.
    gateways []string
    // IP Address for the network adapter. Only to be assigned type STATIC IP.
    ipAddress *string
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4 instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4 and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gateways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGateways(res)
        }
        return nil
    }
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
    return res
}
// GetGateways gets the gateways property value. List of gateways for the IPV4 IP address. Only to be assigned type STATIC IP.
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) GetGateways()([]string) {
    return m.gateways
}
// GetIpAddress gets the ipAddress property value. IP Address for the network adapter. Only to be assigned type STATIC IP.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) GetIpAddress()(*string) {
    return m.ipAddress
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGateways() != nil {
        err := writer.WriteCollectionOfStringValues("gateways", m.GetGateways())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGateways sets the gateways property value. List of gateways for the IPV4 IP address. Only to be assigned type STATIC IP.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) SetGateways(value []string)() {
    m.gateways = value
}
// SetIpAddress sets the ipAddress property value. IP Address for the network adapter. Only to be assigned type STATIC IP.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4) SetIpAddress(value *string)() {
    m.ipAddress = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGateways()([]string)
    GetIpAddress()(*string)
    SetGateways(value []string)()
    SetIpAddress(value *string)()
}
