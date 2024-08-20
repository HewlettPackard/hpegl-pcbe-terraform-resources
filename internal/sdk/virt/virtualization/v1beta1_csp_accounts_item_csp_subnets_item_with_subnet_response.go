package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse a subnet whose values are defined by, and synchronized with, a cloud service provider.
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable instead.
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse struct {
    V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse
}
// NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse instantiates a new V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse()(*V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse) {
    m := &V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse{
        V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse: *NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable instead.
type V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable
}
