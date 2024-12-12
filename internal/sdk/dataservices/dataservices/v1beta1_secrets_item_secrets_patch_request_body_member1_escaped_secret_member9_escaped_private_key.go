package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey private Key secret properties
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Private Key (BASE64-encoded)
    privateKey *string
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["privateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateKey(val)
        }
        return nil
    }
    return res
}
// GetPrivateKey gets the privateKey property value. Private Key (BASE64-encoded)
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) GetPrivateKey()(*string) {
    return m.privateKey
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("privateKey", m.GetPrivateKey())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateKey sets the privateKey property value. Private Key (BASE64-encoded)
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKey) SetPrivateKey(value *string)() {
    m.privateKey = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember9_privateKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateKey()(*string)
    SetPrivateKey(value *string)()
}
