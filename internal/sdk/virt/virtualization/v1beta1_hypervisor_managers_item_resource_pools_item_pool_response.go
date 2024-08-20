package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemResourcePoolsItemPoolResponse represents a hypervisor resource pool.
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable instead.
type V1beta1HypervisorManagersItemResourcePoolsItemPoolResponse struct {
    V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponse
}
// NewV1beta1HypervisorManagersItemResourcePoolsItemPoolResponse instantiates a new V1beta1HypervisorManagersItemResourcePoolsItemPoolResponse and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsItemPoolResponse()(*V1beta1HypervisorManagersItemResourcePoolsItemPoolResponse) {
    m := &V1beta1HypervisorManagersItemResourcePoolsItemPoolResponse{
        V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponse: *NewV1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemResourcePoolsItemPoolResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemResourcePoolsItemPoolResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemResourcePoolsItemPoolResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable instead.
type V1beta1HypervisorManagersItemResourcePoolsItemPoolResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable
}
