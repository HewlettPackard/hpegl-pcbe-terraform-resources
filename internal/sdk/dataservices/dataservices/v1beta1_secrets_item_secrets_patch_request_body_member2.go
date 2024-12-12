package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsPatchRequestBodyMember2 secret Resource Redefinition
type V1beta1SecretsItemSecretsPatchRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description for the resource
    description *string
    // Groups to be associated with the resource
    groups []string
    // Secret policy identifier
    policy *string
}
// NewV1beta1SecretsItemSecretsPatchRequestBodyMember2 instantiates a new V1beta1SecretsItemSecretsPatchRequestBodyMember2 and sets the default values.
func NewV1beta1SecretsItemSecretsPatchRequestBodyMember2()(*V1beta1SecretsItemSecretsPatchRequestBodyMember2) {
    m := &V1beta1SecretsItemSecretsPatchRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecretsPatchRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsPatchRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsPatchRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description for the resource
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Groups to be associated with the resource
// returns a []string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) GetGroups()([]string) {
    return m.groups
}
// GetPolicy gets the policy property value. Secret policy identifier
// returns a *string when successful
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) GetPolicy()(*string) {
    return m.policy
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        err := writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policy", m.GetPolicy())
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
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description for the resource
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) SetDescription(value *string)() {
    m.description = value
}
// SetGroups sets the groups property value. Groups to be associated with the resource
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) SetGroups(value []string)() {
    m.groups = value
}
// SetPolicy sets the policy property value. Secret policy identifier
func (m *V1beta1SecretsItemSecretsPatchRequestBodyMember2) SetPolicy(value *string)() {
    m.policy = value
}
type V1beta1SecretsItemSecretsPatchRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetGroups()([]string)
    GetPolicy()(*string)
    SetDescription(value *string)()
    SetGroups(value []string)()
    SetPolicy(value *string)()
}
