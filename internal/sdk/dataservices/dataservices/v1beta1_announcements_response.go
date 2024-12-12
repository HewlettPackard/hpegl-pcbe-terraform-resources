package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1AnnouncementsGetResponseable instead.
type V1beta1AnnouncementsResponse struct {
    V1beta1AnnouncementsGetResponse
}
// NewV1beta1AnnouncementsResponse instantiates a new V1beta1AnnouncementsResponse and sets the default values.
func NewV1beta1AnnouncementsResponse()(*V1beta1AnnouncementsResponse) {
    m := &V1beta1AnnouncementsResponse{
        V1beta1AnnouncementsGetResponse: *NewV1beta1AnnouncementsGetResponse(),
    }
    return m
}
// CreateV1beta1AnnouncementsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AnnouncementsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AnnouncementsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1AnnouncementsGetResponseable instead.
type V1beta1AnnouncementsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1AnnouncementsGetResponseable
}
