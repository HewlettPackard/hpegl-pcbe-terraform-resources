package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SettingsItemSettingsPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // New value of the Value setting
    currentValue *string
}
// NewV1beta1SettingsItemSettingsPatchRequestBody instantiates a new V1beta1SettingsItemSettingsPatchRequestBody and sets the default values.
func NewV1beta1SettingsItemSettingsPatchRequestBody()(*V1beta1SettingsItemSettingsPatchRequestBody) {
    m := &V1beta1SettingsItemSettingsPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SettingsItemSettingsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SettingsItemSettingsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SettingsItemSettingsPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SettingsItemSettingsPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentValue gets the currentValue property value. New value of the Value setting
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchRequestBody) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SettingsItemSettingsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SettingsItemSettingsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentValue", m.GetCurrentValue())
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
func (m *V1beta1SettingsItemSettingsPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentValue sets the currentValue property value. New value of the Value setting
func (m *V1beta1SettingsItemSettingsPatchRequestBody) SetCurrentValue(value *string)() {
    m.currentValue = value
}
type V1beta1SettingsItemSettingsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentValue()(*string)
    SetCurrentValue(value *string)()
}
