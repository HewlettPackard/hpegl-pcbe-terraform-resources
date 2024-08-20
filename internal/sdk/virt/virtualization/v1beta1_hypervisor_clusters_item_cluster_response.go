package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorClustersItemClusterResponse represents a single instance of a hypervisor cluster (Vmware ESXi Cluster)
// Deprecated: This class is obsolete. Use V1beta1HypervisorClustersItemClusterGetResponseable instead.
type V1beta1HypervisorClustersItemClusterResponse struct {
    V1beta1HypervisorClustersItemClusterGetResponse
}
// NewV1beta1HypervisorClustersItemClusterResponse instantiates a new V1beta1HypervisorClustersItemClusterResponse and sets the default values.
func NewV1beta1HypervisorClustersItemClusterResponse()(*V1beta1HypervisorClustersItemClusterResponse) {
    m := &V1beta1HypervisorClustersItemClusterResponse{
        V1beta1HypervisorClustersItemClusterGetResponse: *NewV1beta1HypervisorClustersItemClusterGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorClustersItemClusterResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersItemClusterResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersItemClusterResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorClustersItemClusterGetResponseable instead.
type V1beta1HypervisorClustersItemClusterResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorClustersItemClusterGetResponseable
}
