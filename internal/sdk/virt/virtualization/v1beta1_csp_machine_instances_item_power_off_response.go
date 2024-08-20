package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemPowerOffPostResponseable instead.
type V1beta1CspMachineInstancesItemPowerOffResponse struct {
    V1beta1CspMachineInstancesItemPowerOffPostResponse
}
// NewV1beta1CspMachineInstancesItemPowerOffResponse instantiates a new V1beta1CspMachineInstancesItemPowerOffResponse and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOffResponse()(*V1beta1CspMachineInstancesItemPowerOffResponse) {
    m := &V1beta1CspMachineInstancesItemPowerOffResponse{
        V1beta1CspMachineInstancesItemPowerOffPostResponse: *NewV1beta1CspMachineInstancesItemPowerOffPostResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesItemPowerOffResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemPowerOffResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemPowerOffResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemPowerOffPostResponseable instead.
type V1beta1CspMachineInstancesItemPowerOffResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesItemPowerOffPostResponseable
}
