package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwarePrechecksPostResponseable instead.
type V1beta1SystemsItemSoftwarePrechecksResponse struct {
    V1beta1SystemsItemSoftwarePrechecksPostResponse
}
// NewV1beta1SystemsItemSoftwarePrechecksResponse instantiates a new V1beta1SystemsItemSoftwarePrechecksResponse and sets the default values.
func NewV1beta1SystemsItemSoftwarePrechecksResponse()(*V1beta1SystemsItemSoftwarePrechecksResponse) {
    m := &V1beta1SystemsItemSoftwarePrechecksResponse{
        V1beta1SystemsItemSoftwarePrechecksPostResponse: *NewV1beta1SystemsItemSoftwarePrechecksPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSoftwarePrechecksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwarePrechecksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwarePrechecksResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwarePrechecksPostResponseable instead.
type V1beta1SystemsItemSoftwarePrechecksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSoftwarePrechecksPostResponseable
}
