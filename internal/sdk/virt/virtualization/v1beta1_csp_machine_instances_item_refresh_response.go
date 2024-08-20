package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemRefreshPostResponseable instead.
type V1beta1CspMachineInstancesItemRefreshResponse struct {
    V1beta1CspMachineInstancesItemRefreshPostResponse
}
// NewV1beta1CspMachineInstancesItemRefreshResponse instantiates a new V1beta1CspMachineInstancesItemRefreshResponse and sets the default values.
func NewV1beta1CspMachineInstancesItemRefreshResponse()(*V1beta1CspMachineInstancesItemRefreshResponse) {
    m := &V1beta1CspMachineInstancesItemRefreshResponse{
        V1beta1CspMachineInstancesItemRefreshPostResponse: *NewV1beta1CspMachineInstancesItemRefreshPostResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesItemRefreshResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemRefreshResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemRefreshResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemRefreshPostResponseable instead.
type V1beta1CspMachineInstancesItemRefreshResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesItemRefreshPostResponseable
}
