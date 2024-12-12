package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsGetResponseable instead.
type V1beta1DataServicesConnectorsResponse struct {
    V1beta1DataServicesConnectorsGetResponse
}
// NewV1beta1DataServicesConnectorsResponse instantiates a new V1beta1DataServicesConnectorsResponse and sets the default values.
func NewV1beta1DataServicesConnectorsResponse()(*V1beta1DataServicesConnectorsResponse) {
    m := &V1beta1DataServicesConnectorsResponse{
        V1beta1DataServicesConnectorsGetResponse: *NewV1beta1DataServicesConnectorsGetResponse(),
    }
    return m
}
// CreateV1beta1DataServicesConnectorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsGetResponseable instead.
type V1beta1DataServicesConnectorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DataServicesConnectorsGetResponseable
}
