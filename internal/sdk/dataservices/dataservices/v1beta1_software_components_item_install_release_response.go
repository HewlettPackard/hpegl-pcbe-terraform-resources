package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareComponentsItemInstallReleaseResponse a release of software.
// Deprecated: This class is obsolete. Use V1beta1SoftwareComponentsItemInstallReleaseGetResponseable instead.
type V1beta1SoftwareComponentsItemInstallReleaseResponse struct {
    V1beta1SoftwareComponentsItemInstallReleaseGetResponse
}
// NewV1beta1SoftwareComponentsItemInstallReleaseResponse instantiates a new V1beta1SoftwareComponentsItemInstallReleaseResponse and sets the default values.
func NewV1beta1SoftwareComponentsItemInstallReleaseResponse()(*V1beta1SoftwareComponentsItemInstallReleaseResponse) {
    m := &V1beta1SoftwareComponentsItemInstallReleaseResponse{
        V1beta1SoftwareComponentsItemInstallReleaseGetResponse: *NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse(),
    }
    return m
}
// CreateV1beta1SoftwareComponentsItemInstallReleaseResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareComponentsItemInstallReleaseResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareComponentsItemInstallReleaseResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SoftwareComponentsItemInstallReleaseGetResponseable instead.
type V1beta1SoftwareComponentsItemInstallReleaseResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SoftwareComponentsItemInstallReleaseGetResponseable
}
