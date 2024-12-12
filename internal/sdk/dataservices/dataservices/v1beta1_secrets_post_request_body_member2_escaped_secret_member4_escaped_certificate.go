package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate certificate secret properties
type V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Distinguished Encoding Rules (DER) (BASE64-encoded)
    der *string
    // Privacy-Enhanced Mail (PEM) (BASE64-encoded)
    pem *string
}
// NewV1beta1SecretsPostRequestBodyMember2_secretMember4_certificate instantiates a new V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember2_secretMember4_certificate()(*V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) {
    m := &V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember2_secretMember4_certificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember2_secretMember4_certificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember2_secretMember4_certificate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDer gets the der property value. Distinguished Encoding Rules (DER) (BASE64-encoded)
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) GetDer()(*string) {
    return m.der
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["der"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDer(val)
        }
        return nil
    }
    res["pem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPem(val)
        }
        return nil
    }
    return res
}
// GetPem gets the pem property value. Privacy-Enhanced Mail (PEM) (BASE64-encoded)
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) GetPem()(*string) {
    return m.pem
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("der", m.GetDer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pem", m.GetPem())
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
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDer sets the der property value. Distinguished Encoding Rules (DER) (BASE64-encoded)
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) SetDer(value *string)() {
    m.der = value
}
// SetPem sets the pem property value. Privacy-Enhanced Mail (PEM) (BASE64-encoded)
func (m *V1beta1SecretsPostRequestBodyMember2_secretMember4_certificate) SetPem(value *string)() {
    m.pem = value
}
type V1beta1SecretsPostRequestBodyMember2_secretMember4_certificateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDer()(*string)
    GetPem()(*string)
    SetDer(value *string)()
    SetPem(value *string)()
}
