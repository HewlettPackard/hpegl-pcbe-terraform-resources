package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareReleasesItemSoftwareReleasesResponse a release of software.
// Deprecated: This class is obsolete. Use V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable instead.
type V1beta1SoftwareReleasesItemSoftwareReleasesResponse struct {
    V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse
}
// NewV1beta1SoftwareReleasesItemSoftwareReleasesResponse instantiates a new V1beta1SoftwareReleasesItemSoftwareReleasesResponse and sets the default values.
func NewV1beta1SoftwareReleasesItemSoftwareReleasesResponse()(*V1beta1SoftwareReleasesItemSoftwareReleasesResponse) {
    m := &V1beta1SoftwareReleasesItemSoftwareReleasesResponse{
        V1beta1SoftwareReleasesItemSoftwareReleasesGetResponse: *NewV1beta1SoftwareReleasesItemSoftwareReleasesGetResponse(),
    }
    return m
}
// CreateV1beta1SoftwareReleasesItemSoftwareReleasesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareReleasesItemSoftwareReleasesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareReleasesItemSoftwareReleasesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable instead.
type V1beta1SoftwareReleasesItemSoftwareReleasesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable
}
