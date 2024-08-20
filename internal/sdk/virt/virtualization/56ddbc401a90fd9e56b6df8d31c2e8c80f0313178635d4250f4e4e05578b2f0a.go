package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IPV4 configurations for a network adapter.
    ipv4 V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ipv4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpv4(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able))
        }
        return nil
    }
    return res
}
// GetIpv4 gets the ipv4 property value. IPV4 configurations for a network adapter.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) GetIpv4()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able) {
    return m.ipv4
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ipv4", m.GetIpv4())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIpv4 sets the ipv4 property value. IPV4 configurations for a network adapter.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter) SetIpv4(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able)() {
    m.ipv4 = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpv4()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able)
    SetIpv4(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4able)()
}
