package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesGetResponse_items_virtualDisks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypervisor specific information.
    appInfo V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable
    // Last known size of the virtual disk File path of the virtual disk
    capacityInBytes *int64
    // Controller key of disk.
    controllerKey *int64
    // File path of the virtual disk
    filePath *string
    // UUID for the virtual disk.
    id *string
    // A unique identifier to select disk for modification.
    key *int64
    // Name of the virtual disk.
    name *string
    // The storageSystemInfo property
    storageSystemInfo []V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable
    // Unique identifier of the virtual disk as reported by the hypervisor.
    uid *string
}
// NewV1beta1VirtualMachinesGetResponse_items_virtualDisks instantiates a new V1beta1VirtualMachinesGetResponse_items_virtualDisks and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_virtualDisks()(*V1beta1VirtualMachinesGetResponse_items_virtualDisks) {
    m := &V1beta1VirtualMachinesGetResponse_items_virtualDisks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_virtualDisksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_virtualDisksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_virtualDisks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Hypervisor specific information.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetAppInfo()(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable) {
    return m.appInfo
}
// GetCapacityInBytes gets the capacityInBytes property value. Last known size of the virtual disk File path of the virtual disk
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetCapacityInBytes()(*int64) {
    return m.capacityInBytes
}
// GetControllerKey gets the controllerKey property value. Controller key of disk.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetControllerKey()(*int64) {
    return m.controllerKey
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable))
        }
        return nil
    }
    res["capacityInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityInBytes(val)
        }
        return nil
    }
    res["controllerKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControllerKey(val)
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
    res["storageSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable)
                }
            }
            m.SetStorageSystemInfo(res)
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    return res
}
// GetFilePath gets the filePath property value. File path of the virtual disk
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetFilePath()(*string) {
    return m.filePath
}
// GetId gets the id property value. UUID for the virtual disk.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetId()(*string) {
    return m.id
}
// GetKey gets the key property value. A unique identifier to select disk for modification.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetKey()(*int64) {
    return m.key
}
// GetName gets the name property value. Name of the virtual disk.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetName()(*string) {
    return m.name
}
// GetStorageSystemInfo gets the storageSystemInfo property value. The storageSystemInfo property
// returns a []V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetStorageSystemInfo()([]V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable) {
    return m.storageSystemInfo
}
// GetUid gets the uid property value. Unique identifier of the virtual disk as reported by the hypervisor.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) GetUid()(*string) {
    return m.uid
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("capacityInBytes", m.GetCapacityInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("controllerKey", m.GetControllerKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePath", m.GetFilePath())
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
        err := writer.WriteInt64Value("key", m.GetKey())
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
    if m.GetStorageSystemInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStorageSystemInfo()))
        for i, v := range m.GetStorageSystemInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("storageSystemInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uid", m.GetUid())
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
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Hypervisor specific information.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetAppInfo(value V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable)() {
    m.appInfo = value
}
// SetCapacityInBytes sets the capacityInBytes property value. Last known size of the virtual disk File path of the virtual disk
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetCapacityInBytes(value *int64)() {
    m.capacityInBytes = value
}
// SetControllerKey sets the controllerKey property value. Controller key of disk.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetControllerKey(value *int64)() {
    m.controllerKey = value
}
// SetFilePath sets the filePath property value. File path of the virtual disk
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetFilePath(value *string)() {
    m.filePath = value
}
// SetId sets the id property value. UUID for the virtual disk.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetId(value *string)() {
    m.id = value
}
// SetKey sets the key property value. A unique identifier to select disk for modification.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetKey(value *int64)() {
    m.key = value
}
// SetName sets the name property value. Name of the virtual disk.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetName(value *string)() {
    m.name = value
}
// SetStorageSystemInfo sets the storageSystemInfo property value. The storageSystemInfo property
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetStorageSystemInfo(value []V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable)() {
    m.storageSystemInfo = value
}
// SetUid sets the uid property value. Unique identifier of the virtual disk as reported by the hypervisor.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks) SetUid(value *string)() {
    m.uid = value
}
type V1beta1VirtualMachinesGetResponse_items_virtualDisksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable)
    GetCapacityInBytes()(*int64)
    GetControllerKey()(*int64)
    GetFilePath()(*string)
    GetId()(*string)
    GetKey()(*int64)
    GetName()(*string)
    GetStorageSystemInfo()([]V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable)
    GetUid()(*string)
    SetAppInfo(value V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfoable)()
    SetCapacityInBytes(value *int64)()
    SetControllerKey(value *int64)()
    SetFilePath(value *string)()
    SetId(value *string)()
    SetKey(value *int64)()
    SetName(value *string)()
    SetStorageSystemInfo(value []V1beta1VirtualMachinesGetResponse_items_virtualDisks_storageSystemInfoable)()
    SetUid(value *string)()
}
