package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksResponse a list of webhooks
// Deprecated: This class is obsolete. Use V1beta1WebhooksGetResponseable instead.
type V1beta1WebhooksResponse struct {
    V1beta1WebhooksGetResponse
}
// NewV1beta1WebhooksResponse instantiates a new V1beta1WebhooksResponse and sets the default values.
func NewV1beta1WebhooksResponse()(*V1beta1WebhooksResponse) {
    m := &V1beta1WebhooksResponse{
        V1beta1WebhooksGetResponse: *NewV1beta1WebhooksGetResponse(),
    }
    return m
}
// CreateV1beta1WebhooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1WebhooksGetResponseable instead.
type V1beta1WebhooksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1WebhooksGetResponseable
}
