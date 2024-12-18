package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable instead.
type V1beta1DataServicesConnectorsItemGenerateTotpResponse struct {
    V1beta1DataServicesConnectorsItemGenerateTotpPostResponse
}
// NewV1beta1DataServicesConnectorsItemGenerateTotpResponse instantiates a new V1beta1DataServicesConnectorsItemGenerateTotpResponse and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateTotpResponse()(*V1beta1DataServicesConnectorsItemGenerateTotpResponse) {
    m := &V1beta1DataServicesConnectorsItemGenerateTotpResponse{
        V1beta1DataServicesConnectorsItemGenerateTotpPostResponse: *NewV1beta1DataServicesConnectorsItemGenerateTotpPostResponse(),
    }
    return m
}
// CreateV1beta1DataServicesConnectorsItemGenerateTotpResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemGenerateTotpResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemGenerateTotpResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable instead.
type V1beta1DataServicesConnectorsItemGenerateTotpResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable
}
