package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SearchMetadataGetResponseable instead.
type V1beta1SearchMetadataResponse struct {
    V1beta1SearchMetadataGetResponse
}
// NewV1beta1SearchMetadataResponse instantiates a new V1beta1SearchMetadataResponse and sets the default values.
func NewV1beta1SearchMetadataResponse()(*V1beta1SearchMetadataResponse) {
    m := &V1beta1SearchMetadataResponse{
        V1beta1SearchMetadataGetResponse: *NewV1beta1SearchMetadataGetResponse(),
    }
    return m
}
// CreateV1beta1SearchMetadataResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchMetadataResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchMetadataResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SearchMetadataGetResponseable instead.
type V1beta1SearchMetadataResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SearchMetadataGetResponseable
}
