package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1IssuesGetResponseable instead.
type V1beta1IssuesResponse struct {
    V1beta1IssuesGetResponse
}
// NewV1beta1IssuesResponse instantiates a new V1beta1IssuesResponse and sets the default values.
func NewV1beta1IssuesResponse()(*V1beta1IssuesResponse) {
    m := &V1beta1IssuesResponse{
        V1beta1IssuesGetResponse: *NewV1beta1IssuesGetResponse(),
    }
    return m
}
// CreateV1beta1IssuesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1IssuesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1IssuesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1IssuesGetResponseable instead.
type V1beta1IssuesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1IssuesGetResponseable
}
