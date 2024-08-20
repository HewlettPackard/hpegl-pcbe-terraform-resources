package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubnetsGetResponseable instead.
type V1beta1CspAccountsItemCspSubnetsResponse struct {
    V1beta1CspAccountsItemCspSubnetsGetResponse
}
// NewV1beta1CspAccountsItemCspSubnetsResponse instantiates a new V1beta1CspAccountsItemCspSubnetsResponse and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsResponse()(*V1beta1CspAccountsItemCspSubnetsResponse) {
    m := &V1beta1CspAccountsItemCspSubnetsResponse{
        V1beta1CspAccountsItemCspSubnetsGetResponse: *NewV1beta1CspAccountsItemCspSubnetsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspSubnetsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubnetsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubnetsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubnetsGetResponseable instead.
type V1beta1CspAccountsItemCspSubnetsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspSubnetsGetResponseable
}
