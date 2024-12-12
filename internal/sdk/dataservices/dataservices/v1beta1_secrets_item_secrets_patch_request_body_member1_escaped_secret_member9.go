package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9 private Key secret redefinition
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Private Key secret properties
    privateKey V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9 instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9 and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["privateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateKey(val.(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable))
        }
        return nil
    }
    return res
}
// GetPrivateKey gets the privateKey property value. Private Key secret properties
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) GetPrivateKey()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable) {
    return m.privateKey
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("privateKey", m.GetPrivateKey())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateKey sets the privateKey property value. Private Key secret properties
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9) SetPrivateKey(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable)() {
    m.privateKey = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateKey()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable)
    SetPrivateKey(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable)()
}
