package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemPowerOnPostResponseable instead.
type V1beta1CspMachineInstancesItemPowerOnResponse struct {
    V1beta1CspMachineInstancesItemPowerOnPostResponse
}
// NewV1beta1CspMachineInstancesItemPowerOnResponse instantiates a new V1beta1CspMachineInstancesItemPowerOnResponse and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOnResponse()(*V1beta1CspMachineInstancesItemPowerOnResponse) {
    m := &V1beta1CspMachineInstancesItemPowerOnResponse{
        V1beta1CspMachineInstancesItemPowerOnPostResponse: *NewV1beta1CspMachineInstancesItemPowerOnPostResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesItemPowerOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemPowerOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemPowerOnResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemPowerOnPostResponseable instead.
type V1beta1CspMachineInstancesItemPowerOnResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesItemPowerOnPostResponseable
}
