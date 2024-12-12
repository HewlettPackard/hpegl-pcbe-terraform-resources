package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Configurations for a disk
    diskConfig V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDiskConfig gets the diskConfig property value. Configurations for a disk
// returns a V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) GetDiskConfig()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable) {
    return m.diskConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["diskConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskConfig(val.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("diskConfig", m.GetDiskConfig())
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDiskConfig sets the diskConfig property value. Configurations for a disk
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks) SetDiskConfig(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable)() {
    m.diskConfig = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiskConfig()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable)
    SetDiskConfig(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisks_diskConfigable)()
}
