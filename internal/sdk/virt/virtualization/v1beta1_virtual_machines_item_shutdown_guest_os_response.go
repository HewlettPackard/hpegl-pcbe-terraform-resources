package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemShutdownGuestOsResponse struct {
    V1beta1VirtualMachinesItemShutdownGuestOsPostResponse
}
// NewV1beta1VirtualMachinesItemShutdownGuestOsResponse instantiates a new V1beta1VirtualMachinesItemShutdownGuestOsResponse and sets the default values.
func NewV1beta1VirtualMachinesItemShutdownGuestOsResponse()(*V1beta1VirtualMachinesItemShutdownGuestOsResponse) {
    m := &V1beta1VirtualMachinesItemShutdownGuestOsResponse{
        V1beta1VirtualMachinesItemShutdownGuestOsPostResponse: *NewV1beta1VirtualMachinesItemShutdownGuestOsPostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemShutdownGuestOsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemShutdownGuestOsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemShutdownGuestOsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemShutdownGuestOsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable
}
