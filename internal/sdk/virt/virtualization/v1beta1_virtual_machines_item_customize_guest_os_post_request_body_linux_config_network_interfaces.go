package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces network interface configurations
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of network adapter configuration.
    networkAdapter []V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["networkAdapter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable)
                }
            }
            m.SetNetworkAdapter(res)
        }
        return nil
    }
    return res
}
// GetNetworkAdapter gets the networkAdapter property value. List of network adapter configuration.
// returns a []V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) GetNetworkAdapter()([]V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable) {
    return m.networkAdapter
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNetworkAdapter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkAdapter()))
        for i, v := range m.GetNetworkAdapter() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkAdapter", cast)
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetworkAdapter sets the networkAdapter property value. List of network adapter configuration.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces) SetNetworkAdapter(value []V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable)() {
    m.networkAdapter = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkAdapter()([]V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable)
    SetNetworkAdapter(value []V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapterable)()
}
