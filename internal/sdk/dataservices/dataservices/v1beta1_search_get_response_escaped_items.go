package dataservices

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SearchGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Map of fields that matched the search query. The map contains one property for each resource field that matched the search term. The property value is an array of one or more sub-strings that matched the search.
    matchingFields V1beta1SearchGetResponse_items_matchingFieldsable
    // The resource property
    resource V1beta1SearchGetResponse_items_resourceable
    // A score assigned by the search engine.
    score *float32
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1SearchGetResponse_items instantiates a new V1beta1SearchGetResponse_items and sets the default values.
func NewV1beta1SearchGetResponse_items()(*V1beta1SearchGetResponse_items) {
    m := &V1beta1SearchGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SearchGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SearchGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SearchGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["matchingFields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SearchGetResponse_items_matchingFieldsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingFields(val.(V1beta1SearchGetResponse_items_matchingFieldsable))
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SearchGetResponse_items_resourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(V1beta1SearchGetResponse_items_resourceable))
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
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
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1SearchGetResponse_items) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetMatchingFields gets the matchingFields property value. Map of fields that matched the search query. The map contains one property for each resource field that matched the search term. The property value is an array of one or more sub-strings that matched the search.
// returns a V1beta1SearchGetResponse_items_matchingFieldsable when successful
func (m *V1beta1SearchGetResponse_items) GetMatchingFields()(V1beta1SearchGetResponse_items_matchingFieldsable) {
    return m.matchingFields
}
// GetResource gets the resource property value. The resource property
// returns a V1beta1SearchGetResponse_items_resourceable when successful
func (m *V1beta1SearchGetResponse_items) GetResource()(V1beta1SearchGetResponse_items_resourceable) {
    return m.resource
}
// GetScore gets the score property value. A score assigned by the search engine.
// returns a *float32 when successful
func (m *V1beta1SearchGetResponse_items) GetScore()(*float32) {
    return m.score
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SearchGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SearchGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("matchingFields", m.GetMatchingFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("score", m.GetScore())
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
func (m *V1beta1SearchGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SearchGetResponse_items) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetMatchingFields sets the matchingFields property value. Map of fields that matched the search query. The map contains one property for each resource field that matched the search term. The property value is an array of one or more sub-strings that matched the search.
func (m *V1beta1SearchGetResponse_items) SetMatchingFields(value V1beta1SearchGetResponse_items_matchingFieldsable)() {
    m.matchingFields = value
}
// SetResource sets the resource property value. The resource property
func (m *V1beta1SearchGetResponse_items) SetResource(value V1beta1SearchGetResponse_items_resourceable)() {
    m.resource = value
}
// SetScore sets the score property value. A score assigned by the search engine.
func (m *V1beta1SearchGetResponse_items) SetScore(value *float32)() {
    m.score = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SearchGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SearchGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetMatchingFields()(V1beta1SearchGetResponse_items_matchingFieldsable)
    GetResource()(V1beta1SearchGetResponse_items_resourceable)
    GetScore()(*float32)
    GetTypeEscaped()(*string)
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetMatchingFields(value V1beta1SearchGetResponse_items_matchingFieldsable)()
    SetResource(value V1beta1SearchGetResponse_items_resourceable)()
    SetScore(value *float32)()
    SetTypeEscaped(value *string)()
}
