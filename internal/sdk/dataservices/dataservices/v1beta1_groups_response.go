package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1GroupsGetResponseable instead.
type V1beta1GroupsResponse struct {
    V1beta1GroupsGetResponse
}
// NewV1beta1GroupsResponse instantiates a new V1beta1GroupsResponse and sets the default values.
func NewV1beta1GroupsResponse()(*V1beta1GroupsResponse) {
    m := &V1beta1GroupsResponse{
        V1beta1GroupsGetResponse: *NewV1beta1GroupsGetResponse(),
    }
    return m
}
// CreateV1beta1GroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1GroupsGetResponseable instead.
type V1beta1GroupsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1GroupsGetResponseable
}
