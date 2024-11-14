package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemRemoveHypervisorServersPostResponseable instead.
type V1beta1SystemsItemRemoveHypervisorServersResponse struct {
    V1beta1SystemsItemRemoveHypervisorServersPostResponse
}
// NewV1beta1SystemsItemRemoveHypervisorServersResponse instantiates a new V1beta1SystemsItemRemoveHypervisorServersResponse and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorServersResponse()(*V1beta1SystemsItemRemoveHypervisorServersResponse) {
    m := &V1beta1SystemsItemRemoveHypervisorServersResponse{
        V1beta1SystemsItemRemoveHypervisorServersPostResponse: *NewV1beta1SystemsItemRemoveHypervisorServersPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemRemoveHypervisorServersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemRemoveHypervisorServersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemRemoveHypervisorServersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemRemoveHypervisorServersPostResponseable instead.
type V1beta1SystemsItemRemoveHypervisorServersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemRemoveHypervisorServersPostResponseable
}
