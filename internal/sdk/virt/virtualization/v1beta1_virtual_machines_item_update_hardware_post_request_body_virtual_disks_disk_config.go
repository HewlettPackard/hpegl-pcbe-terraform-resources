package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig configurations for a disk
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Size of the disk in mebibytes. Mandatory parameter for ADD and EDIT operation.
    capacityInMb *int32
    // The UUID of the disk. Mandatory parameter for EDIT and DELETE operation.
    id *string
    // Specifies if the backing files of the disk need to be retained after disk deletion. Mandatory parameter for DELETE operation.
    retainFiles *bool
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapacityInMb gets the capacityInMb property value. Size of the disk in mebibytes. Mandatory parameter for ADD and EDIT operation.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) GetCapacityInMb()(*int32) {
    return m.capacityInMb
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capacityInMb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityInMb(val)
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
    res["retainFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainFiles(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The UUID of the disk. Mandatory parameter for EDIT and DELETE operation.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) GetId()(*string) {
    return m.id
}
// GetRetainFiles gets the retainFiles property value. Specifies if the backing files of the disk need to be retained after disk deletion. Mandatory parameter for DELETE operation.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) GetRetainFiles()(*bool) {
    return m.retainFiles
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("capacityInMb", m.GetCapacityInMb())
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
        err := writer.WriteBoolValue("retainFiles", m.GetRetainFiles())
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapacityInMb sets the capacityInMb property value. Size of the disk in mebibytes. Mandatory parameter for ADD and EDIT operation.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) SetCapacityInMb(value *int32)() {
    m.capacityInMb = value
}
// SetId sets the id property value. The UUID of the disk. Mandatory parameter for EDIT and DELETE operation.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) SetId(value *string)() {
    m.id = value
}
// SetRetainFiles sets the retainFiles property value. Specifies if the backing files of the disk need to be retained after disk deletion. Mandatory parameter for DELETE operation.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfig) SetRetainFiles(value *bool)() {
    m.retainFiles = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacityInMb()(*int32)
    GetId()(*string)
    GetRetainFiles()(*bool)
    SetCapacityInMb(value *int32)()
    SetId(value *string)()
    SetRetainFiles(value *bool)()
}
