package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SearchGetResponseable instead.
type V1beta1SearchResponse struct {
    V1beta1SearchGetResponse
}
// NewV1beta1SearchResponse instantiates a new V1beta1SearchResponse and sets the default values.
func NewV1beta1SearchResponse()(*V1beta1SearchResponse) {
    m := &V1beta1SearchResponse{
        V1beta1SearchGetResponse: *NewV1beta1SearchGetResponse(),
    }
    return m
}
// CreateV1beta1SearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SearchGetResponseable instead.
type V1beta1SearchResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SearchGetResponseable
}
