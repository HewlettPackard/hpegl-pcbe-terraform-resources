package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspAccountsItemCspTagsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the resource, usually a UUID.
    id *string
    // A CSP tag key.
    key *string
    // The type of resource.
    typeEscaped *string
    // A CSP tag value.
    value *string
}
// NewV1beta1CspAccountsItemCspTagsGetResponse_items instantiates a new V1beta1CspAccountsItemCspTagsGetResponse_items and sets the default values.
func NewV1beta1CspAccountsItemCspTagsGetResponse_items()(*V1beta1CspAccountsItemCspTagsGetResponse_items) {
    m := &V1beta1CspAccountsItemCspTagsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspTagsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspTagsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspTagsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetId()(*string) {
    return m.id
}
// GetKey gets the key property value. A CSP tag key.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetKey()(*string) {
    return m.key
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. A CSP tag value.
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetKey sets the key property value. A CSP tag key.
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) SetKey(value *string)() {
    m.key = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. A CSP tag value.
func (m *V1beta1CspAccountsItemCspTagsGetResponse_items) SetValue(value *string)() {
    m.value = value
}
type V1beta1CspAccountsItemCspTagsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetKey()(*string)
    GetTypeEscaped()(*string)
    GetValue()(*string)
    SetId(value *string)()
    SetKey(value *string)()
    SetTypeEscaped(value *string)()
    SetValue(value *string)()
}
