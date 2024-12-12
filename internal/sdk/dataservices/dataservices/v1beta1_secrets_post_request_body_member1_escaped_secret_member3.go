package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember3 bearer Token Authentication secret definition
type V1beta1SecretsPostRequestBodyMember1_secretMember3 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Bearer Token Authentication secret properties
    bearerAuth V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember3 instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember3 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember3()(*V1beta1SecretsPostRequestBodyMember1_secretMember3) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember3{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember3FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember3FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember3(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBearerAuth gets the bearerAuth property value. Bearer Token Authentication secret properties
// returns a V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) GetBearerAuth()(V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable) {
    return m.bearerAuth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bearerAuth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBearerAuth(val.(V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bearerAuth", m.GetBearerAuth())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBearerAuth sets the bearerAuth property value. Bearer Token Authentication secret properties
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember3) SetBearerAuth(value V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable)() {
    m.bearerAuth = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember3able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBearerAuth()(V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable)
    SetBearerAuth(value V1beta1SecretsPostRequestBodyMember1_secretMember3_bearerAuthable)()
}
