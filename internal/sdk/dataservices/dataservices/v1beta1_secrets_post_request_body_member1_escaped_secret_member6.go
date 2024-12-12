package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostRequestBodyMember1_secretMember6 jSON Web Token secret definition
type V1beta1SecretsPostRequestBodyMember1_secretMember6 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // JSON Web Token secret properties
    jsonWebToken V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable
}
// NewV1beta1SecretsPostRequestBodyMember1_secretMember6 instantiates a new V1beta1SecretsPostRequestBodyMember1_secretMember6 and sets the default values.
func NewV1beta1SecretsPostRequestBodyMember1_secretMember6()(*V1beta1SecretsPostRequestBodyMember1_secretMember6) {
    m := &V1beta1SecretsPostRequestBodyMember1_secretMember6{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostRequestBodyMember1_secretMember6FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostRequestBodyMember1_secretMember6FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostRequestBodyMember1_secretMember6(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["jsonWebToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonWebToken(val.(V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable))
        }
        return nil
    }
    return res
}
// GetJsonWebToken gets the jsonWebToken property value. JSON Web Token secret properties
// returns a V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable when successful
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) GetJsonWebToken()(V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable) {
    return m.jsonWebToken
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("jsonWebToken", m.GetJsonWebToken())
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
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJsonWebToken sets the jsonWebToken property value. JSON Web Token secret properties
func (m *V1beta1SecretsPostRequestBodyMember1_secretMember6) SetJsonWebToken(value V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable)() {
    m.jsonWebToken = value
}
type V1beta1SecretsPostRequestBodyMember1_secretMember6able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetJsonWebToken()(V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable)
    SetJsonWebToken(value V1beta1SecretsPostRequestBodyMember1_secretMember6_jsonWebTokenable)()
}
