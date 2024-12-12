package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksPostResponse_authorization an object containing an identity that is used to authorize published events
type V1beta1WebhooksPostResponse_authorization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientCredentials property
    clientCredentials V1beta1WebhooksPostResponse_authorization_clientCredentialsable
}
// NewV1beta1WebhooksPostResponse_authorization instantiates a new V1beta1WebhooksPostResponse_authorization and sets the default values.
func NewV1beta1WebhooksPostResponse_authorization()(*V1beta1WebhooksPostResponse_authorization) {
    m := &V1beta1WebhooksPostResponse_authorization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksPostResponse_authorizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksPostResponse_authorizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksPostResponse_authorization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksPostResponse_authorization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientCredentials gets the clientCredentials property value. The clientCredentials property
// returns a V1beta1WebhooksPostResponse_authorization_clientCredentialsable when successful
func (m *V1beta1WebhooksPostResponse_authorization) GetClientCredentials()(V1beta1WebhooksPostResponse_authorization_clientCredentialsable) {
    return m.clientCredentials
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksPostResponse_authorization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksPostResponse_authorization_clientCredentialsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientCredentials(val.(V1beta1WebhooksPostResponse_authorization_clientCredentialsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksPostResponse_authorization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("clientCredentials", m.GetClientCredentials())
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
func (m *V1beta1WebhooksPostResponse_authorization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientCredentials sets the clientCredentials property value. The clientCredentials property
func (m *V1beta1WebhooksPostResponse_authorization) SetClientCredentials(value V1beta1WebhooksPostResponse_authorization_clientCredentialsable)() {
    m.clientCredentials = value
}
type V1beta1WebhooksPostResponse_authorizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientCredentials()(V1beta1WebhooksPostResponse_authorization_clientCredentialsable)
    SetClientCredentials(value V1beta1WebhooksPostResponse_authorization_clientCredentialsable)()
}
