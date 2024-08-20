package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemNetworksItemNetworkResponse represents a hypervisor network resource
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable instead.
type V1beta1HypervisorManagersItemNetworksItemNetworkResponse struct {
    V1beta1HypervisorManagersItemNetworksItemNetworkGetResponse
}
// NewV1beta1HypervisorManagersItemNetworksItemNetworkResponse instantiates a new V1beta1HypervisorManagersItemNetworksItemNetworkResponse and sets the default values.
func NewV1beta1HypervisorManagersItemNetworksItemNetworkResponse()(*V1beta1HypervisorManagersItemNetworksItemNetworkResponse) {
    m := &V1beta1HypervisorManagersItemNetworksItemNetworkResponse{
        V1beta1HypervisorManagersItemNetworksItemNetworkGetResponse: *NewV1beta1HypervisorManagersItemNetworksItemNetworkGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemNetworksItemNetworkResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemNetworksItemNetworkResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemNetworksItemNetworkResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable instead.
type V1beta1HypervisorManagersItemNetworksItemNetworkResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable
}
