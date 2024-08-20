package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesUpdateCspTagsPostResponseable instead.
type V1beta1CspMachineInstancesUpdateCspTagsResponse struct {
    V1beta1CspMachineInstancesUpdateCspTagsPostResponse
}
// NewV1beta1CspMachineInstancesUpdateCspTagsResponse instantiates a new V1beta1CspMachineInstancesUpdateCspTagsResponse and sets the default values.
func NewV1beta1CspMachineInstancesUpdateCspTagsResponse()(*V1beta1CspMachineInstancesUpdateCspTagsResponse) {
    m := &V1beta1CspMachineInstancesUpdateCspTagsResponse{
        V1beta1CspMachineInstancesUpdateCspTagsPostResponse: *NewV1beta1CspMachineInstancesUpdateCspTagsPostResponse(),
    }
    return m
}
// CreateV1beta1CspMachineInstancesUpdateCspTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesUpdateCspTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesUpdateCspTagsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineInstancesUpdateCspTagsPostResponseable instead.
type V1beta1CspMachineInstancesUpdateCspTagsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineInstancesUpdateCspTagsPostResponseable
}
