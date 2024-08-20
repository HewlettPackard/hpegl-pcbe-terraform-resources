package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemValidatePostResponseable instead.
type V1beta1CspAccountsItemValidateResponse struct {
    V1beta1CspAccountsItemValidatePostResponse
}
// NewV1beta1CspAccountsItemValidateResponse instantiates a new V1beta1CspAccountsItemValidateResponse and sets the default values.
func NewV1beta1CspAccountsItemValidateResponse()(*V1beta1CspAccountsItemValidateResponse) {
    m := &V1beta1CspAccountsItemValidateResponse{
        V1beta1CspAccountsItemValidatePostResponse: *NewV1beta1CspAccountsItemValidatePostResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemValidateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemValidateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemValidateResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemValidatePostResponseable instead.
type V1beta1CspAccountsItemValidateResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemValidatePostResponseable
}
