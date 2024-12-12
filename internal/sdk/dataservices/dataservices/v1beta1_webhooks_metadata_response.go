package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksMetadataResponse a list of webhook metadata
// Deprecated: This class is obsolete. Use V1beta1WebhooksMetadataGetResponseable instead.
type V1beta1WebhooksMetadataResponse struct {
    V1beta1WebhooksMetadataGetResponse
}
// NewV1beta1WebhooksMetadataResponse instantiates a new V1beta1WebhooksMetadataResponse and sets the default values.
func NewV1beta1WebhooksMetadataResponse()(*V1beta1WebhooksMetadataResponse) {
    m := &V1beta1WebhooksMetadataResponse{
        V1beta1WebhooksMetadataGetResponse: *NewV1beta1WebhooksMetadataGetResponse(),
    }
    return m
}
// CreateV1beta1WebhooksMetadataResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksMetadataResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksMetadataResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1WebhooksMetadataGetResponseable instead.
type V1beta1WebhooksMetadataResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1WebhooksMetadataGetResponseable
}
