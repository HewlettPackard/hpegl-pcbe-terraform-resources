package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksItemWebhooksResponse a Webhook resource
// Deprecated: This class is obsolete. Use V1beta1WebhooksItemWebhooksGetResponseable instead.
type V1beta1WebhooksItemWebhooksResponse struct {
    V1beta1WebhooksItemWebhooksGetResponse
}
// NewV1beta1WebhooksItemWebhooksResponse instantiates a new V1beta1WebhooksItemWebhooksResponse and sets the default values.
func NewV1beta1WebhooksItemWebhooksResponse()(*V1beta1WebhooksItemWebhooksResponse) {
    m := &V1beta1WebhooksItemWebhooksResponse{
        V1beta1WebhooksItemWebhooksGetResponse: *NewV1beta1WebhooksItemWebhooksGetResponse(),
    }
    return m
}
// CreateV1beta1WebhooksItemWebhooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksItemWebhooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksItemWebhooksResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1WebhooksItemWebhooksGetResponseable instead.
type V1beta1WebhooksItemWebhooksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1WebhooksItemWebhooksGetResponseable
}
