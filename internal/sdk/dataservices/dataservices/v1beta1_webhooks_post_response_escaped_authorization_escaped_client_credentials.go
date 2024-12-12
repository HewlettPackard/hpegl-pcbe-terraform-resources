package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksPostResponse_authorization_clientCredentials struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reference to a Secret resource
    secret V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable
}
// NewV1beta1WebhooksPostResponse_authorization_clientCredentials instantiates a new V1beta1WebhooksPostResponse_authorization_clientCredentials and sets the default values.
func NewV1beta1WebhooksPostResponse_authorization_clientCredentials()(*V1beta1WebhooksPostResponse_authorization_clientCredentials) {
    m := &V1beta1WebhooksPostResponse_authorization_clientCredentials{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksPostResponse_authorization_clientCredentialsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksPostResponse_authorization_clientCredentialsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksPostResponse_authorization_clientCredentials(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksPostResponse_authorization_clientCredentials_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable))
        }
        return nil
    }
    return res
}
// GetSecret gets the secret property value. Reference to a Secret resource
// returns a V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable when successful
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) GetSecret()(V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable) {
    return m.secret
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecret sets the secret property value. Reference to a Secret resource
func (m *V1beta1WebhooksPostResponse_authorization_clientCredentials) SetSecret(value V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable)() {
    m.secret = value
}
type V1beta1WebhooksPostResponse_authorization_clientCredentialsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecret()(V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable)
    SetSecret(value V1beta1WebhooksPostResponse_authorization_clientCredentials_secretable)()
}
