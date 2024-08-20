package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubscriptionsGetResponseable instead.
type V1beta1CspAccountsItemCspSubscriptionsResponse struct {
    V1beta1CspAccountsItemCspSubscriptionsGetResponse
}
// NewV1beta1CspAccountsItemCspSubscriptionsResponse instantiates a new V1beta1CspAccountsItemCspSubscriptionsResponse and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsResponse()(*V1beta1CspAccountsItemCspSubscriptionsResponse) {
    m := &V1beta1CspAccountsItemCspSubscriptionsResponse{
        V1beta1CspAccountsItemCspSubscriptionsGetResponse: *NewV1beta1CspAccountsItemCspSubscriptionsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspSubscriptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubscriptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubscriptionsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspSubscriptionsGetResponseable instead.
type V1beta1CspAccountsItemCspSubscriptionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspSubscriptionsGetResponseable
}
