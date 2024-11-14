package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_iloUserCredential iLO user name and authentication credential secret. Property iloAdminUserPasswordSecret will be deprecated.Use this property instead.
type V1beta1SystemsItemServersGetResponse_items_iloUserCredential struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // iLO username
    name *string
    // Secret information
    secret V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable
}
// NewV1beta1SystemsItemServersGetResponse_items_iloUserCredential instantiates a new V1beta1SystemsItemServersGetResponse_items_iloUserCredential and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_iloUserCredential()(*V1beta1SystemsItemServersGetResponse_items_iloUserCredential) {
    m := &V1beta1SystemsItemServersGetResponse_items_iloUserCredential{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_iloUserCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_iloUserCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_iloUserCredential(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. iLO username
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) GetName()(*string) {
    return m.name
}
// GetSecret gets the secret property value. Secret information
// returns a V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) GetSecret()(V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable) {
    return m.secret
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret", m.GetSecret())
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
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. iLO username
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) SetName(value *string)() {
    m.name = value
}
// SetSecret sets the secret property value. Secret information
func (m *V1beta1SystemsItemServersGetResponse_items_iloUserCredential) SetSecret(value V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable)() {
    m.secret = value
}
type V1beta1SystemsItemServersGetResponse_items_iloUserCredentialable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetSecret()(V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable)
    SetName(value *string)()
    SetSecret(value V1beta1SystemsItemServersGetResponse_items_iloUserCredential_secretable)()
}
