package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo information about the VMware tools installed in this virtual machine.
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Type of the VMware tool installed in this virtual machine.
    typeEscaped *string
    // Version of the VMware tool installed in this virtual machine.
    version *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. Type of the VMware tool installed in this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetVersion gets the version property value. Version of the VMware tool installed in this virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. Type of the VMware tool installed in this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetVersion sets the version property value. Version of the VMware tool installed in this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfo) SetVersion(value *string)() {
    m.version = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_toolsInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*string)
    GetVersion()(*string)
    SetTypeEscaped(value *string)()
    SetVersion(value *string)()
}
