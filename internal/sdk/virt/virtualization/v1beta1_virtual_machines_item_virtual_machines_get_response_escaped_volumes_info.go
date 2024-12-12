package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo information of volume or snapshot.
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A user-friendly name that identifies the volume.
    displayName *string
    // UUID string uniquely identifying the volume.
    id *string
    // Name of the volume.
    name *string
    // The URI reference for this resource.
    resourceUri *string
    // SCSI identifier of the volume or snapshot .
    scsiIdentifier *string
    // Size of the volume or snapshot in bytes.
    sizeInBytes *int64
    // Information of storage folder.
    storageFolderInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable
    // Describes a storage pool.
    storagePoolInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable
    // Describes a storage system.
    storageSystemInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable
    // Type of volume.
    typeEscaped *string
    // Describes a volume set.
    volumeSetInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the volume.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["scsiIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScsiIdentifier(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["storageFolderInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageFolderInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable))
        }
        return nil
    }
    res["storagePoolInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStoragePoolInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable))
        }
        return nil
    }
    res["storageSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSystemInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable))
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
    res["volumeSetInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeSetInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. UUID string uniquely identifying the volume.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the volume.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetScsiIdentifier gets the scsiIdentifier property value. SCSI identifier of the volume or snapshot .
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetScsiIdentifier()(*string) {
    return m.scsiIdentifier
}
// GetSizeInBytes gets the sizeInBytes property value. Size of the volume or snapshot in bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetSizeInBytes()(*int64) {
    return m.sizeInBytes
}
// GetStorageFolderInfo gets the storageFolderInfo property value. Information of storage folder.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetStorageFolderInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable) {
    return m.storageFolderInfo
}
// GetStoragePoolInfo gets the storagePoolInfo property value. Describes a storage pool.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetStoragePoolInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable) {
    return m.storagePoolInfo
}
// GetStorageSystemInfo gets the storageSystemInfo property value. Describes a storage system.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetStorageSystemInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable) {
    return m.storageSystemInfo
}
// GetTypeEscaped gets the type property value. Type of volume.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetVolumeSetInfo gets the volumeSetInfo property value. Describes a volume set.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) GetVolumeSetInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable) {
    return m.volumeSetInfo
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("scsiIdentifier", m.GetScsiIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageFolderInfo", m.GetStorageFolderInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storagePoolInfo", m.GetStoragePoolInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageSystemInfo", m.GetStorageSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("volumeSetInfo", m.GetVolumeSetInfo())
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the volume.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. UUID string uniquely identifying the volume.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the volume.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetScsiIdentifier sets the scsiIdentifier property value. SCSI identifier of the volume or snapshot .
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetScsiIdentifier(value *string)() {
    m.scsiIdentifier = value
}
// SetSizeInBytes sets the sizeInBytes property value. Size of the volume or snapshot in bytes.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
// SetStorageFolderInfo sets the storageFolderInfo property value. Information of storage folder.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetStorageFolderInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable)() {
    m.storageFolderInfo = value
}
// SetStoragePoolInfo sets the storagePoolInfo property value. Describes a storage pool.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetStoragePoolInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable)() {
    m.storagePoolInfo = value
}
// SetStorageSystemInfo sets the storageSystemInfo property value. Describes a storage system.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetStorageSystemInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable)() {
    m.storageSystemInfo = value
}
// SetTypeEscaped sets the type property value. Type of volume.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetVolumeSetInfo sets the volumeSetInfo property value. Describes a volume set.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo) SetVolumeSetInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable)() {
    m.volumeSetInfo = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetScsiIdentifier()(*string)
    GetSizeInBytes()(*int64)
    GetStorageFolderInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable)
    GetStoragePoolInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable)
    GetStorageSystemInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable)
    GetTypeEscaped()(*string)
    GetVolumeSetInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetScsiIdentifier(value *string)()
    SetSizeInBytes(value *int64)()
    SetStorageFolderInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageFolderInfoable)()
    SetStoragePoolInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storagePoolInfoable)()
    SetStorageSystemInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_storageSystemInfoable)()
    SetTypeEscaped(value *string)()
    SetVolumeSetInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfo_volumeSetInfoable)()
}
