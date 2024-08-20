package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesGetResponseable instead.
type V1beta1CspMachineInstancesResponse struct {
    V1beta1CspMachineInstancesGetResponse
}
// NewV1beta1CspMachineInstancesResponse instantiates a new V1beta1CspMachineInstancesResponse and sets the default values.
func NewV1beta1CspMachineInstancesResponse()(*V1beta1CspMachineInstancesResponse) {
    m := &V1beta1CspMachineInstancesResponse{
        V1beta1CspMachineInstancesGetResponse: *NewV1beta1CspMachineInstancesGetResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesGetResponseable instead.
type V1beta1CspMachineInstancesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesGetResponseable
}
