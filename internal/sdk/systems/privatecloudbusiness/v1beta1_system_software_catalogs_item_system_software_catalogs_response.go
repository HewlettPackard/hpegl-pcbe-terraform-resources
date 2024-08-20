package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse system Software Catalog with all the detailed information of software components, end user license agreement and a list of systems that have an update path to the catalog.
// Deprecated: This class is obsolete. Use V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable instead.
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse struct {
    V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse
}
// NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse instantiates a new V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse and sets the default values.
func NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse()(*V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse) {
    m := &V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse{
        V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse: *NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponse(),
    }
    return m
}
// CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable instead.
type V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemSoftwareCatalogsItemSystemSoftwareCatalogsGetResponseable
}
