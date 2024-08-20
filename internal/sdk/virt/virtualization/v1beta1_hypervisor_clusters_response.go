package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorClustersGetResponseable instead.
type V1beta1HypervisorClustersResponse struct {
    V1beta1HypervisorClustersGetResponse
}
// NewV1beta1HypervisorClustersResponse instantiates a new V1beta1HypervisorClustersResponse and sets the default values.
func NewV1beta1HypervisorClustersResponse()(*V1beta1HypervisorClustersResponse) {
    m := &V1beta1HypervisorClustersResponse{
        V1beta1HypervisorClustersGetResponse: *NewV1beta1HypervisorClustersGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorClustersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorClustersGetResponseable instead.
type V1beta1HypervisorClustersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorClustersGetResponseable
}
