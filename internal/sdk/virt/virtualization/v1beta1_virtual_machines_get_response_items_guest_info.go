package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_guestInfo information of this guest OS running on the virtual machine.
type V1beta1VirtualMachinesGetResponse_items_guestInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Build version of the guest operating system on this virtual machine.
    buildVersion *string
    // Name of the guest operating system on this virtual machine.
    name *string
    // Network Interfaces Information of the guest operating system on this virtual machine.
    networkInterfacesInfo []V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable
    // Release version of the guest operating system on this virtual machine.
    releaseVersion *string
    // Gateway Information of the guest operating system on this virtual machine.
    routes []V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable
    // Dns settings of the guest operating system on this virtual machine.
    vmDnsSetting V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable
}
// NewV1beta1VirtualMachinesGetResponse_items_guestInfo instantiates a new V1beta1VirtualMachinesGetResponse_items_guestInfo and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_guestInfo()(*V1beta1VirtualMachinesGetResponse_items_guestInfo) {
    m := &V1beta1VirtualMachinesGetResponse_items_guestInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_guestInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_guestInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_guestInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuildVersion gets the buildVersion property value. Build version of the guest operating system on this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetBuildVersion()(*string) {
    return m.buildVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildVersion(val)
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
    res["networkInterfacesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable)
                }
            }
            m.SetNetworkInterfacesInfo(res)
        }
        return nil
    }
    res["releaseVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseVersion(val)
        }
        return nil
    }
    res["routes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_routesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable)
                }
            }
            m.SetRoutes(res)
        }
        return nil
    }
    res["vmDnsSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmDnsSetting(val.(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the guest operating system on this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetName()(*string) {
    return m.name
}
// GetNetworkInterfacesInfo gets the networkInterfacesInfo property value. Network Interfaces Information of the guest operating system on this virtual machine.
// returns a []V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetNetworkInterfacesInfo()([]V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable) {
    return m.networkInterfacesInfo
}
// GetReleaseVersion gets the releaseVersion property value. Release version of the guest operating system on this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetReleaseVersion()(*string) {
    return m.releaseVersion
}
// GetRoutes gets the routes property value. Gateway Information of the guest operating system on this virtual machine.
// returns a []V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetRoutes()([]V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable) {
    return m.routes
}
// GetVmDnsSetting gets the vmDnsSetting property value. Dns settings of the guest operating system on this virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) GetVmDnsSetting()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable) {
    return m.vmDnsSetting
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("buildVersion", m.GetBuildVersion())
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
    if m.GetNetworkInterfacesInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkInterfacesInfo()))
        for i, v := range m.GetNetworkInterfacesInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkInterfacesInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("releaseVersion", m.GetReleaseVersion())
        if err != nil {
            return err
        }
    }
    if m.GetRoutes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoutes()))
        for i, v := range m.GetRoutes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("routes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vmDnsSetting", m.GetVmDnsSetting())
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
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuildVersion sets the buildVersion property value. Build version of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetBuildVersion(value *string)() {
    m.buildVersion = value
}
// SetName sets the name property value. Name of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetName(value *string)() {
    m.name = value
}
// SetNetworkInterfacesInfo sets the networkInterfacesInfo property value. Network Interfaces Information of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetNetworkInterfacesInfo(value []V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable)() {
    m.networkInterfacesInfo = value
}
// SetReleaseVersion sets the releaseVersion property value. Release version of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetReleaseVersion(value *string)() {
    m.releaseVersion = value
}
// SetRoutes sets the routes property value. Gateway Information of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetRoutes(value []V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable)() {
    m.routes = value
}
// SetVmDnsSetting sets the vmDnsSetting property value. Dns settings of the guest operating system on this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo) SetVmDnsSetting(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable)() {
    m.vmDnsSetting = value
}
type V1beta1VirtualMachinesGetResponse_items_guestInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuildVersion()(*string)
    GetName()(*string)
    GetNetworkInterfacesInfo()([]V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable)
    GetReleaseVersion()(*string)
    GetRoutes()([]V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable)
    GetVmDnsSetting()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable)
    SetBuildVersion(value *string)()
    SetName(value *string)()
    SetNetworkInterfacesInfo(value []V1beta1VirtualMachinesGetResponse_items_guestInfo_networkInterfacesInfoable)()
    SetReleaseVersion(value *string)()
    SetRoutes(value []V1beta1VirtualMachinesGetResponse_items_guestInfo_routesable)()
    SetVmDnsSetting(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable)()
}
