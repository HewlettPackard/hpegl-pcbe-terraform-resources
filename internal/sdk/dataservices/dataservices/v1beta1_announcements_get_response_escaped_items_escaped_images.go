package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AnnouncementsGetResponse_items_images a reference to an image associated with an announcement.
type V1beta1AnnouncementsGetResponse_items_images struct {
    // URL of an image associated with an announcement.
    url *string
}
// NewV1beta1AnnouncementsGetResponse_items_images instantiates a new V1beta1AnnouncementsGetResponse_items_images and sets the default values.
func NewV1beta1AnnouncementsGetResponse_items_images()(*V1beta1AnnouncementsGetResponse_items_images) {
    m := &V1beta1AnnouncementsGetResponse_items_images{
    }
    return m
}
// CreateV1beta1AnnouncementsGetResponse_items_imagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AnnouncementsGetResponse_items_imagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AnnouncementsGetResponse_items_images(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AnnouncementsGetResponse_items_images) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. URL of an image associated with an announcement.
// returns a *string when successful
func (m *V1beta1AnnouncementsGetResponse_items_images) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *V1beta1AnnouncementsGetResponse_items_images) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUrl sets the url property value. URL of an image associated with an announcement.
func (m *V1beta1AnnouncementsGetResponse_items_images) SetUrl(value *string)() {
    m.url = value
}
type V1beta1AnnouncementsGetResponse_items_imagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUrl()(*string)
    SetUrl(value *string)()
}
