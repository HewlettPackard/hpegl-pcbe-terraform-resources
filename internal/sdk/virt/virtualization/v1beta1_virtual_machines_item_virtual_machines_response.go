package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable instead.
type V1beta1VirtualMachinesItemVirtualMachinesResponse struct {
    V1beta1VirtualMachinesItemVirtualMachinesDeleteResponse
}
// NewV1beta1VirtualMachinesItemVirtualMachinesResponse instantiates a new V1beta1VirtualMachinesItemVirtualMachinesResponse and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesResponse()(*V1beta1VirtualMachinesItemVirtualMachinesResponse) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesResponse{
        V1beta1VirtualMachinesItemVirtualMachinesDeleteResponse: *NewV1beta1VirtualMachinesItemVirtualMachinesDeleteResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable instead.
type V1beta1VirtualMachinesItemVirtualMachinesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable
}
