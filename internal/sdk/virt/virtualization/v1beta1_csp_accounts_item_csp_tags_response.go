package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspTagsGetResponseable instead.
type V1beta1CspAccountsItemCspTagsResponse struct {
    V1beta1CspAccountsItemCspTagsGetResponse
}
// NewV1beta1CspAccountsItemCspTagsResponse instantiates a new V1beta1CspAccountsItemCspTagsResponse and sets the default values.
func NewV1beta1CspAccountsItemCspTagsResponse()(*V1beta1CspAccountsItemCspTagsResponse) {
    m := &V1beta1CspAccountsItemCspTagsResponse{
        V1beta1CspAccountsItemCspTagsGetResponse: *NewV1beta1CspAccountsItemCspTagsGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspTagsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspTagsGetResponseable instead.
type V1beta1CspAccountsItemCspTagsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspTagsGetResponseable
}
