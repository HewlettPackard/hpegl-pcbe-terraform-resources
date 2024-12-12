package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig hPE Cloud module configuration.
type V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *string
    // The instanceType property
    instanceType *string
}
// NewV1beta1SystemsItemServersGetResponse_items_cloudModuleConfig instantiates a new V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_cloudModuleConfig()(*V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) {
    m := &V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_cloudModuleConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_cloudModuleConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_cloudModuleConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["instanceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) GetId()(*string) {
    return m.id
}
// GetInstanceType gets the instanceType property value. The instanceType property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) GetInstanceType()(*string) {
    return m.instanceType
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceType", m.GetInstanceType())
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
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) SetId(value *string)() {
    m.id = value
}
// SetInstanceType sets the instanceType property value. The instanceType property
func (m *V1beta1SystemsItemServersGetResponse_items_cloudModuleConfig) SetInstanceType(value *string)() {
    m.instanceType = value
}
type V1beta1SystemsItemServersGetResponse_items_cloudModuleConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetInstanceType()(*string)
    SetId(value *string)()
    SetInstanceType(value *string)()
}
