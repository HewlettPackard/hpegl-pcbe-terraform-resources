package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties automatic Os Upgrade Properties.
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Is automatic OS upgrade supported.
    automaticOsUpgradeSupported *bool
}
// NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties instantiates a new V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties and sets the default values.
func NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties()(*V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) {
    m := &V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradePropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradePropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutomaticOsUpgradeSupported gets the automaticOsUpgradeSupported property value. Is automatic OS upgrade supported.
// returns a *bool when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) GetAutomaticOsUpgradeSupported()(*bool) {
    return m.automaticOsUpgradeSupported
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["automaticOsUpgradeSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticOsUpgradeSupported(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("automaticOsUpgradeSupported", m.GetAutomaticOsUpgradeSupported())
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
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutomaticOsUpgradeSupported sets the automaticOsUpgradeSupported property value. Is automatic OS upgrade supported.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradeProperties) SetAutomaticOsUpgradeSupported(value *bool)() {
    m.automaticOsUpgradeSupported = value
}
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_automaticOsUpgradePropertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutomaticOsUpgradeSupported()(*bool)
    SetAutomaticOsUpgradeSupported(value *bool)()
}
