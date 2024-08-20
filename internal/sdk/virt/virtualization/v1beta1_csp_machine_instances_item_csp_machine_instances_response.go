package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable instead.
type V1beta1CspMachineInstancesItemCspMachineInstancesResponse struct {
    V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponse
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesResponse instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesResponse and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesResponse()(*V1beta1CspMachineInstancesItemCspMachineInstancesResponse) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesResponse{
        V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponse: *NewV1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable instead.
type V1beta1CspMachineInstancesItemCspMachineInstancesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable
}
