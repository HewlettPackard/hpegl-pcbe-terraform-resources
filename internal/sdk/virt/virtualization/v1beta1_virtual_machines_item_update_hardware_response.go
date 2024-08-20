package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemUpdateHardwarePostResponseable instead.
type V1beta1VirtualMachinesItemUpdateHardwareResponse struct {
    V1beta1VirtualMachinesItemUpdateHardwarePostResponse
}
// NewV1beta1VirtualMachinesItemUpdateHardwareResponse instantiates a new V1beta1VirtualMachinesItemUpdateHardwareResponse and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwareResponse()(*V1beta1VirtualMachinesItemUpdateHardwareResponse) {
    m := &V1beta1VirtualMachinesItemUpdateHardwareResponse{
        V1beta1VirtualMachinesItemUpdateHardwarePostResponse: *NewV1beta1VirtualMachinesItemUpdateHardwarePostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwareResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwareResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwareResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemUpdateHardwarePostResponseable instead.
type V1beta1VirtualMachinesItemUpdateHardwareResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemUpdateHardwarePostResponseable
}
