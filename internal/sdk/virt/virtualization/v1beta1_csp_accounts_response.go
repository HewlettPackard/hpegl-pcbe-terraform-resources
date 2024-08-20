package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsGetResponseable instead.
type V1beta1CspAccountsResponse struct {
    V1beta1CspAccountsGetResponse
}
// NewV1beta1CspAccountsResponse instantiates a new V1beta1CspAccountsResponse and sets the default values.
func NewV1beta1CspAccountsResponse()(*V1beta1CspAccountsResponse) {
    m := &V1beta1CspAccountsResponse{
        V1beta1CspAccountsGetResponse: *NewV1beta1CspAccountsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsGetResponseable instead.
type V1beta1CspAccountsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsGetResponseable
}
