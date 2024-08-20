package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesGetResponseable instead.
type V1beta1VirtualMachinesResponse struct {
    V1beta1VirtualMachinesGetResponse
}
// NewV1beta1VirtualMachinesResponse instantiates a new V1beta1VirtualMachinesResponse and sets the default values.
func NewV1beta1VirtualMachinesResponse()(*V1beta1VirtualMachinesResponse) {
    m := &V1beta1VirtualMachinesResponse{
        V1beta1VirtualMachinesGetResponse: *NewV1beta1VirtualMachinesGetResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesGetResponseable instead.
type V1beta1VirtualMachinesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesGetResponseable
}
