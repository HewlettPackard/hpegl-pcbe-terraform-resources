package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspVolumesGetResponseable instead.
type V1beta1CspVolumesResponse struct {
    V1beta1CspVolumesGetResponse
}
// NewV1beta1CspVolumesResponse instantiates a new V1beta1CspVolumesResponse and sets the default values.
func NewV1beta1CspVolumesResponse()(*V1beta1CspVolumesResponse) {
    m := &V1beta1CspVolumesResponse{
        V1beta1CspVolumesGetResponse: *NewV1beta1CspVolumesGetResponse(),
    }
    return m
}
// CreateV1beta1CspVolumesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspVolumesGetResponseable instead.
type V1beta1CspVolumesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspVolumesGetResponseable
}
