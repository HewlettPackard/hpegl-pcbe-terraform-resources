package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsGetResponse_classifier resource Classifier Definition (e.g. SECRET)
type V1beta1SecretsItemSecretsGetResponse_classifier struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the resource classifier
    name *string
}
// NewV1beta1SecretsItemSecretsGetResponse_classifier instantiates a new V1beta1SecretsItemSecretsGetResponse_classifier and sets the default values.
func NewV1beta1SecretsItemSecretsGetResponse_classifier()(*V1beta1SecretsItemSecretsGetResponse_classifier) {
    m := &V1beta1SecretsItemSecretsGetResponse_classifier{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsGetResponse_classifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsGetResponse_classifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsGetResponse_classifier(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the resource classifier
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the resource classifier
func (m *V1beta1SecretsItemSecretsGetResponse_classifier) SetName(value *string)() {
    m.name = value
}
type V1beta1SecretsItemSecretsGetResponse_classifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    SetName(value *string)()
}
