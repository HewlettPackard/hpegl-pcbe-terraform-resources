package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesItemCspMachineImagesResponse details of a CSP machine image
// Deprecated: This class is obsolete. Use V1beta1CspMachineImagesItemCspMachineImagesGetResponseable instead.
type V1beta1CspMachineImagesItemCspMachineImagesResponse struct {
    V1beta1CspMachineImagesItemCspMachineImagesGetResponse
}
// NewV1beta1CspMachineImagesItemCspMachineImagesResponse instantiates a new V1beta1CspMachineImagesItemCspMachineImagesResponse and sets the default values.
func NewV1beta1CspMachineImagesItemCspMachineImagesResponse()(*V1beta1CspMachineImagesItemCspMachineImagesResponse) {
    m := &V1beta1CspMachineImagesItemCspMachineImagesResponse{
        V1beta1CspMachineImagesItemCspMachineImagesGetResponse: *NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse(),
    }
    return m
}
// CreateV1beta1CspMachineImagesItemCspMachineImagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesItemCspMachineImagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesItemCspMachineImagesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspMachineImagesItemCspMachineImagesGetResponseable instead.
type V1beta1CspMachineImagesItemCspMachineImagesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspMachineImagesItemCspMachineImagesGetResponseable
}
