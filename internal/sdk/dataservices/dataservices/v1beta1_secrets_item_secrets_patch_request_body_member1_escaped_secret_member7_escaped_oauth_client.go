package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient oAuth Client secret properties
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // OAuth Client ID
    clientId *string
    // OAuth Client secret
    clientSecret *string
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the clientId property value. OAuth Client ID
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. OAuth Client secret
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientSecret", m.GetClientSecret())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the clientId property value. OAuth Client ID
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. OAuth Client secret
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClient) SetClientSecret(value *string)() {
    m.clientSecret = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember7_oauthClientable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
}
