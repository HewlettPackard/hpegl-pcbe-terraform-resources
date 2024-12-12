package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8 password secret redefinition
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Password secret properties
    password V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8 instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8 and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val.(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable))
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. Password secret properties
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) GetPassword()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable) {
    return m.password
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("password", m.GetPassword())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPassword sets the password property value. Password secret properties
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8) SetPassword(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable)() {
    m.password = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable)
    SetPassword(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember8_passwordable)()
}
