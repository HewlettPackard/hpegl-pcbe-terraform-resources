package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the property to be set to the virtual machine
    key *string
    // Value of the property to be set to the virtual machine
    value *string
}
// NewV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig instantiates a new V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig()(*V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) {
    m := &V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. Name of the property to be set to the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) GetKey()(*string) {
    return m.key
}
// GetValue gets the value property value. Value of the property to be set to the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Name of the property to be set to the virtual machine
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) SetKey(value *string)() {
    m.key = value
}
// SetValue sets the value property value. Value of the property to be set to the virtual machine
func (m *V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfig) SetValue(value *string)() {
    m.value = value
}
type V1beta1VirtualMachinesPostRequestBody_vmConfig_propertyConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetValue()(*string)
    SetKey(value *string)()
    SetValue(value *string)()
}
