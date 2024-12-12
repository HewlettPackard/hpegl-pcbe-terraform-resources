package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1IssuesMetadataGetResponseable instead.
type V1beta1IssuesMetadataResponse struct {
    V1beta1IssuesMetadataGetResponse
}
// NewV1beta1IssuesMetadataResponse instantiates a new V1beta1IssuesMetadataResponse and sets the default values.
func NewV1beta1IssuesMetadataResponse()(*V1beta1IssuesMetadataResponse) {
    m := &V1beta1IssuesMetadataResponse{
        V1beta1IssuesMetadataGetResponse: *NewV1beta1IssuesMetadataGetResponse(),
    }
    return m
}
// CreateV1beta1IssuesMetadataResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1IssuesMetadataResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1IssuesMetadataResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1IssuesMetadataGetResponseable instead.
type V1beta1IssuesMetadataResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1IssuesMetadataGetResponseable
}
