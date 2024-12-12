package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Type of the VM disk.
    vmDiskType *string
}
// NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed instantiates a new V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed and sets the default values.
func NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed()(*V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) {
    m := &V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["vmDiskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmDiskType(val)
        }
        return nil
    }
    return res
}
// GetVmDiskType gets the vmDiskType property value. Type of the VM disk.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) GetVmDiskType()(*string) {
    return m.vmDiskType
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("vmDiskType", m.GetVmDiskType())
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
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVmDiskType sets the vmDiskType property value. Type of the VM disk.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowed) SetVmDiskType(value *string)() {
    m.vmDiskType = value
}
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_disallowedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVmDiskType()(*string)
    SetVmDiskType(value *string)()
}
