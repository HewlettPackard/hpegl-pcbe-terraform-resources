package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspVolumesUpdateCspTagsPostResponseable instead.
type V1beta1CspVolumesUpdateCspTagsResponse struct {
    V1beta1CspVolumesUpdateCspTagsPostResponse
}
// NewV1beta1CspVolumesUpdateCspTagsResponse instantiates a new V1beta1CspVolumesUpdateCspTagsResponse and sets the default values.
func NewV1beta1CspVolumesUpdateCspTagsResponse()(*V1beta1CspVolumesUpdateCspTagsResponse) {
    m := &V1beta1CspVolumesUpdateCspTagsResponse{
        V1beta1CspVolumesUpdateCspTagsPostResponse: *NewV1beta1CspVolumesUpdateCspTagsPostResponse(),
    }
    return m
}
// CreateV1beta1CspVolumesUpdateCspTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesUpdateCspTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesUpdateCspTagsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspVolumesUpdateCspTagsPostResponseable instead.
type V1beta1CspVolumesUpdateCspTagsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspVolumesUpdateCspTagsPostResponseable
}
