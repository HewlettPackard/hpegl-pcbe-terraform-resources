package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksMetadataGetResponse_items_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Machine readable item identifier
    key *string
    // Additional attributes for a Resource metadata item
    resource V1beta1WebhooksMetadataGetResponse_items_items_resourceable
    // Human readable item name
    value *string
}
// NewV1beta1WebhooksMetadataGetResponse_items_items instantiates a new V1beta1WebhooksMetadataGetResponse_items_items and sets the default values.
func NewV1beta1WebhooksMetadataGetResponse_items_items()(*V1beta1WebhooksMetadataGetResponse_items_items) {
    m := &V1beta1WebhooksMetadataGetResponse_items_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksMetadataGetResponse_items_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksMetadataGetResponse_items_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksMetadataGetResponse_items_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksMetadataGetResponse_items_items_resourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(V1beta1WebhooksMetadataGetResponse_items_items_resourceable))
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
func (m *V1beta1WebhooksMetadataGetResponse_items_items) GetKey()(*string) {
    return m.key
}
// GetResource gets the resource property value. Additional attributes for a Resource metadata item
// returns a V1beta1WebhooksMetadataGetResponse_items_items_resourceable when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items) GetResource()(V1beta1WebhooksMetadataGetResponse_items_items_resourceable) {
    return m.resource
}
// GetValue gets the value property value. Human readable item name
// returns a *string when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksMetadataGetResponse_items_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
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
func (m *V1beta1WebhooksMetadataGetResponse_items_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Machine readable item identifier
func (m *V1beta1WebhooksMetadataGetResponse_items_items) SetKey(value *string)() {
    m.key = value
}
// SetResource sets the resource property value. Additional attributes for a Resource metadata item
func (m *V1beta1WebhooksMetadataGetResponse_items_items) SetResource(value V1beta1WebhooksMetadataGetResponse_items_items_resourceable)() {
    m.resource = value
}
// SetValue sets the value property value. Human readable item name
func (m *V1beta1WebhooksMetadataGetResponse_items_items) SetValue(value *string)() {
    m.value = value
}
type V1beta1WebhooksMetadataGetResponse_items_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetResource()(V1beta1WebhooksMetadataGetResponse_items_items_resourceable)
    GetValue()(*string)
    SetKey(value *string)()
    SetResource(value V1beta1WebhooksMetadataGetResponse_items_items_resourceable)()
    SetValue(value *string)()
}
