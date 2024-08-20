package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1HypervisorHostsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total number of records returned.
    count *int32
    // The items property
    items []V1beta1HypervisorHostsGetResponse_itemsable
    // The number of items to skip before starting to collect the result set
    offset *int32
    // Total number of documents matching filter criteria.
    total *int32
}
// NewV1beta1HypervisorHostsGetResponse instantiates a new V1beta1HypervisorHostsGetResponse and sets the default values.
func NewV1beta1HypervisorHostsGetResponse()(*V1beta1HypervisorHostsGetResponse) {
    m := &V1beta1HypervisorHostsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Total number of records returned.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorHostsGetResponse_itemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorHostsGetResponse_itemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorHostsGetResponse_itemsable)
                }
            }
            m.SetItems(res)
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
// GetItems gets the items property value. The items property
// returns a []V1beta1HypervisorHostsGetResponse_itemsable when successful
func (m *V1beta1HypervisorHostsGetResponse) GetItems()([]V1beta1HypervisorHostsGetResponse_itemsable) {
    return m.items
}
// GetOffset gets the offset property value. The number of items to skip before starting to collect the result set
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse) GetOffset()(*int32) {
    return m.offset
}
// GetTotal gets the total property value. Total number of documents matching filter criteria.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1HypervisorHostsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Total number of records returned.
func (m *V1beta1HypervisorHostsGetResponse) SetCount(value *int32)() {
    m.count = value
}
// SetItems sets the items property value. The items property
func (m *V1beta1HypervisorHostsGetResponse) SetItems(value []V1beta1HypervisorHostsGetResponse_itemsable)() {
    m.items = value
}
// SetOffset sets the offset property value. The number of items to skip before starting to collect the result set
func (m *V1beta1HypervisorHostsGetResponse) SetOffset(value *int32)() {
    m.offset = value
}
// SetTotal sets the total property value. Total number of documents matching filter criteria.
func (m *V1beta1HypervisorHostsGetResponse) SetTotal(value *int32)() {
    m.total = value
}
type V1beta1HypervisorHostsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetItems()([]V1beta1HypervisorHostsGetResponse_itemsable)
    GetOffset()(*int32)
    GetTotal()(*int32)
    SetCount(value *int32)()
    SetItems(value []V1beta1HypervisorHostsGetResponse_itemsable)()
    SetOffset(value *int32)()
    SetTotal(value *int32)()
}
