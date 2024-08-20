package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemTagsGetResponseable instead.
type V1beta1HypervisorManagersItemTagsResponse struct {
    V1beta1HypervisorManagersItemTagsGetResponse
}
// NewV1beta1HypervisorManagersItemTagsResponse instantiates a new V1beta1HypervisorManagersItemTagsResponse and sets the default values.
func NewV1beta1HypervisorManagersItemTagsResponse()(*V1beta1HypervisorManagersItemTagsResponse) {
    m := &V1beta1HypervisorManagersItemTagsResponse{
        V1beta1HypervisorManagersItemTagsGetResponse: *NewV1beta1HypervisorManagersItemTagsGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemTagsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemTagsGetResponseable instead.
type V1beta1HypervisorManagersItemTagsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemTagsGetResponseable
}
