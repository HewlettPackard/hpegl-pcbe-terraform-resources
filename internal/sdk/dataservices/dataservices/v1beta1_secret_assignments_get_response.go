package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsGetResponse report Assignments Response
type V1beta1SecretAssignmentsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Current results page count
    count *int32
    // Assignment resource definitions
    items []V1beta1SecretAssignmentsGetResponse_itemsable
    // Cursor for next results page
    next *string
    // Offset of current results page
    offset *int32
    // Total results count
    total *int32
}
// NewV1beta1SecretAssignmentsGetResponse instantiates a new V1beta1SecretAssignmentsGetResponse and sets the default values.
func NewV1beta1SecretAssignmentsGetResponse()(*V1beta1SecretAssignmentsGetResponse) {
    m := &V1beta1SecretAssignmentsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretAssignmentsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Current results page count
// returns a *int32 when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SecretAssignmentsGetResponse_itemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SecretAssignmentsGetResponse_itemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SecretAssignmentsGetResponse_itemsable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["next"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNext(val)
        }
        return nil
    }
    res["offset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. Assignment resource definitions
// returns a []V1beta1SecretAssignmentsGetResponse_itemsable when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetItems()([]V1beta1SecretAssignmentsGetResponse_itemsable) {
    return m.items
}
// GetNext gets the next property value. Cursor for next results page
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetNext()(*string) {
    return m.next
}
// GetOffset gets the offset property value. Offset of current results page
// returns a *int32 when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetOffset()(*int32) {
    return m.offset
}
// GetTotal gets the total property value. Total results count
// returns a *int32 when successful
func (m *V1beta1SecretAssignmentsGetResponse) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *V1beta1SecretAssignmentsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("next", m.GetNext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offset", m.GetOffset())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *V1beta1SecretAssignmentsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Current results page count
func (m *V1beta1SecretAssignmentsGetResponse) SetCount(value *int32)() {
    m.count = value
}
// SetItems sets the items property value. Assignment resource definitions
func (m *V1beta1SecretAssignmentsGetResponse) SetItems(value []V1beta1SecretAssignmentsGetResponse_itemsable)() {
    m.items = value
}
// SetNext sets the next property value. Cursor for next results page
func (m *V1beta1SecretAssignmentsGetResponse) SetNext(value *string)() {
    m.next = value
}
// SetOffset sets the offset property value. Offset of current results page
func (m *V1beta1SecretAssignmentsGetResponse) SetOffset(value *int32)() {
    m.offset = value
}
// SetTotal sets the total property value. Total results count
func (m *V1beta1SecretAssignmentsGetResponse) SetTotal(value *int32)() {
    m.total = value
}
type V1beta1SecretAssignmentsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetItems()([]V1beta1SecretAssignmentsGetResponse_itemsable)
    GetNext()(*string)
    GetOffset()(*int32)
    GetTotal()(*int32)
    SetCount(value *int32)()
    SetItems(value []V1beta1SecretAssignmentsGetResponse_itemsable)()
    SetNext(value *string)()
    SetOffset(value *int32)()
    SetTotal(value *int32)()
}
