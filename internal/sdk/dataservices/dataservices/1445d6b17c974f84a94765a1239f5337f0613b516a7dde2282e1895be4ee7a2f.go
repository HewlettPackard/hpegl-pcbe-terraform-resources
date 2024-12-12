package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reference to a Secret resource
    secret V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable
}
// NewV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken instantiates a new V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken and sets the default values.
func NewV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken()(*V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) {
    m := &V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable))
        }
        return nil
    }
    return res
}
// GetSecret gets the secret property value. Reference to a Secret resource
// returns a V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable when successful
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) GetSecret()(V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable) {
    return m.secret
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("secret", m.GetSecret())
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
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecret sets the secret property value. Reference to a Secret resource
func (m *V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken) SetSecret(value V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable)() {
    m.secret = value
}
type V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecret()(V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable)
    SetSecret(value V1beta1WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_bearerToken_secretable)()
}
