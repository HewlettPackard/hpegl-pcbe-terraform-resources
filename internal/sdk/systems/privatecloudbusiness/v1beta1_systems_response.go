package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsGetResponseable instead.
type V1beta1SystemsResponse struct {
    V1beta1SystemsGetResponse
}
// NewV1beta1SystemsResponse instantiates a new V1beta1SystemsResponse and sets the default values.
func NewV1beta1SystemsResponse()(*V1beta1SystemsResponse) {
    m := &V1beta1SystemsResponse{
        V1beta1SystemsGetResponse: *NewV1beta1SystemsGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsGetResponseable instead.
type V1beta1SystemsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsGetResponseable
}
