package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo describes a storage system.
type V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A user-friendly name that identifies the storage system.
    displayName *string
    // UUID string uniquely identifying the storage system.
    id *string
    // Specify if the storage system is registered.
    managed *bool
    // Name of the storage system.
    name *string
    // The URI reference for this resource.
    resourceUri *string
    // Serial number of the storage system.
    serialNumber *string
    // Storage system provider name.
    vendorName *string
    // The wwn property
    wwn *string
}
// NewV1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo instantiates a new V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo()(*V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) {
    m := &V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the storage system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["managed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManaged(val)
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
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["vendorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorName(val)
        }
        return nil
    }
    res["wwn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWwn(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. UUID string uniquely identifying the storage system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetId()(*string) {
    return m.id
}
// GetManaged gets the managed property value. Specify if the storage system is registered.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetManaged()(*bool) {
    return m.managed
}
// GetName gets the name property value. Name of the storage system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSerialNumber gets the serialNumber property value. Serial number of the storage system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetVendorName gets the vendorName property value. Storage system provider name.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetVendorName()(*string) {
    return m.vendorName
}
// GetWwn gets the wwn property value. The wwn property
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) GetWwn()(*string) {
    return m.wwn
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("managed", m.GetManaged())
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
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendorName", m.GetVendorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wwn", m.GetWwn())
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
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the storage system.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. UUID string uniquely identifying the storage system.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetId(value *string)() {
    m.id = value
}
// SetManaged sets the managed property value. Specify if the storage system is registered.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetManaged(value *bool)() {
    m.managed = value
}
// SetName sets the name property value. Name of the storage system.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSerialNumber sets the serialNumber property value. Serial number of the storage system.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetVendorName sets the vendorName property value. Storage system provider name.
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetVendorName(value *string)() {
    m.vendorName = value
}
// SetWwn sets the wwn property value. The wwn property
func (m *V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfo) SetWwn(value *string)() {
    m.wwn = value
}
type V1beta1VirtualMachinesGetResponse_items_volumesInfo_storageSystemInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetManaged()(*bool)
    GetName()(*string)
    GetResourceUri()(*string)
    GetSerialNumber()(*string)
    GetVendorName()(*string)
    GetWwn()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetManaged(value *bool)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetSerialNumber(value *string)()
    SetVendorName(value *string)()
    SetWwn(value *string)()
}
