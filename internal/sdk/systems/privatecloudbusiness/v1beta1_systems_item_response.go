package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemResponse details of the system with select information.
// Deprecated: This class is obsolete. Use V1beta1SystemsItemGetResponseable instead.
type V1beta1SystemsItemResponse struct {
    V1beta1SystemsItemGetResponse
}
// NewV1beta1SystemsItemResponse instantiates a new V1beta1SystemsItemResponse and sets the default values.
func NewV1beta1SystemsItemResponse()(*V1beta1SystemsItemResponse) {
    m := &V1beta1SystemsItemResponse{
        V1beta1SystemsItemGetResponse: *NewV1beta1SystemsItemGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemGetResponseable instead.
type V1beta1SystemsItemResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemGetResponseable
}
