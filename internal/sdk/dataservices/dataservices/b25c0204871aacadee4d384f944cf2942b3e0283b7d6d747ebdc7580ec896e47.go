package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken jSON Web Token secret properties
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // JSON Web Token value
    token *string
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetToken gets the token property value. JSON Web Token value
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetToken sets the token property value. JSON Web Token value
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebToken) SetToken(value *string)() {
    m.token = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember6_jsonWebTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetToken()(*string)
    SetToken(value *string)()
}
