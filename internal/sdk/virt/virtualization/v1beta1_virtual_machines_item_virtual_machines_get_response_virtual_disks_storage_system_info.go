package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *string
    // The managed property
    managed *bool
    // The model property
    model *string
    // The name property
    name *string
    // The serialNumber property
    serialNumber *string
    // The vendorName property
    vendorName *string
    // The wwns property
    wwns []string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetId()(*string) {
    return m.id
}
// GetManaged gets the managed property value. The managed property
// returns a *bool when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetManaged()(*bool) {
    return m.managed
}
// GetModel gets the model property value. The model property
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetName()(*string) {
    return m.name
}
// GetSerialNumber gets the serialNumber property value. The serialNumber property
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetVendorName gets the vendorName property value. The vendorName property
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetVendorName()(*string) {
    return m.vendorName
}
// GetWwns gets the wwns property value. The wwns property
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) GetWwns()([]string) {
    return m.wwns
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetId(value *string)() {
    m.id = value
}
// SetManaged sets the managed property value. The managed property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetManaged(value *bool)() {
    m.managed = value
}
// SetModel sets the model property value. The model property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. The name property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetName(value *string)() {
    m.name = value
}
// SetSerialNumber sets the serialNumber property value. The serialNumber property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetVendorName sets the vendorName property value. The vendorName property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetVendorName(value *string)() {
    m.vendorName = value
}
// SetWwns sets the wwns property value. The wwns property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfo) SetWwns(value []string)() {
    m.wwns = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisks_storageSystemInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetManaged()(*bool)
    GetModel()(*string)
    GetName()(*string)
    GetSerialNumber()(*string)
    GetVendorName()(*string)
    GetWwns()([]string)
    SetId(value *string)()
    SetManaged(value *bool)()
    SetModel(value *string)()
    SetName(value *string)()
    SetSerialNumber(value *string)()
    SetVendorName(value *string)()
    SetWwns(value []string)()
}
