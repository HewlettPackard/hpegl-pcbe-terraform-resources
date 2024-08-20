package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemTagsItemTagResponse represents a hypervisor tag resource
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemTagsItemTagGetResponseable instead.
type V1beta1HypervisorManagersItemTagsItemTagResponse struct {
    V1beta1HypervisorManagersItemTagsItemTagGetResponse
}
// NewV1beta1HypervisorManagersItemTagsItemTagResponse instantiates a new V1beta1HypervisorManagersItemTagsItemTagResponse and sets the default values.
func NewV1beta1HypervisorManagersItemTagsItemTagResponse()(*V1beta1HypervisorManagersItemTagsItemTagResponse) {
    m := &V1beta1HypervisorManagersItemTagsItemTagResponse{
        V1beta1HypervisorManagersItemTagsItemTagGetResponse: *NewV1beta1HypervisorManagersItemTagsItemTagGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemTagsItemTagResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemTagsItemTagResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemTagsItemTagResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemTagsItemTagGetResponseable instead.
type V1beta1HypervisorManagersItemTagsItemTagResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemTagsItemTagGetResponseable
}
