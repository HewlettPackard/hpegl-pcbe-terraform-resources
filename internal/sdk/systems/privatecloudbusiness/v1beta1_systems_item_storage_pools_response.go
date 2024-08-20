package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemStoragePoolsGetResponseable instead.
type V1beta1SystemsItemStoragePoolsResponse struct {
    V1beta1SystemsItemStoragePoolsGetResponse
}
// NewV1beta1SystemsItemStoragePoolsResponse instantiates a new V1beta1SystemsItemStoragePoolsResponse and sets the default values.
func NewV1beta1SystemsItemStoragePoolsResponse()(*V1beta1SystemsItemStoragePoolsResponse) {
    m := &V1beta1SystemsItemStoragePoolsResponse{
        V1beta1SystemsItemStoragePoolsGetResponse: *NewV1beta1SystemsItemStoragePoolsGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemStoragePoolsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStoragePoolsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStoragePoolsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemStoragePoolsGetResponseable instead.
type V1beta1SystemsItemStoragePoolsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemStoragePoolsGetResponseable
}
