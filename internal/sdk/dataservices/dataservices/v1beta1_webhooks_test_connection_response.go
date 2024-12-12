package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1WebhooksTestConnectionPostResponseable instead.
type V1beta1WebhooksTestConnectionResponse struct {
    V1beta1WebhooksTestConnectionPostResponse
}
// NewV1beta1WebhooksTestConnectionResponse instantiates a new V1beta1WebhooksTestConnectionResponse and sets the default values.
func NewV1beta1WebhooksTestConnectionResponse()(*V1beta1WebhooksTestConnectionResponse) {
    m := &V1beta1WebhooksTestConnectionResponse{
        V1beta1WebhooksTestConnectionPostResponse: *NewV1beta1WebhooksTestConnectionPostResponse(),
    }
    return m
}
// CreateV1beta1WebhooksTestConnectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksTestConnectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksTestConnectionResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1WebhooksTestConnectionPostResponseable instead.
type V1beta1WebhooksTestConnectionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1WebhooksTestConnectionPostResponseable
}
