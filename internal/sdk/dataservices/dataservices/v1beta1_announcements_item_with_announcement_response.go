package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AnnouncementsItemWithAnnouncementResponse all possible properties of an announcement resource.
// Deprecated: This class is obsolete. Use V1beta1AnnouncementsItemWithAnnouncementGetResponseable instead.
type V1beta1AnnouncementsItemWithAnnouncementResponse struct {
    V1beta1AnnouncementsItemWithAnnouncementGetResponse
}
// NewV1beta1AnnouncementsItemWithAnnouncementResponse instantiates a new V1beta1AnnouncementsItemWithAnnouncementResponse and sets the default values.
func NewV1beta1AnnouncementsItemWithAnnouncementResponse()(*V1beta1AnnouncementsItemWithAnnouncementResponse) {
    m := &V1beta1AnnouncementsItemWithAnnouncementResponse{
        V1beta1AnnouncementsItemWithAnnouncementGetResponse: *NewV1beta1AnnouncementsItemWithAnnouncementGetResponse(),
    }
    return m
}
// CreateV1beta1AnnouncementsItemWithAnnouncementResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AnnouncementsItemWithAnnouncementResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AnnouncementsItemWithAnnouncementResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1AnnouncementsItemWithAnnouncementGetResponseable instead.
type V1beta1AnnouncementsItemWithAnnouncementResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1AnnouncementsItemWithAnnouncementGetResponseable
}