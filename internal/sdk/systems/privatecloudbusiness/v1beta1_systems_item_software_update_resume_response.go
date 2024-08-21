package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareUpdateResumePostResponseable instead.
type V1beta1SystemsItemSoftwareUpdateResumeResponse struct {
    V1beta1SystemsItemSoftwareUpdateResumePostResponse
}
// NewV1beta1SystemsItemSoftwareUpdateResumeResponse instantiates a new V1beta1SystemsItemSoftwareUpdateResumeResponse and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateResumeResponse()(*V1beta1SystemsItemSoftwareUpdateResumeResponse) {
    m := &V1beta1SystemsItemSoftwareUpdateResumeResponse{
        V1beta1SystemsItemSoftwareUpdateResumePostResponse: *NewV1beta1SystemsItemSoftwareUpdateResumePostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSoftwareUpdateResumeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareUpdateResumeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareUpdateResumeResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareUpdateResumePostResponseable instead.
type V1beta1SystemsItemSoftwareUpdateResumeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSoftwareUpdateResumePostResponseable
}