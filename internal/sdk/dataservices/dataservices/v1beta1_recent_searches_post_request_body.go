package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1RecentSearchesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The resource property
    resource V1beta1RecentSearchesPostRequestBody_resourceable
    // The service that the search was made within
    searchedService *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1RecentSearchesPostRequestBody instantiates a new V1beta1RecentSearchesPostRequestBody and sets the default values.
func NewV1beta1RecentSearchesPostRequestBody()(*V1beta1RecentSearchesPostRequestBody) {
    m := &V1beta1RecentSearchesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1RecentSearchesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1RecentSearchesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1RecentSearchesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1RecentSearchesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1RecentSearchesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1RecentSearchesPostRequestBody_resourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(V1beta1RecentSearchesPostRequestBody_resourceable))
        }
        return nil
    }
    res["searchedService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchedService(val)
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
// GetResource gets the resource property value. The resource property
// returns a V1beta1RecentSearchesPostRequestBody_resourceable when successful
func (m *V1beta1RecentSearchesPostRequestBody) GetResource()(V1beta1RecentSearchesPostRequestBody_resourceable) {
    return m.resource
}
// GetSearchedService gets the searchedService property value. The service that the search was made within
// returns a *string when successful
func (m *V1beta1RecentSearchesPostRequestBody) GetSearchedService()(*string) {
    return m.searchedService
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1RecentSearchesPostRequestBody) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1RecentSearchesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("searchedService", m.GetSearchedService())
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
func (m *V1beta1RecentSearchesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResource sets the resource property value. The resource property
func (m *V1beta1RecentSearchesPostRequestBody) SetResource(value V1beta1RecentSearchesPostRequestBody_resourceable)() {
    m.resource = value
}
// SetSearchedService sets the searchedService property value. The service that the search was made within
func (m *V1beta1RecentSearchesPostRequestBody) SetSearchedService(value *string)() {
    m.searchedService = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1RecentSearchesPostRequestBody) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1RecentSearchesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResource()(V1beta1RecentSearchesPostRequestBody_resourceable)
    GetSearchedService()(*string)
    GetTypeEscaped()(*string)
    SetResource(value V1beta1RecentSearchesPostRequestBody_resourceable)()
    SetSearchedService(value *string)()
    SetTypeEscaped(value *string)()
}
