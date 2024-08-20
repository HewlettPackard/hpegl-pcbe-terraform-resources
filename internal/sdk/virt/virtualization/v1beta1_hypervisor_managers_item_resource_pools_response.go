package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemResourcePoolsGetResponseable instead.
type V1beta1HypervisorManagersItemResourcePoolsResponse struct {
    V1beta1HypervisorManagersItemResourcePoolsGetResponse
}
// NewV1beta1HypervisorManagersItemResourcePoolsResponse instantiates a new V1beta1HypervisorManagersItemResourcePoolsResponse and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsResponse()(*V1beta1HypervisorManagersItemResourcePoolsResponse) {
    m := &V1beta1HypervisorManagersItemResourcePoolsResponse{
        V1beta1HypervisorManagersItemResourcePoolsGetResponse: *NewV1beta1HypervisorManagersItemResourcePoolsGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemResourcePoolsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemResourcePoolsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemResourcePoolsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemResourcePoolsGetResponseable instead.
type V1beta1HypervisorManagersItemResourcePoolsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemResourcePoolsGetResponseable
}
