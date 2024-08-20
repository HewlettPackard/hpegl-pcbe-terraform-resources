package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails details of the replication partnerDetails:
type V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the storage system, applicable only for Nimble storage systems
    id *string
    // Name of the replication partner Array
    name *string
    // storage system wwn in case of Primera
    systemWwn *string
}
// NewV1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails instantiates a new V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails and sets the default values.
func NewV1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails()(*V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) {
    m := &V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresGetResponse_items_replicationInfo_partnerDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresGetResponse_items_replicationInfo_partnerDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["systemWwn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemWwn(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Id of the storage system, applicable only for Nimble storage systems
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the replication partner Array
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) GetName()(*string) {
    return m.name
}
// GetSystemWwn gets the systemWwn property value. storage system wwn in case of Primera
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) GetSystemWwn()(*string) {
    return m.systemWwn
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("systemWwn", m.GetSystemWwn())
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
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Id of the storage system, applicable only for Nimble storage systems
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the replication partner Array
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) SetName(value *string)() {
    m.name = value
}
// SetSystemWwn sets the systemWwn property value. storage system wwn in case of Primera
func (m *V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetails) SetSystemWwn(value *string)() {
    m.systemWwn = value
}
type V1beta1DatastoresGetResponse_items_replicationInfo_partnerDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetSystemWwn()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetSystemWwn(value *string)()
}
