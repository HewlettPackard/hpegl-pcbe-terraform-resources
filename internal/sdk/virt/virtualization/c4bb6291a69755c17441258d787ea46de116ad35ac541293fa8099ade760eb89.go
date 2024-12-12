package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings dNS configurations for the virtual machine.
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of DNS servers for the virtual network adapter.
    dnsServers []string
    // List of name resolutions for the virtual network adapter.
    searchDomain []string
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDnsServers gets the dnsServers property value. List of DNS servers for the virtual network adapter.
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) GetDnsServers()([]string) {
    return m.dnsServers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dnsServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDnsServers(res)
        }
        return nil
    }
    res["searchDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSearchDomain(res)
        }
        return nil
    }
    return res
}
// GetSearchDomain gets the searchDomain property value. List of name resolutions for the virtual network adapter.
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) GetSearchDomain()([]string) {
    return m.searchDomain
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDnsServers() != nil {
        err := writer.WriteCollectionOfStringValues("dnsServers", m.GetDnsServers())
        if err != nil {
            return err
        }
    }
    if m.GetSearchDomain() != nil {
        err := writer.WriteCollectionOfStringValues("searchDomain", m.GetSearchDomain())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDnsServers sets the dnsServers property value. List of DNS servers for the virtual network adapter.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) SetDnsServers(value []string)() {
    m.dnsServers = value
}
// SetSearchDomain sets the searchDomain property value. List of name resolutions for the virtual network adapter.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettings) SetSearchDomain(value []string)() {
    m.searchDomain = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDnsServers()([]string)
    GetSearchDomain()([]string)
    SetDnsServers(value []string)()
    SetSearchDomain(value []string)()
}
