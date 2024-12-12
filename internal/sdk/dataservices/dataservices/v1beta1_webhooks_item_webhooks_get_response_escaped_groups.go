package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksItemWebhooksGetResponse_groups reference to a Group resource
type V1beta1WebhooksItemWebhooksGetResponse_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Group's unique identifier
    id *string
    // The Group's name
    name *string
}
// NewV1beta1WebhooksItemWebhooksGetResponse_groups instantiates a new V1beta1WebhooksItemWebhooksGetResponse_groups and sets the default values.
func NewV1beta1WebhooksItemWebhooksGetResponse_groups()(*V1beta1WebhooksItemWebhooksGetResponse_groups) {
    m := &V1beta1WebhooksItemWebhooksGetResponse_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksItemWebhooksGetResponse_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksItemWebhooksGetResponse_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksItemWebhooksGetResponse_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The Group's unique identifier
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The Group's name
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The Group's unique identifier
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The Group's name
func (m *V1beta1WebhooksItemWebhooksGetResponse_groups) SetName(value *string)() {
    m.name = value
}
type V1beta1WebhooksItemWebhooksGetResponse_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    SetId(value *string)()
    SetName(value *string)()
}
