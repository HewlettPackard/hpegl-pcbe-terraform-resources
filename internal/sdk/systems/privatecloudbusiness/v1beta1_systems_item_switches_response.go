package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSwitchesGetResponseable instead.
type V1beta1SystemsItemSwitchesResponse struct {
    V1beta1SystemsItemSwitchesGetResponse
}
// NewV1beta1SystemsItemSwitchesResponse instantiates a new V1beta1SystemsItemSwitchesResponse and sets the default values.
func NewV1beta1SystemsItemSwitchesResponse()(*V1beta1SystemsItemSwitchesResponse) {
    m := &V1beta1SystemsItemSwitchesResponse{
        V1beta1SystemsItemSwitchesGetResponse: *NewV1beta1SystemsItemSwitchesGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSwitchesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSwitchesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSwitchesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSwitchesGetResponseable instead.
type V1beta1SystemsItemSwitchesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSwitchesGetResponseable
}
