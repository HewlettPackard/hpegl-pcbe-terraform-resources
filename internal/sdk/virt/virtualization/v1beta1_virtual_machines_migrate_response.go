package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesMigratePostResponseable instead.
type V1beta1VirtualMachinesMigrateResponse struct {
    V1beta1VirtualMachinesMigratePostResponse
}
// NewV1beta1VirtualMachinesMigrateResponse instantiates a new V1beta1VirtualMachinesMigrateResponse and sets the default values.
func NewV1beta1VirtualMachinesMigrateResponse()(*V1beta1VirtualMachinesMigrateResponse) {
    m := &V1beta1VirtualMachinesMigrateResponse{
        V1beta1VirtualMachinesMigratePostResponse: *NewV1beta1VirtualMachinesMigratePostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesMigrateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesMigrateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesMigrateResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesMigratePostResponseable instead.
type V1beta1VirtualMachinesMigrateResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesMigratePostResponseable
}
