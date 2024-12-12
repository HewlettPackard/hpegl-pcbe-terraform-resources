package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AnnouncementsGetResponse_items_services a reference to a specific service relevant to an announcement.
type V1beta1AnnouncementsGetResponse_items_services struct {
    // A unique service identifier.
    id *string
    // Human-readable service name.
    name *string
}
// NewV1beta1AnnouncementsGetResponse_items_services instantiates a new V1beta1AnnouncementsGetResponse_items_services and sets the default values.
func NewV1beta1AnnouncementsGetResponse_items_services()(*V1beta1AnnouncementsGetResponse_items_services) {
    m := &V1beta1AnnouncementsGetResponse_items_services{
    }
    return m
}
// CreateV1beta1AnnouncementsGetResponse_items_servicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AnnouncementsGetResponse_items_servicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AnnouncementsGetResponse_items_services(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AnnouncementsGetResponse_items_services) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetId gets the id property value. A unique service identifier.
// returns a *string when successful
func (m *V1beta1AnnouncementsGetResponse_items_services) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Human-readable service name.
// returns a *string when successful
func (m *V1beta1AnnouncementsGetResponse_items_services) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1AnnouncementsGetResponse_items_services) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetId sets the id property value. A unique service identifier.
func (m *V1beta1AnnouncementsGetResponse_items_services) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Human-readable service name.
func (m *V1beta1AnnouncementsGetResponse_items_services) SetName(value *string)() {
    m.name = value
}
type V1beta1AnnouncementsGetResponse_items_servicesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    SetId(value *string)()
    SetName(value *string)()
}
