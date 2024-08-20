package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorHostsGetResponseable instead.
type V1beta1HypervisorHostsResponse struct {
    V1beta1HypervisorHostsGetResponse
}
// NewV1beta1HypervisorHostsResponse instantiates a new V1beta1HypervisorHostsResponse and sets the default values.
func NewV1beta1HypervisorHostsResponse()(*V1beta1HypervisorHostsResponse) {
    m := &V1beta1HypervisorHostsResponse{
        V1beta1HypervisorHostsGetResponse: *NewV1beta1HypervisorHostsGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorHostsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorHostsGetResponseable instead.
type V1beta1HypervisorHostsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorHostsGetResponseable
}
