package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspVpcsGetResponseable instead.
type V1beta1CspAccountsItemCspVpcsResponse struct {
    V1beta1CspAccountsItemCspVpcsGetResponse
}
// NewV1beta1CspAccountsItemCspVpcsResponse instantiates a new V1beta1CspAccountsItemCspVpcsResponse and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsResponse()(*V1beta1CspAccountsItemCspVpcsResponse) {
    m := &V1beta1CspAccountsItemCspVpcsResponse{
        V1beta1CspAccountsItemCspVpcsGetResponse: *NewV1beta1CspAccountsItemCspVpcsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspVpcsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspVpcsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspVpcsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspVpcsGetResponseable instead.
type V1beta1CspAccountsItemCspVpcsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspVpcsGetResponseable
}
