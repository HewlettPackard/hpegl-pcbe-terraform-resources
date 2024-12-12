package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1OneTimeTokensPostResponseable instead.
type V1beta1OneTimeTokensResponse struct {
    V1beta1OneTimeTokensPostResponse
}
// NewV1beta1OneTimeTokensResponse instantiates a new V1beta1OneTimeTokensResponse and sets the default values.
func NewV1beta1OneTimeTokensResponse()(*V1beta1OneTimeTokensResponse) {
    m := &V1beta1OneTimeTokensResponse{
        V1beta1OneTimeTokensPostResponse: *NewV1beta1OneTimeTokensPostResponse(),
    }
    return m
}
// CreateV1beta1OneTimeTokensResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1OneTimeTokensResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1OneTimeTokensResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1OneTimeTokensPostResponseable instead.
type V1beta1OneTimeTokensResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1OneTimeTokensPostResponseable
}
