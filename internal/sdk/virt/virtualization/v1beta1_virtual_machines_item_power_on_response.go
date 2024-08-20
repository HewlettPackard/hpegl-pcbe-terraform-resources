package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemPowerOnPostResponseable instead.
type V1beta1VirtualMachinesItemPowerOnResponse struct {
    V1beta1VirtualMachinesItemPowerOnPostResponse
}
// NewV1beta1VirtualMachinesItemPowerOnResponse instantiates a new V1beta1VirtualMachinesItemPowerOnResponse and sets the default values.
func NewV1beta1VirtualMachinesItemPowerOnResponse()(*V1beta1VirtualMachinesItemPowerOnResponse) {
    m := &V1beta1VirtualMachinesItemPowerOnResponse{
        V1beta1VirtualMachinesItemPowerOnPostResponse: *NewV1beta1VirtualMachinesItemPowerOnPostResponse(),
    }
    return m
}
// CreateV1beta1VirtualMachinesItemPowerOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemPowerOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemPowerOnResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1VirtualMachinesItemPowerOnPostResponseable instead.
type V1beta1VirtualMachinesItemPowerOnResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1VirtualMachinesItemPowerOnPostResponseable
}
