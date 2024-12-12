package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1EmailSubscriptionsItemEmailSubscriptionsResponse common properties included in all resource models.
// Deprecated: This class is obsolete. Use V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable instead.
type V1beta1EmailSubscriptionsItemEmailSubscriptionsResponse struct {
    V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponse
}
// NewV1beta1EmailSubscriptionsItemEmailSubscriptionsResponse instantiates a new V1beta1EmailSubscriptionsItemEmailSubscriptionsResponse and sets the default values.
func NewV1beta1EmailSubscriptionsItemEmailSubscriptionsResponse()(*V1beta1EmailSubscriptionsItemEmailSubscriptionsResponse) {
    m := &V1beta1EmailSubscriptionsItemEmailSubscriptionsResponse{
        V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponse: *NewV1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponse(),
    }
    return m
}
// CreateV1beta1EmailSubscriptionsItemEmailSubscriptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1EmailSubscriptionsItemEmailSubscriptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1EmailSubscriptionsItemEmailSubscriptionsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable instead.
type V1beta1EmailSubscriptionsItemEmailSubscriptionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable
}
