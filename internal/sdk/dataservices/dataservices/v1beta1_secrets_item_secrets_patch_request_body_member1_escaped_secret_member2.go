package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2 basic Access Authentication secret redefinition
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Basic Access Authentication secret properties
    basicAuth V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2 instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2 and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBasicAuth gets the basicAuth property value. Basic Access Authentication secret properties
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) GetBasicAuth()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable) {
    return m.basicAuth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["basicAuth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasicAuth(val.(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("basicAuth", m.GetBasicAuth())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBasicAuth sets the basicAuth property value. Basic Access Authentication secret properties
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2) SetBasicAuth(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable)() {
    m.basicAuth = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBasicAuth()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable)
    SetBasicAuth(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember2_basicAuthable)()
}
