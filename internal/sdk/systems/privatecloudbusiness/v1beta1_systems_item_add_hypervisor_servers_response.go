package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemAddHypervisorServersPostResponseable instead.
type V1beta1SystemsItemAddHypervisorServersResponse struct {
    V1beta1SystemsItemAddHypervisorServersPostResponse
}
// NewV1beta1SystemsItemAddHypervisorServersResponse instantiates a new V1beta1SystemsItemAddHypervisorServersResponse and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersResponse()(*V1beta1SystemsItemAddHypervisorServersResponse) {
    m := &V1beta1SystemsItemAddHypervisorServersResponse{
        V1beta1SystemsItemAddHypervisorServersPostResponse: *NewV1beta1SystemsItemAddHypervisorServersPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemAddHypervisorServersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorServersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorServersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemAddHypervisorServersPostResponseable instead.
type V1beta1SystemsItemAddHypervisorServersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemAddHypervisorServersPostResponseable
}
