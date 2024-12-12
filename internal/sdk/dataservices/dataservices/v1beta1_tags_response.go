package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1TagsGetResponseable instead.
type V1beta1TagsResponse struct {
    V1beta1TagsGetResponse
}
// NewV1beta1TagsResponse instantiates a new V1beta1TagsResponse and sets the default values.
func NewV1beta1TagsResponse()(*V1beta1TagsResponse) {
    m := &V1beta1TagsResponse{
        V1beta1TagsGetResponse: *NewV1beta1TagsGetResponse(),
    }
    return m
}
// CreateV1beta1TagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1TagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1TagsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1TagsGetResponseable instead.
type V1beta1TagsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1TagsGetResponseable
}
