package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemHypervisorResponse represents a single instance of a hypervisor manager (Vmware vCenter)
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemHypervisorGetResponseable instead.
type V1beta1HypervisorManagersItemHypervisorResponse struct {
    V1beta1HypervisorManagersItemHypervisorGetResponse
}
// NewV1beta1HypervisorManagersItemHypervisorResponse instantiates a new V1beta1HypervisorManagersItemHypervisorResponse and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorResponse()(*V1beta1HypervisorManagersItemHypervisorResponse) {
    m := &V1beta1HypervisorManagersItemHypervisorResponse{
        V1beta1HypervisorManagersItemHypervisorGetResponse: *NewV1beta1HypervisorManagersItemHypervisorGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemHypervisorGetResponseable instead.
type V1beta1HypervisorManagersItemHypervisorResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemHypervisorGetResponseable
}
