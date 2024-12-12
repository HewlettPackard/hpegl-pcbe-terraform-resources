package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SearchGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of items in this response.
    count *int32
    // Indicates whether or not items is the result of a fuzzy text searchrather than an exact search query term string match
    isFuzzy *bool
    // List of search results
    items []V1beta1SearchGetResponse_itemsable
    // The offset query parameter from the request.
    offset *int32
    // Timestamp at which the search was performed
    searchedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Total number of items matching the filter parameter in the request.
    total *int32
}
// NewV1beta1SearchGetResponse instantiates a new V1beta1SearchGetResponse and sets the default values.
func NewV1beta1SearchGetResponse()(*V1beta1SearchGetResponse) {
    m := &V1beta1SearchGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SearchGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SearchGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Number of items in this response.
// returns a *int32 when successful
func (m *V1beta1SearchGetResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SearchGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isFuzzy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFuzzy(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SearchGetResponse_itemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SearchGetResponse_itemsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SearchGetResponse_itemsable)
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
    res["searchedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchedAt(val)
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
// GetIsFuzzy gets the isFuzzy property value. Indicates whether or not items is the result of a fuzzy text searchrather than an exact search query term string match
// returns a *bool when successful
func (m *V1beta1SearchGetResponse) GetIsFuzzy()(*bool) {
    return m.isFuzzy
}
// GetItems gets the items property value. List of search results
// returns a []V1beta1SearchGetResponse_itemsable when successful
func (m *V1beta1SearchGetResponse) GetItems()([]V1beta1SearchGetResponse_itemsable) {
    return m.items
}
// GetOffset gets the offset property value. The offset query parameter from the request.
// returns a *int32 when successful
func (m *V1beta1SearchGetResponse) GetOffset()(*int32) {
    return m.offset
}
// GetSearchedAt gets the searchedAt property value. Timestamp at which the search was performed
// returns a *Time when successful
func (m *V1beta1SearchGetResponse) GetSearchedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.searchedAt
}
// GetTotal gets the total property value. Total number of items matching the filter parameter in the request.
// returns a *int32 when successful
func (m *V1beta1SearchGetResponse) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *V1beta1SearchGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFuzzy", m.GetIsFuzzy())
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
        err := writer.WriteTimeValue("searchedAt", m.GetSearchedAt())
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
func (m *V1beta1SearchGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Number of items in this response.
func (m *V1beta1SearchGetResponse) SetCount(value *int32)() {
    m.count = value
}
// SetIsFuzzy sets the isFuzzy property value. Indicates whether or not items is the result of a fuzzy text searchrather than an exact search query term string match
func (m *V1beta1SearchGetResponse) SetIsFuzzy(value *bool)() {
    m.isFuzzy = value
}
// SetItems sets the items property value. List of search results
func (m *V1beta1SearchGetResponse) SetItems(value []V1beta1SearchGetResponse_itemsable)() {
    m.items = value
}
// SetOffset sets the offset property value. The offset query parameter from the request.
func (m *V1beta1SearchGetResponse) SetOffset(value *int32)() {
    m.offset = value
}
// SetSearchedAt sets the searchedAt property value. Timestamp at which the search was performed
func (m *V1beta1SearchGetResponse) SetSearchedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.searchedAt = value
}
// SetTotal sets the total property value. Total number of items matching the filter parameter in the request.
func (m *V1beta1SearchGetResponse) SetTotal(value *int32)() {
    m.total = value
}
type V1beta1SearchGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetIsFuzzy()(*bool)
    GetItems()([]V1beta1SearchGetResponse_itemsable)
    GetOffset()(*int32)
    GetSearchedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotal()(*int32)
    SetCount(value *int32)()
    SetIsFuzzy(value *bool)()
    SetItems(value []V1beta1SearchGetResponse_itemsable)()
    SetOffset(value *int32)()
    SetSearchedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotal(value *int32)()
}
