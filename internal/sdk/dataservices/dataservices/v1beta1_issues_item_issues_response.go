package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1IssuesItemIssuesGetResponseable instead.
type V1beta1IssuesItemIssuesResponse struct {
    V1beta1IssuesItemIssuesGetResponse
}
// NewV1beta1IssuesItemIssuesResponse instantiates a new V1beta1IssuesItemIssuesResponse and sets the default values.
func NewV1beta1IssuesItemIssuesResponse()(*V1beta1IssuesItemIssuesResponse) {
    m := &V1beta1IssuesItemIssuesResponse{
        V1beta1IssuesItemIssuesGetResponse: *NewV1beta1IssuesItemIssuesGetResponse(),
    }
    return m
}
// CreateV1beta1IssuesItemIssuesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1IssuesItemIssuesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1IssuesItemIssuesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1IssuesItemIssuesGetResponseable instead.
type V1beta1IssuesItemIssuesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1IssuesItemIssuesGetResponseable
}
