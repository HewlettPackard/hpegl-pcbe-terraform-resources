package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1EmailSubscriptionsGetResponseable instead.
type V1beta1EmailSubscriptionsResponse struct {
    V1beta1EmailSubscriptionsGetResponse
}
// NewV1beta1EmailSubscriptionsResponse instantiates a new V1beta1EmailSubscriptionsResponse and sets the default values.
func NewV1beta1EmailSubscriptionsResponse()(*V1beta1EmailSubscriptionsResponse) {
    m := &V1beta1EmailSubscriptionsResponse{
        V1beta1EmailSubscriptionsGetResponse: *NewV1beta1EmailSubscriptionsGetResponse(),
    }
    return m
}
// CreateV1beta1EmailSubscriptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1EmailSubscriptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1EmailSubscriptionsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1EmailSubscriptionsGetResponseable instead.
type V1beta1EmailSubscriptionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1EmailSubscriptionsGetResponseable
}
