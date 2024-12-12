package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SearchSummariesGetResponseable instead.
type V1beta1SearchSummariesResponse struct {
    V1beta1SearchSummariesGetResponse
}
// NewV1beta1SearchSummariesResponse instantiates a new V1beta1SearchSummariesResponse and sets the default values.
func NewV1beta1SearchSummariesResponse()(*V1beta1SearchSummariesResponse) {
    m := &V1beta1SearchSummariesResponse{
        V1beta1SearchSummariesGetResponse: *NewV1beta1SearchSummariesGetResponse(),
    }
    return m
}
// CreateV1beta1SearchSummariesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SearchSummariesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SearchSummariesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SearchSummariesGetResponseable instead.
type V1beta1SearchSummariesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SearchSummariesGetResponseable
}
