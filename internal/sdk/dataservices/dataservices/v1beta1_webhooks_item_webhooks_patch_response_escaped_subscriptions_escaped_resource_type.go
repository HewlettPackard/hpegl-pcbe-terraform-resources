package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the resource type that is subscribed to
    applicationResource *string
    // Name of the resource type
    name *string
}
// NewV1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType instantiates a new V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType and sets the default values.
func NewV1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType()(*V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) {
    m := &V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationResource gets the applicationResource property value. Identifies the resource type that is subscribed to
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) GetApplicationResource()(*string) {
    return m.applicationResource
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationResource(val)
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
// GetName gets the name property value. Name of the resource type
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationResource", m.GetApplicationResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationResource sets the applicationResource property value. Identifies the resource type that is subscribed to
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) SetApplicationResource(value *string)() {
    m.applicationResource = value
}
// SetName sets the name property value. Name of the resource type
func (m *V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceType) SetName(value *string)() {
    m.name = value
}
type V1beta1WebhooksItemWebhooksPatchResponse_subscriptions_resourceTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationResource()(*string)
    GetName()(*string)
    SetApplicationResource(value *string)()
    SetName(value *string)()
}
