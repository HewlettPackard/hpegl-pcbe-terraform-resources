package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An object containing the endpoint authentication attributes
    authentication V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable
    // The webhook endpoint address
    endpoint *string
}
// NewV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint instantiates a new V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint and sets the default values.
func NewV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint()(*V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) {
    m := &V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthentication gets the authentication property value. An object containing the endpoint authentication attributes
// returns a V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) GetAuthentication()(V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable) {
    return m.authentication
}
// GetEndpoint gets the endpoint property value. The webhook endpoint address
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) GetEndpoint()(*string) {
    return m.endpoint
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthentication(val.(V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable))
        }
        return nil
    }
    res["endpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpoint(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endpoint", m.GetEndpoint())
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
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthentication sets the authentication property value. An object containing the endpoint authentication attributes
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) SetAuthentication(value V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable)() {
    m.authentication = value
}
// SetEndpoint sets the endpoint property value. The webhook endpoint address
func (m *V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint) SetEndpoint(value *string)() {
    m.endpoint = value
}
type V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthentication()(V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable)
    GetEndpoint()(*string)
    SetAuthentication(value V1beta1WebhooksItemWebhooksGetResponse_destinations_cloudEventEndpoint_authenticationable)()
    SetEndpoint(value *string)()
}
