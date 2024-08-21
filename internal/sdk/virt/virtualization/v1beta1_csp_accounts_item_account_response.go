package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemAccountDeleteResponseable instead.
type V1beta1CspAccountsItemAccountResponse struct {
    V1beta1CspAccountsItemAccountDeleteResponse
}
// NewV1beta1CspAccountsItemAccountResponse instantiates a new V1beta1CspAccountsItemAccountResponse and sets the default values.
func NewV1beta1CspAccountsItemAccountResponse()(*V1beta1CspAccountsItemAccountResponse) {
    m := &V1beta1CspAccountsItemAccountResponse{
        V1beta1CspAccountsItemAccountDeleteResponse: *NewV1beta1CspAccountsItemAccountDeleteResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemAccountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemAccountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemAccountResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemAccountDeleteResponseable instead.
type V1beta1CspAccountsItemAccountResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemAccountDeleteResponseable
}