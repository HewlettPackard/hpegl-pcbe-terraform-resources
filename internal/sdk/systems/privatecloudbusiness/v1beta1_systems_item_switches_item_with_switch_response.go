package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemSwitchesItemWithSwitchResponse details of the Switch with select information.
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable instead.
type V1beta1SystemsItemSwitchesItemWithSwitchResponse struct {
    V1beta1SystemsItemSwitchesItemWithSwitchGetResponse
}
// NewV1beta1SystemsItemSwitchesItemWithSwitchResponse instantiates a new V1beta1SystemsItemSwitchesItemWithSwitchResponse and sets the default values.
func NewV1beta1SystemsItemSwitchesItemWithSwitchResponse()(*V1beta1SystemsItemSwitchesItemWithSwitchResponse) {
    m := &V1beta1SystemsItemSwitchesItemWithSwitchResponse{
        V1beta1SystemsItemSwitchesItemWithSwitchGetResponse: *NewV1beta1SystemsItemSwitchesItemWithSwitchGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSwitchesItemWithSwitchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSwitchesItemWithSwitchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSwitchesItemWithSwitchResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable instead.
type V1beta1SystemsItemSwitchesItemWithSwitchResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable
}
