package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemStorageArraysGetResponseable instead.
type V1beta1SystemsItemStorageArraysResponse struct {
    V1beta1SystemsItemStorageArraysGetResponse
}
// NewV1beta1SystemsItemStorageArraysResponse instantiates a new V1beta1SystemsItemStorageArraysResponse and sets the default values.
func NewV1beta1SystemsItemStorageArraysResponse()(*V1beta1SystemsItemStorageArraysResponse) {
    m := &V1beta1SystemsItemStorageArraysResponse{
        V1beta1SystemsItemStorageArraysGetResponse: *NewV1beta1SystemsItemStorageArraysGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemStorageArraysResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageArraysResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageArraysResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemStorageArraysGetResponseable instead.
type V1beta1SystemsItemStorageArraysResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemStorageArraysGetResponseable
}
