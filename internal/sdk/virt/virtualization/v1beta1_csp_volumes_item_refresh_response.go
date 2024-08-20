package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspVolumesItemRefreshPostResponseable instead.
type V1beta1CspVolumesItemRefreshResponse struct {
    V1beta1CspVolumesItemRefreshPostResponse
}
// NewV1beta1CspVolumesItemRefreshResponse instantiates a new V1beta1CspVolumesItemRefreshResponse and sets the default values.
func NewV1beta1CspVolumesItemRefreshResponse()(*V1beta1CspVolumesItemRefreshResponse) {
    m := &V1beta1CspVolumesItemRefreshResponse{
        V1beta1CspVolumesItemRefreshPostResponse: *NewV1beta1CspVolumesItemRefreshPostResponse(),
    }
    return m
}
// CreateV1beta1CspVolumesItemRefreshResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspVolumesItemRefreshResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspVolumesItemRefreshResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspVolumesItemRefreshPostResponseable instead.
type V1beta1CspVolumesItemRefreshResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspVolumesItemRefreshPostResponseable
}
