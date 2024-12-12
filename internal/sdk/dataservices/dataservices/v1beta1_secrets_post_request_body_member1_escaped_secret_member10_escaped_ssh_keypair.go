package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair sSH Key Pair secret properties
type V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Private Key (BASE64-encoded)
    privateKey *string
    // Public Key (BASE64-encoded)
    publicKey *string
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair()(*V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["publicKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKey(val)
        }
        return nil
    }
    return res
}
// GetPrivateKey gets the privateKey property value. Private Key (BASE64-encoded)
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) GetPrivateKey()(*string) {
    return m.privateKey
}
// GetPublicKey gets the publicKey property value. Public Key (BASE64-encoded)
// returns a *string when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) GetPublicKey()(*string) {
    return m.publicKey
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("privateKey", m.GetPrivateKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publicKey", m.GetPublicKey())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateKey sets the privateKey property value. Private Key (BASE64-encoded)
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) SetPrivateKey(value *string)() {
    m.privateKey = value
}
// SetPublicKey sets the publicKey property value. Public Key (BASE64-encoded)
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypair) SetPublicKey(value *string)() {
    m.publicKey = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember10_sshKeypairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateKey()(*string)
    GetPublicKey()(*string)
    SetPrivateKey(value *string)()
    SetPublicKey(value *string)()
}
