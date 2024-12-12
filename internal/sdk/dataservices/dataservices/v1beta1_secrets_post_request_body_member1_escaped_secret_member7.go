package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember7 oAuth Client secret definition
type V1beta1SecretsPostRequestBodyMember1_secretMember7 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // OAuth Client secret properties
    oauthClient V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember7 instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember7 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember7()(*V1beta1SecretsPostRequestBodyMember1_secretMember7) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember7{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember7FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember7FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember7(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["oauthClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOauthClient(val.(V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable))
        }
        return nil
    }
    return res
}
// GetOauthClient gets the oauthClient property value. OAuth Client secret properties
// returns a V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) GetOauthClient()(V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable) {
    return m.oauthClient
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("oauthClient", m.GetOauthClient())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOauthClient sets the oauthClient property value. OAuth Client secret properties
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember7) SetOauthClient(value V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable)() {
    m.oauthClient = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember7able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOauthClient()(V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable)
    SetOauthClient(value V1beta1SecretsPostRequestBodyMember1_secretMember7_oauthClientable)()
}
