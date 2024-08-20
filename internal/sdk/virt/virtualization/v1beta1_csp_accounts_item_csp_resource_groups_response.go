package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspResourceGroupsGetResponseable instead.
type V1beta1CspAccountsItemCspResourceGroupsResponse struct {
    V1beta1CspAccountsItemCspResourceGroupsGetResponse
}
// NewV1beta1CspAccountsItemCspResourceGroupsResponse instantiates a new V1beta1CspAccountsItemCspResourceGroupsResponse and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsResponse()(*V1beta1CspAccountsItemCspResourceGroupsResponse) {
    m := &V1beta1CspAccountsItemCspResourceGroupsResponse{
        V1beta1CspAccountsItemCspResourceGroupsGetResponse: *NewV1beta1CspAccountsItemCspResourceGroupsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspResourceGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspResourceGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspResourceGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspResourceGroupsGetResponseable instead.
type V1beta1CspAccountsItemCspResourceGroupsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspResourceGroupsGetResponseable
}
