package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemRestartGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemRestartGuestOsResponse struct {
    V1beta1VirtualMachinesItemRestartGuestOsPostResponse
}
// NewV1beta1VirtualMachinesItemRestartGuestOsResponse instantiates a new V1beta1VirtualMachinesItemRestartGuestOsResponse and sets the default values.
func NewV1beta1VirtualMachinesItemRestartGuestOsResponse()(*V1beta1VirtualMachinesItemRestartGuestOsResponse) {
    m := &V1beta1VirtualMachinesItemRestartGuestOsResponse{
        V1beta1VirtualMachinesItemRestartGuestOsPostResponse: *NewV1beta1VirtualMachinesItemRestartGuestOsPostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemRestartGuestOsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemRestartGuestOsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemRestartGuestOsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemRestartGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemRestartGuestOsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemRestartGuestOsPostResponseable
}
