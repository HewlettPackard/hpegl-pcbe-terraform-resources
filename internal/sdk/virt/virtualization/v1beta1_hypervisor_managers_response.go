package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersGetResponseable instead.
type V1beta1HypervisorManagersResponse struct {
    V1beta1HypervisorManagersGetResponse
}
// NewV1beta1HypervisorManagersResponse instantiates a new V1beta1HypervisorManagersResponse and sets the default values.
func NewV1beta1HypervisorManagersResponse()(*V1beta1HypervisorManagersResponse) {
    m := &V1beta1HypervisorManagersResponse{
        V1beta1HypervisorManagersGetResponse: *NewV1beta1HypervisorManagersGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersGetResponseable instead.
type V1beta1HypervisorManagersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersGetResponseable
}
