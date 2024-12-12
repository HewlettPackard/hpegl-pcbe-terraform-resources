package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsGetResponse_items_dateTime this details the current date and time of the Data Services Connector, how the current date and time is determined, as well as the user specified timezone.
type V1beta1DataServicesConnectorsGetResponse_items_dateTime struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time zone in which the Data Services Connector is deployed.
    timezone *string
}
// NewV1beta1DataServicesConnectorsGetResponse_items_dateTime instantiates a new V1beta1DataServicesConnectorsGetResponse_items_dateTime and sets the default values.
func NewV1beta1DataServicesConnectorsGetResponse_items_dateTime()(*V1beta1DataServicesConnectorsGetResponse_items_dateTime) {
    m := &V1beta1DataServicesConnectorsGetResponse_items_dateTime{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsGetResponse_items_dateTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsGetResponse_items_dateTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsGetResponse_items_dateTime(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["timezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimezone(val)
        }
        return nil
    }
    return res
}
// GetTimezone gets the timezone property value. Time zone in which the Data Services Connector is deployed.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) GetTimezone()(*string) {
    return m.timezone
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("timezone", m.GetTimezone())
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
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTimezone sets the timezone property value. Time zone in which the Data Services Connector is deployed.
func (m *V1beta1DataServicesConnectorsGetResponse_items_dateTime) SetTimezone(value *string)() {
    m.timezone = value
}
type V1beta1DataServicesConnectorsGetResponse_items_dateTimeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTimezone()(*string)
    SetTimezone(value *string)()
}
