package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable instead.
type V1beta1HypervisorManagersItemHypervisorLibraryImagesResponse struct {
    V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponse
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesResponse instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesResponse and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesResponse()(*V1beta1HypervisorManagersItemHypervisorLibraryImagesResponse) {
    m := &V1beta1HypervisorManagersItemHypervisorLibraryImagesResponse{
        V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponse: *NewV1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable instead.
type V1beta1HypervisorManagersItemHypervisorLibraryImagesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemHypervisorLibraryImagesGetResponseable
}