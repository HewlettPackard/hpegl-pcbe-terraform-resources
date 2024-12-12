package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1GroupsItemGroupsGetResponseable instead.
type V1beta1GroupsItemGroupsResponse struct {
    V1beta1GroupsItemGroupsGetResponse
}
// NewV1beta1GroupsItemGroupsResponse instantiates a new V1beta1GroupsItemGroupsResponse and sets the default values.
func NewV1beta1GroupsItemGroupsResponse()(*V1beta1GroupsItemGroupsResponse) {
    m := &V1beta1GroupsItemGroupsResponse{
        V1beta1GroupsItemGroupsGetResponse: *NewV1beta1GroupsItemGroupsGetResponse(),
    }
    return m
}
// CreateV1beta1GroupsItemGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsItemGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsItemGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1GroupsItemGroupsGetResponseable instead.
type V1beta1GroupsItemGroupsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1GroupsItemGroupsGetResponseable
}
