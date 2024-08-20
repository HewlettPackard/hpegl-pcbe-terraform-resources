package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo processor information.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The speed of the processor, in GHz.
    sustainedClockSpeedInGhz *float64
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sustainedClockSpeedInGhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSustainedClockSpeedInGhz(val)
        }
        return nil
    }
    return res
}
// GetSustainedClockSpeedInGhz gets the sustainedClockSpeedInGhz property value. The speed of the processor, in GHz.
// returns a *float64 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) GetSustainedClockSpeedInGhz()(*float64) {
    return m.sustainedClockSpeedInGhz
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("sustainedClockSpeedInGhz", m.GetSustainedClockSpeedInGhz())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSustainedClockSpeedInGhz sets the sustainedClockSpeedInGhz property value. The speed of the processor, in GHz.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfo) SetSustainedClockSpeedInGhz(value *float64)() {
    m.sustainedClockSpeedInGhz = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_processorInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSustainedClockSpeedInGhz()(*float64)
    SetSustainedClockSpeedInGhz(value *float64)()
}
