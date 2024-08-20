package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareVersionRefreshPostResponseable instead.
type V1beta1SystemsItemSoftwareVersionRefreshResponse struct {
    V1beta1SystemsItemSoftwareVersionRefreshPostResponse
}
// NewV1beta1SystemsItemSoftwareVersionRefreshResponse instantiates a new V1beta1SystemsItemSoftwareVersionRefreshResponse and sets the default values.
func NewV1beta1SystemsItemSoftwareVersionRefreshResponse()(*V1beta1SystemsItemSoftwareVersionRefreshResponse) {
    m := &V1beta1SystemsItemSoftwareVersionRefreshResponse{
        V1beta1SystemsItemSoftwareVersionRefreshPostResponse: *NewV1beta1SystemsItemSoftwareVersionRefreshPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSoftwareVersionRefreshResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareVersionRefreshResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareVersionRefreshResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareVersionRefreshPostResponseable instead.
type V1beta1SystemsItemSoftwareVersionRefreshResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSoftwareVersionRefreshPostResponseable
}
