package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1GroupsItemAssociatedResourcesGetResponseable instead.
type V1beta1GroupsItemAssociatedResourcesResponse struct {
    V1beta1GroupsItemAssociatedResourcesGetResponse
}
// NewV1beta1GroupsItemAssociatedResourcesResponse instantiates a new V1beta1GroupsItemAssociatedResourcesResponse and sets the default values.
func NewV1beta1GroupsItemAssociatedResourcesResponse()(*V1beta1GroupsItemAssociatedResourcesResponse) {
    m := &V1beta1GroupsItemAssociatedResourcesResponse{
        V1beta1GroupsItemAssociatedResourcesGetResponse: *NewV1beta1GroupsItemAssociatedResourcesGetResponse(),
    }
    return m
}
// CreateV1beta1GroupsItemAssociatedResourcesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsItemAssociatedResourcesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsItemAssociatedResourcesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1GroupsItemAssociatedResourcesGetResponseable instead.
type V1beta1GroupsItemAssociatedResourcesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1GroupsItemAssociatedResourcesGetResponseable
}