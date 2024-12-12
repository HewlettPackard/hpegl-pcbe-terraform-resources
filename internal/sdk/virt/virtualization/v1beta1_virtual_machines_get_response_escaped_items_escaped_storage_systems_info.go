package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // id of the system
    id *string
    // Whether the storage system is managed or not
    managed *bool
    // model of storage system
    model *string
    // name of the storage system
    name *string
    // Serial number of the storage system.
    serialNumber *string
    // wwn of the system
    systemWwn *string
    // Storage system provider name
    vendorName *string
    // wwn of the volumes
    wwns []string
}
// NewV1beta1VirtualMachinesGetResponse_items_storageSystemsInfo instantiates a new V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_storageSystemsInfo()(*V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) {
    m := &V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_storageSystemsInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_storageSystemsInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_storageSystemsInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
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
    res["systemWwn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemWwn(val)
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
    res["wwns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetWwns(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. id of the system
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetId()(*string) {
    return m.id
}
// GetManaged gets the managed property value. Whether the storage system is managed or not
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetManaged()(*bool) {
    return m.managed
}
// GetModel gets the model property value. model of storage system
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. name of the storage system
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetName()(*string) {
    return m.name
}
// GetSerialNumber gets the serialNumber property value. Serial number of the storage system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSystemWwn gets the systemWwn property value. wwn of the system
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetSystemWwn()(*string) {
    return m.systemWwn
}
// GetVendorName gets the vendorName property value. Storage system provider name
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetVendorName()(*string) {
    return m.vendorName
}
// GetWwns gets the wwns property value. wwn of the volumes
// returns a []string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) GetWwns()([]string) {
    return m.wwns
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemWwn", m.GetSystemWwn())
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
    if m.GetWwns() != nil {
        err := writer.WriteCollectionOfStringValues("wwns", m.GetWwns())
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
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. id of the system
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetId(value *string)() {
    m.id = value
}
// SetManaged sets the managed property value. Whether the storage system is managed or not
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetManaged(value *bool)() {
    m.managed = value
}
// SetModel sets the model property value. model of storage system
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. name of the storage system
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetName(value *string)() {
    m.name = value
}
// SetSerialNumber sets the serialNumber property value. Serial number of the storage system.
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSystemWwn sets the systemWwn property value. wwn of the system
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetSystemWwn(value *string)() {
    m.systemWwn = value
}
// SetVendorName sets the vendorName property value. Storage system provider name
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetVendorName(value *string)() {
    m.vendorName = value
}
// SetWwns sets the wwns property value. wwn of the volumes
func (m *V1beta1VirtualMachinesGetResponse_items_storageSystemsInfo) SetWwns(value []string)() {
    m.wwns = value
}
type V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetManaged()(*bool)
    GetModel()(*string)
    GetName()(*string)
    GetSerialNumber()(*string)
    GetSystemWwn()(*string)
    GetVendorName()(*string)
    GetWwns()([]string)
    SetId(value *string)()
    SetManaged(value *bool)()
    SetModel(value *string)()
    SetName(value *string)()
    SetSerialNumber(value *string)()
    SetSystemWwn(value *string)()
    SetVendorName(value *string)()
    SetWwns(value []string)()
}
