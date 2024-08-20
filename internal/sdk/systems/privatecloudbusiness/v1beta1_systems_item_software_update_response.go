package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareUpdatePostResponseable instead.
type V1beta1SystemsItemSoftwareUpdateResponse struct {
    V1beta1SystemsItemSoftwareUpdatePostResponse
}
// NewV1beta1SystemsItemSoftwareUpdateResponse instantiates a new V1beta1SystemsItemSoftwareUpdateResponse and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateResponse()(*V1beta1SystemsItemSoftwareUpdateResponse) {
    m := &V1beta1SystemsItemSoftwareUpdateResponse{
        V1beta1SystemsItemSoftwareUpdatePostResponse: *NewV1beta1SystemsItemSoftwareUpdatePostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSoftwareUpdateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareUpdateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareUpdateResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareUpdatePostResponseable instead.
type V1beta1SystemsItemSoftwareUpdateResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSoftwareUpdatePostResponseable
}
