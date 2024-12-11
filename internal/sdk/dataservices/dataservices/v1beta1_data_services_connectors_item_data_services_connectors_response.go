package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsItemDataServicesConnectorsResponse detailed information of Data Services Collector
// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable instead.
type V1beta1DataServicesConnectorsItemDataServicesConnectorsResponse struct {
    V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsResponse instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsResponse and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsResponse()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsResponse) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsResponse{
        V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse: *NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse(),
    }
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable instead.
type V1beta1DataServicesConnectorsItemDataServicesConnectorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable
}