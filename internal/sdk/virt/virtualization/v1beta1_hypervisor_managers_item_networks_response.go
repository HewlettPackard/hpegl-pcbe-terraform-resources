package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemNetworksGetResponseable instead.
type V1beta1HypervisorManagersItemNetworksResponse struct {
    V1beta1HypervisorManagersItemNetworksGetResponse
}
// NewV1beta1HypervisorManagersItemNetworksResponse instantiates a new V1beta1HypervisorManagersItemNetworksResponse and sets the default values.
func NewV1beta1HypervisorManagersItemNetworksResponse()(*V1beta1HypervisorManagersItemNetworksResponse) {
    m := &V1beta1HypervisorManagersItemNetworksResponse{
        V1beta1HypervisorManagersItemNetworksGetResponse: *NewV1beta1HypervisorManagersItemNetworksGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemNetworksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemNetworksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemNetworksResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemNetworksGetResponseable instead.
type V1beta1HypervisorManagersItemNetworksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemNetworksGetResponseable
}
