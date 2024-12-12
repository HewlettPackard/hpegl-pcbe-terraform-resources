package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication an object containing the endpoint authentication attributes
type V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The bearerToken property
    bearerToken V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable
}
// NewV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication instantiates a new V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication and sets the default values.
func NewV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication()(*V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) {
    m := &V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBearerToken gets the bearerToken property value. The bearerToken property
// returns a V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable when successful
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) GetBearerToken()(V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable) {
    return m.bearerToken
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bearerToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBearerToken(val.(V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bearerToken", m.GetBearerToken())
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
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBearerToken sets the bearerToken property value. The bearerToken property
func (m *V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication) SetBearerToken(value V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable)() {
    m.bearerToken = value
}
type V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authenticationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBearerToken()(V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable)
    SetBearerToken(value V1beta1WebhooksPostResponse_destinations_cloudEventEndpoint_authentication_bearerTokenable)()
}
