package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspAccountsItemCspTagKeysGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of items in this response.
    count *int32
    // List of CSP tag keys.
    items []V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable
}
// NewV1beta1CspAccountsItemCspTagKeysGetResponse instantiates a new V1beta1CspAccountsItemCspTagKeysGetResponse and sets the default values.
func NewV1beta1CspAccountsItemCspTagKeysGetResponse()(*V1beta1CspAccountsItemCspTagKeysGetResponse) {
    m := &V1beta1CspAccountsItemCspTagKeysGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspTagKeysGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspTagKeysGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspTagKeysGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Number of items in this response.
// returns a *int32 when successful
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspAccountsItemCspTagKeysGetResponse_itemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. List of CSP tag keys.
// returns a []V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable when successful
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) GetItems()([]V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable) {
    return m.items
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
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
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Number of items in this response.
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) SetCount(value *int32)() {
    m.count = value
}
// SetItems sets the items property value. List of CSP tag keys.
func (m *V1beta1CspAccountsItemCspTagKeysGetResponse) SetItems(value []V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable)() {
    m.items = value
}
type V1beta1CspAccountsItemCspTagKeysGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetItems()([]V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable)
    SetCount(value *int32)()
    SetItems(value []V1beta1CspAccountsItemCspTagKeysGetResponse_itemsable)()
}