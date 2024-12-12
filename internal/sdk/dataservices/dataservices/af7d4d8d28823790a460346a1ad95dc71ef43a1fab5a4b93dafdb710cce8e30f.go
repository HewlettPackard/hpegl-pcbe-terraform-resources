package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient azure Service Principal client secret properties
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Azure Service Principal client ID (aka application ID)
    clientId *string
    // Azure Service Principal client secret
    clientSecret *string
    // Azure Service Principal tenant ID (aka directory ID)
    tenantId *string
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient()(*V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the clientId property value. Azure Service Principal client ID (aka application ID)
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. Azure Service Principal client secret
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. Azure Service Principal tenant ID (aka directory ID)
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the clientId property value. Azure Service Principal client ID (aka application ID)
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. Azure Service Principal client secret
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetTenantId sets the tenantId property value. Azure Service Principal tenant ID (aka directory ID)
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClient) SetTenantId(value *string)() {
    m.tenantId = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember1_secretMember1_azureSPClientable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetTenantId()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetTenantId(value *string)()
}
