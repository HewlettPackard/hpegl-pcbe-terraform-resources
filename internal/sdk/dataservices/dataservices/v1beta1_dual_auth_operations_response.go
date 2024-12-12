package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DualAuthOperationsGetResponseable instead.
type V1beta1DualAuthOperationsResponse struct {
    V1beta1DualAuthOperationsGetResponse
}
// NewV1beta1DualAuthOperationsResponse instantiates a new V1beta1DualAuthOperationsResponse and sets the default values.
func NewV1beta1DualAuthOperationsResponse()(*V1beta1DualAuthOperationsResponse) {
    m := &V1beta1DualAuthOperationsResponse{
        V1beta1DualAuthOperationsGetResponse: *NewV1beta1DualAuthOperationsGetResponse(),
    }
    return m
}
// CreateV1beta1DualAuthOperationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DualAuthOperationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DualAuthOperationsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DualAuthOperationsGetResponseable instead.
type V1beta1DualAuthOperationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DualAuthOperationsGetResponseable
}
