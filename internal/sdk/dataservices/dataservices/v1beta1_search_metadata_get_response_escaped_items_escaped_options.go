package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SearchMetadataGetResponse_items_options struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Machine readable item identifier
    key *string
    // Human readable item name
    value *string
}
// NewV1beta1SearchMetadataGetResponse_items_options instantiates a new V1beta1SearchMetadataGetResponse_items_options and sets the default values.
func NewV1beta1SearchMetadataGetResponse_items_options()(*V1beta1SearchMetadataGetResponse_items_options) {
    m := &V1beta1SearchMetadataGetResponse_items_options{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SearchMetadataGetResponse_items_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchMetadataGetResponse_items_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchMetadataGetResponse_items_options(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SearchMetadataGetResponse_items_options) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SearchMetadataGetResponse_items_options) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetKey gets the key property value. Machine readable item identifier
// returns a *string when successful
func (m *V1beta1SearchMetadataGetResponse_items_options) GetKey()(*string) {
    return m.key
}
// GetValue gets the value property value. Human readable item name
// returns a *string when successful
func (m *V1beta1SearchMetadataGetResponse_items_options) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *V1beta1SearchMetadataGetResponse_items_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SearchMetadataGetResponse_items_options) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Machine readable item identifier
func (m *V1beta1SearchMetadataGetResponse_items_options) SetKey(value *string)() {
    m.key = value
}
// SetValue sets the value property value. Human readable item name
func (m *V1beta1SearchMetadataGetResponse_items_options) SetValue(value *string)() {
    m.value = value
}
type V1beta1SearchMetadataGetResponse_items_optionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetValue()(*string)
    SetKey(value *string)()
    SetValue(value *string)()
}
