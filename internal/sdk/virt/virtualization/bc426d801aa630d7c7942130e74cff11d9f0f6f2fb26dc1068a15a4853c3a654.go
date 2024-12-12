package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware vMware specific app info.
type V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Information of the datastore where the virtual disk is residing.
    datastoreInfo V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable
    // True if the disk UUID is enabled for the virtual machine.
    diskUuidEnabled *bool
}
// NewV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware instantiates a new V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware()(*V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) {
    m := &V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmwareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmwareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDatastoreInfo gets the datastoreInfo property value. Information of the datastore where the virtual disk is residing.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) GetDatastoreInfo()(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable) {
    return m.datastoreInfo
}
// GetDiskUuidEnabled gets the diskUuidEnabled property value. True if the disk UUID is enabled for the virtual machine.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) GetDiskUuidEnabled()(*bool) {
    return m.diskUuidEnabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datastoreInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatastoreInfo(val.(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable))
        }
        return nil
    }
    res["diskUuidEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskUuidEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("datastoreInfo", m.GetDatastoreInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("diskUuidEnabled", m.GetDiskUuidEnabled())
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
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDatastoreInfo sets the datastoreInfo property value. Information of the datastore where the virtual disk is residing.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) SetDatastoreInfo(value V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable)() {
    m.datastoreInfo = value
}
// SetDiskUuidEnabled sets the diskUuidEnabled property value. True if the disk UUID is enabled for the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware) SetDiskUuidEnabled(value *bool)() {
    m.diskUuidEnabled = value
}
type V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmwareable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatastoreInfo()(V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable)
    GetDiskUuidEnabled()(*bool)
    SetDatastoreInfo(value V1beta1VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_datastoreInfoable)()
    SetDiskUuidEnabled(value *bool)()
}
