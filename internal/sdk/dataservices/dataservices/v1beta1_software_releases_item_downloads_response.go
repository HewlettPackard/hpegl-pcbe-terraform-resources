package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareReleasesItemDownloadsResponse the values needed to download a file within a Software Release.
// Deprecated: This class is obsolete. Use V1beta1SoftwareReleasesItemDownloadsPostResponseable instead.
type V1beta1SoftwareReleasesItemDownloadsResponse struct {
    V1beta1SoftwareReleasesItemDownloadsPostResponse
}
// NewV1beta1SoftwareReleasesItemDownloadsResponse instantiates a new V1beta1SoftwareReleasesItemDownloadsResponse and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadsResponse()(*V1beta1SoftwareReleasesItemDownloadsResponse) {
    m := &V1beta1SoftwareReleasesItemDownloadsResponse{
        V1beta1SoftwareReleasesItemDownloadsPostResponse: *NewV1beta1SoftwareReleasesItemDownloadsPostResponse(),
    }
    return m
}
// CreateV1beta1SoftwareReleasesItemDownloadsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareReleasesItemDownloadsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareReleasesItemDownloadsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SoftwareReleasesItemDownloadsPostResponseable instead.
type V1beta1SoftwareReleasesItemDownloadsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SoftwareReleasesItemDownloadsPostResponseable
}
