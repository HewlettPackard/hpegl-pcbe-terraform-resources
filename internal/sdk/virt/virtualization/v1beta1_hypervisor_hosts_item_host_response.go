package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsItemHostResponse represents a single instance of a hypervisor host (Vmware ESXi)
// Deprecated: This class is obsolete. Use V1beta1HypervisorHostsItemHostGetResponseable instead.
type V1beta1HypervisorHostsItemHostResponse struct {
    V1beta1HypervisorHostsItemHostGetResponse
}
// NewV1beta1HypervisorHostsItemHostResponse instantiates a new V1beta1HypervisorHostsItemHostResponse and sets the default values.
func NewV1beta1HypervisorHostsItemHostResponse()(*V1beta1HypervisorHostsItemHostResponse) {
    m := &V1beta1HypervisorHostsItemHostResponse{
        V1beta1HypervisorHostsItemHostGetResponse: *NewV1beta1HypervisorHostsItemHostGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorHostsItemHostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsItemHostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsItemHostResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorHostsItemHostGetResponseable instead.
type V1beta1HypervisorHostsItemHostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorHostsItemHostGetResponseable
}
