package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsGetResponse_items_location system Location Information.
type V1beta1SystemsGetResponse_items_location struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // City
    city *string
    // Country Code
    countryCode *string
    // Latitude coordinate of the system location
    latitude *string
    // Longitude coordinate of the system location
    longitude *string
    // State Code
    stateCode *string
    // Zip Code
    zipCode *string
}
// NewV1beta1SystemsGetResponse_items_location instantiates a new V1beta1SystemsGetResponse_items_location and sets the default values.
func NewV1beta1SystemsGetResponse_items_location()(*V1beta1SystemsGetResponse_items_location) {
    m := &V1beta1SystemsGetResponse_items_location{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsGetResponse_items_locationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsGetResponse_items_locationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsGetResponse_items_location(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsGetResponse_items_location) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCity gets the city property value. City
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetCity()(*string) {
    return m.city
}
// GetCountryCode gets the countryCode property value. Country Code
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetCountryCode()(*string) {
    return m.countryCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsGetResponse_items_location) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
        }
        return nil
    }
    res["stateCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateCode(val)
        }
        return nil
    }
    res["zipCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZipCode(val)
        }
        return nil
    }
    return res
}
// GetLatitude gets the latitude property value. Latitude coordinate of the system location
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetLatitude()(*string) {
    return m.latitude
}
// GetLongitude gets the longitude property value. Longitude coordinate of the system location
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetLongitude()(*string) {
    return m.longitude
}
// GetStateCode gets the stateCode property value. State Code
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetStateCode()(*string) {
    return m.stateCode
}
// GetZipCode gets the zipCode property value. Zip Code
// returns a *string when successful
func (m *V1beta1SystemsGetResponse_items_location) GetZipCode()(*string) {
    return m.zipCode
}
// Serialize serializes information the current object
func (m *V1beta1SystemsGetResponse_items_location) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("longitude", m.GetLongitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateCode", m.GetStateCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("zipCode", m.GetZipCode())
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
func (m *V1beta1SystemsGetResponse_items_location) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCity sets the city property value. City
func (m *V1beta1SystemsGetResponse_items_location) SetCity(value *string)() {
    m.city = value
}
// SetCountryCode sets the countryCode property value. Country Code
func (m *V1beta1SystemsGetResponse_items_location) SetCountryCode(value *string)() {
    m.countryCode = value
}
// SetLatitude sets the latitude property value. Latitude coordinate of the system location
func (m *V1beta1SystemsGetResponse_items_location) SetLatitude(value *string)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. Longitude coordinate of the system location
func (m *V1beta1SystemsGetResponse_items_location) SetLongitude(value *string)() {
    m.longitude = value
}
// SetStateCode sets the stateCode property value. State Code
func (m *V1beta1SystemsGetResponse_items_location) SetStateCode(value *string)() {
    m.stateCode = value
}
// SetZipCode sets the zipCode property value. Zip Code
func (m *V1beta1SystemsGetResponse_items_location) SetZipCode(value *string)() {
    m.zipCode = value
}
type V1beta1SystemsGetResponse_items_locationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryCode()(*string)
    GetLatitude()(*string)
    GetLongitude()(*string)
    GetStateCode()(*string)
    GetZipCode()(*string)
    SetCity(value *string)()
    SetCountryCode(value *string)()
    SetLatitude(value *string)()
    SetLongitude(value *string)()
    SetStateCode(value *string)()
    SetZipCode(value *string)()
}
