package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_osUserCredential operating system user name and authentication credential secret. Property hypervisorHostRootUserPasswordSecret will be deprecated.Use this property instead.
type V1beta1SystemsItemServersGetResponse_items_osUserCredential struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Operating system username
    name *string
    // Secret information
    secret V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable
}
// NewV1beta1SystemsItemServersGetResponse_items_osUserCredential instantiates a new V1beta1SystemsItemServersGetResponse_items_osUserCredential and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_osUserCredential()(*V1beta1SystemsItemServersGetResponse_items_osUserCredential) {
    m := &V1beta1SystemsItemServersGetResponse_items_osUserCredential{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_osUserCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_osUserCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_osUserCredential(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemServersGetResponse_items_osUserCredential_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Operating system username
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) GetName()(*string) {
    return m.name
}
// GetSecret gets the secret property value. Secret information
// returns a V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable when successful
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) GetSecret()(V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable) {
    return m.secret
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Operating system username
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) SetName(value *string)() {
    m.name = value
}
// SetSecret sets the secret property value. Secret information
func (m *V1beta1SystemsItemServersGetResponse_items_osUserCredential) SetSecret(value V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable)() {
    m.secret = value
}
type V1beta1SystemsItemServersGetResponse_items_osUserCredentialable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetSecret()(V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable)
    SetName(value *string)()
    SetSecret(value V1beta1SystemsItemServersGetResponse_items_osUserCredential_secretable)()
}
