package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemCustomizeGuestOsResponse struct {
    V1beta1VirtualMachinesItemCustomizeGuestOsPostResponse
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsResponse instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsResponse and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsResponse()(*V1beta1VirtualMachinesItemCustomizeGuestOsResponse) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsResponse{
        V1beta1VirtualMachinesItemCustomizeGuestOsPostResponse: *NewV1beta1VirtualMachinesItemCustomizeGuestOsPostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable instead.
type V1beta1VirtualMachinesItemCustomizeGuestOsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable
}
