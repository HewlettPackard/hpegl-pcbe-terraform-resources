package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchResponse_subclassifier resource Subclassifier Definition (e.g. BASIC_AUTH)
type V1beta1SecretsItemSecretsPatchResponse_subclassifier struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the resource subclassifier
    name *string
    // Resource subclassifier properties
    properties V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable
}
// NewV1beta1SecretsItemSecretsPatchResponse_subclassifier instantiates a new V1beta1SecretsItemSecretsPatchResponse_subclassifier and sets the default values.
func NewV1beta1SecretsItemSecretsPatchResponse_subclassifier()(*V1beta1SecretsItemSecretsPatchResponse_subclassifier) {
    m := &V1beta1SecretsItemSecretsPatchResponse_subclassifier{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchResponse_subclassifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchResponse_subclassifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchResponse_subclassifier(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the resource subclassifier
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) GetName()(*string) {
    return m.name
}
// GetProperties gets the properties property value. Resource subclassifier properties
// returns a V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable when successful
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) GetProperties()(V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("properties", m.GetProperties())
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
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the resource subclassifier
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the properties property value. Resource subclassifier properties
func (m *V1beta1SecretsItemSecretsPatchResponse_subclassifier) SetProperties(value V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable)() {
    m.properties = value
}
type V1beta1SecretsItemSecretsPatchResponse_subclassifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetProperties()(V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable)
    SetName(value *string)()
    SetProperties(value V1beta1SecretsItemSecretsPatchResponse_subclassifier_propertiesable)()
}
