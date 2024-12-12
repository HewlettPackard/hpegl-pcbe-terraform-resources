package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember1 azure Service Principal client secret definition
type V1beta1SecretsPostRequestBodyMember1_secretMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Azure Service Principal client secret properties
    azureSPClient V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember1 instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember1 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember1()(*V1beta1SecretsPostRequestBodyMember1_secretMember1) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAzureSPClient gets the azureSPClient property value. Azure Service Principal client secret properties
// returns a V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) GetAzureSPClient()(V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable) {
    return m.azureSPClient
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureSPClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSPClient(val.(V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("azureSPClient", m.GetAzureSPClient())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAzureSPClient sets the azureSPClient property value. Azure Service Principal client secret properties
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember1) SetAzureSPClient(value V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable)() {
    m.azureSPClient = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureSPClient()(V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable)
    SetAzureSPClient(value V1beta1SecretsPostRequestBodyMember1_secretMember1_azureSPClientable)()
}
