package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret secret information
type V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // secret id
    id *string
    // secret name
    name *string
    // secret type
    typeEscaped *string
}
// NewV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret instantiates a new V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret and sets the default values.
func NewV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret()(*V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) {
    m := &V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secretFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secretFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. secret id
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. secret name
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. secret type
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. secret id
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. secret name
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. secret type
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secret) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SystemsItemServersItemWithServerGetResponse_iloUserCredential_secretable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
}
