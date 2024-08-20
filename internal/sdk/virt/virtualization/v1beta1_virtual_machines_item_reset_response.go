package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemResetPostResponseable instead.
type V1beta1VirtualMachinesItemResetResponse struct {
    V1beta1VirtualMachinesItemResetPostResponse
}
// NewV1beta1VirtualMachinesItemResetResponse instantiates a new V1beta1VirtualMachinesItemResetResponse and sets the default values.
func NewV1beta1VirtualMachinesItemResetResponse()(*V1beta1VirtualMachinesItemResetResponse) {
    m := &V1beta1VirtualMachinesItemResetResponse{
        V1beta1VirtualMachinesItemResetPostResponse: *NewV1beta1VirtualMachinesItemResetPostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemResetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemResetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemResetResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemResetPostResponseable instead.
type V1beta1VirtualMachinesItemResetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemResetPostResponseable
}
