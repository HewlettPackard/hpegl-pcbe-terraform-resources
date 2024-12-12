package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage the os disk image information.
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The operating system of the osDiskImage.
    operatingSystem *string
    // Size of the operating system of the osDiskImage.
    sizeInGb *int32
}
// NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage instantiates a new V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage and sets the default values.
func NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage()(*V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) {
    m := &V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["sizeInGb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInGb(val)
        }
        return nil
    }
    return res
}
// GetOperatingSystem gets the operatingSystem property value. The operating system of the osDiskImage.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetSizeInGb gets the sizeInGb property value. Size of the operating system of the osDiskImage.
// returns a *int32 when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) GetSizeInGb()(*int32) {
    return m.sizeInGb
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sizeInGb", m.GetSizeInGb())
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
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOperatingSystem sets the operatingSystem property value. The operating system of the osDiskImage.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetSizeInGb sets the sizeInGb property value. Size of the operating system of the osDiskImage.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImage) SetSizeInGb(value *int32)() {
    m.sizeInGb = value
}
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_osDiskImageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperatingSystem()(*string)
    GetSizeInGb()(*int32)
    SetOperatingSystem(value *string)()
    SetSizeInGb(value *int32)()
}
