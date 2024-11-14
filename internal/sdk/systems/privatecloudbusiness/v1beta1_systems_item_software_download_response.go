package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareDownloadPostResponseable instead.
type V1beta1SystemsItemSoftwareDownloadResponse struct {
    V1beta1SystemsItemSoftwareDownloadPostResponse
}
// NewV1beta1SystemsItemSoftwareDownloadResponse instantiates a new V1beta1SystemsItemSoftwareDownloadResponse and sets the default values.
func NewV1beta1SystemsItemSoftwareDownloadResponse()(*V1beta1SystemsItemSoftwareDownloadResponse) {
    m := &V1beta1SystemsItemSoftwareDownloadResponse{
        V1beta1SystemsItemSoftwareDownloadPostResponse: *NewV1beta1SystemsItemSoftwareDownloadPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemSoftwareDownloadResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareDownloadResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareDownloadResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemSoftwareDownloadPostResponseable instead.
type V1beta1SystemsItemSoftwareDownloadResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemSoftwareDownloadPostResponseable
}
