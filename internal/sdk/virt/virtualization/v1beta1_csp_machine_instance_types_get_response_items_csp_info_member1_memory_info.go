package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo the memory for the instance type.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The size of the memory, in MiB.
    sizeInMiB *int32
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sizeInMiB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInMiB(val)
        }
        return nil
    }
    return res
}
// GetSizeInMiB gets the sizeInMiB property value. The size of the memory, in MiB.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) GetSizeInMiB()(*int32) {
    return m.sizeInMiB
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("sizeInMiB", m.GetSizeInMiB())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSizeInMiB sets the sizeInMiB property value. The size of the memory, in MiB.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfo) SetSizeInMiB(value *int32)() {
    m.sizeInMiB = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_memoryInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSizeInMiB()(*int32)
    SetSizeInMiB(value *int32)()
}
