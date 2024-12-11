package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareUpgradesResponse a list of Software Upgrades.
// Deprecated: This class is obsolete. Use V1beta1SoftwareUpgradesGetResponseable instead.
type V1beta1SoftwareUpgradesResponse struct {
    V1beta1SoftwareUpgradesGetResponse
}
// NewV1beta1SoftwareUpgradesResponse instantiates a new V1beta1SoftwareUpgradesResponse and sets the default values.
func NewV1beta1SoftwareUpgradesResponse()(*V1beta1SoftwareUpgradesResponse) {
    m := &V1beta1SoftwareUpgradesResponse{
        V1beta1SoftwareUpgradesGetResponse: *NewV1beta1SoftwareUpgradesGetResponse(),
    }
    return m
}
// CreateV1beta1SoftwareUpgradesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareUpgradesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareUpgradesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SoftwareUpgradesGetResponseable instead.
type V1beta1SoftwareUpgradesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SoftwareUpgradesGetResponseable
}