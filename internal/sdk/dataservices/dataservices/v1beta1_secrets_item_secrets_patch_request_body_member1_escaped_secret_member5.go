package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5 generic secret redefinition
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Generic secret properties
    generic V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5 instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5 and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["generic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneric(val.(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable))
        }
        return nil
    }
    return res
}
// GetGeneric gets the generic property value. Generic secret properties
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) GetGeneric()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable) {
    return m.generic
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("generic", m.GetGeneric())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGeneric sets the generic property value. Generic secret properties
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5) SetGeneric(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable)() {
    m.generic = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGeneric()(V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable)
    SetGeneric(value V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember5_genericable)()
}