package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth bearer Token Authentication secret properties
type V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Token (e.g. JWT)
    token *string
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth()(*V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetToken gets the token property value. Token (e.g. JWT)
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetToken sets the token property value. Token (e.g. JWT)
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuth) SetToken(value *string)() {
    m.token = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetToken()(*string)
    SetToken(value *string)()
}
