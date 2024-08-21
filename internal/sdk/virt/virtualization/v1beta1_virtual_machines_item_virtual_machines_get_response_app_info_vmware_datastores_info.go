package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A user-friendly name that identifies the datastore.
    displayName *string
    // UUID string uniquely identifying the datastore
    id *string
    // Name of the datastore as reported by the hypervisor manager.
    name *string
    // The URI reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the datastore.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
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
    return res
}
// GetId gets the id property value. UUID string uniquely identifying the datastore
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the datastore as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the datastore.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. UUID string uniquely identifying the datastore
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the datastore as reported by the hypervisor manager.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfo_vmware_datastoresInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}