package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers provides Ahci Disk Controller details of the virtual machine.
type V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Bus number of Controller.
    busNumber *int64
    // Controller key of disk.
    controllerKey *int64
    // All the device associated with this controller.
    device []int64
    // A unique identifier to select disk for modification.
    key *int64
    // Label of Controller.
    label *string
    // Summary of Controller.
    summary *string
    // Unit Number of Controller.
    unitNumber *int64
}
// NewV1beta1VirtualMachinesGetResponse_items_virtualAhciControllers instantiates a new V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_virtualAhciControllers()(*V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) {
    m := &V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_virtualAhciControllersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_virtualAhciControllersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_virtualAhciControllers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusNumber gets the busNumber property value. Bus number of Controller.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetBusNumber()(*int64) {
    return m.busNumber
}
// GetControllerKey gets the controllerKey property value. Controller key of disk.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetControllerKey()(*int64) {
    return m.controllerKey
}
// GetDevice gets the device property value. All the device associated with this controller.
// returns a []int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetDevice()([]int64) {
    return m.device
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["busNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusNumber(val)
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
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int64))
                }
            }
            m.SetDevice(res)
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
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    res["unitNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitNumber(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. A unique identifier to select disk for modification.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetKey()(*int64) {
    return m.key
}
// GetLabel gets the label property value. Label of Controller.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetLabel()(*string) {
    return m.label
}
// GetSummary gets the summary property value. Summary of Controller.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetSummary()(*string) {
    return m.summary
}
// GetUnitNumber gets the unitNumber property value. Unit Number of Controller.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) GetUnitNumber()(*int64) {
    return m.unitNumber
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("busNumber", m.GetBusNumber())
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
    if m.GetDevice() != nil {
        err := writer.WriteCollectionOfInt64Values("device", m.GetDevice())
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
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("unitNumber", m.GetUnitNumber())
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
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusNumber sets the busNumber property value. Bus number of Controller.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetBusNumber(value *int64)() {
    m.busNumber = value
}
// SetControllerKey sets the controllerKey property value. Controller key of disk.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetControllerKey(value *int64)() {
    m.controllerKey = value
}
// SetDevice sets the device property value. All the device associated with this controller.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetDevice(value []int64)() {
    m.device = value
}
// SetKey sets the key property value. A unique identifier to select disk for modification.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetKey(value *int64)() {
    m.key = value
}
// SetLabel sets the label property value. Label of Controller.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetLabel(value *string)() {
    m.label = value
}
// SetSummary sets the summary property value. Summary of Controller.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetSummary(value *string)() {
    m.summary = value
}
// SetUnitNumber sets the unitNumber property value. Unit Number of Controller.
func (m *V1beta1VirtualMachinesGetResponse_items_virtualAhciControllers) SetUnitNumber(value *int64)() {
    m.unitNumber = value
}
type V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusNumber()(*int64)
    GetControllerKey()(*int64)
    GetDevice()([]int64)
    GetKey()(*int64)
    GetLabel()(*string)
    GetSummary()(*string)
    GetUnitNumber()(*int64)
    SetBusNumber(value *int64)()
    SetControllerKey(value *int64)()
    SetDevice(value []int64)()
    SetKey(value *int64)()
    SetLabel(value *string)()
    SetSummary(value *string)()
    SetUnitNumber(value *int64)()
}
