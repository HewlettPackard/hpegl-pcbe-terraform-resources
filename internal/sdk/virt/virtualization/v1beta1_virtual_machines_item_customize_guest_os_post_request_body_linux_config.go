package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig configurations for linux operating system.
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // DNS configurations for the virtual machine.
    dnsSettings V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable
    // The domain name for the virtual machine.
    domainName *string
    // Hostname for the virtual machine.
    hostname V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable
    // Network interface configurations
    networkInterfaces V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDnsSettings gets the dnsSettings property value. DNS configurations for the virtual machine.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetDnsSettings()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable) {
    return m.dnsSettings
}
// GetDomainName gets the domainName property value. The domain name for the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dnsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsSettings(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable))
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable))
        }
        return nil
    }
    res["networkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInterfaces(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable))
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. Hostname for the virtual machine.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetHostname()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable) {
    return m.hostname
}
// GetNetworkInterfaces gets the networkInterfaces property value. Network interface configurations
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) GetNetworkInterfaces()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable) {
    return m.networkInterfaces
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dnsSettings", m.GetDnsSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkInterfaces", m.GetNetworkInterfaces())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDnsSettings sets the dnsSettings property value. DNS configurations for the virtual machine.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) SetDnsSettings(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable)() {
    m.dnsSettings = value
}
// SetDomainName sets the domainName property value. The domain name for the virtual machine.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHostname sets the hostname property value. Hostname for the virtual machine.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) SetHostname(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable)() {
    m.hostname = value
}
// SetNetworkInterfaces sets the networkInterfaces property value. Network interface configurations
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig) SetNetworkInterfaces(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable)() {
    m.networkInterfaces = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDnsSettings()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable)
    GetDomainName()(*string)
    GetHostname()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable)
    GetNetworkInterfaces()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable)
    SetDnsSettings(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_dnsSettingsable)()
    SetDomainName(value *string)()
    SetHostname(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_hostnameable)()
    SetNetworkInterfaces(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfacesable)()
}
