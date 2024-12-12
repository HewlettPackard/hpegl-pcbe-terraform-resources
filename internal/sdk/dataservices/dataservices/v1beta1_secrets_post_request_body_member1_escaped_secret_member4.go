package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember4 certificate secret definition
type V1beta1SecretsPostRequestBodyMember1_secretMember4 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Certificate secret properties
    certificate V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember4 instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember4 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember4()(*V1beta1SecretsPostRequestBodyMember1_secretMember4) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember4{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember4FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember4FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember4(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificate gets the certificate property value. Certificate secret properties
// returns a V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) GetCertificate()(V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable) {
    return m.certificate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1_secretMember4_certificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val.(V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("certificate", m.GetCertificate())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificate sets the certificate property value. Certificate secret properties
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember4) SetCertificate(value V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable)() {
    m.certificate = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember4able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()(V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable)
    SetCertificate(value V1beta1SecretsPostRequestBodyMember1_secretMember4_certificateable)()
}
