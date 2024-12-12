package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1StorageLocationsGetResponseable instead.
type V1beta1StorageLocationsResponse struct {
    V1beta1StorageLocationsGetResponse
}
// NewV1beta1StorageLocationsResponse instantiates a new V1beta1StorageLocationsResponse and sets the default values.
func NewV1beta1StorageLocationsResponse()(*V1beta1StorageLocationsResponse) {
    m := &V1beta1StorageLocationsResponse{
        V1beta1StorageLocationsGetResponse: *NewV1beta1StorageLocationsGetResponse(),
    }
    return m
}
// CreateV1beta1StorageLocationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1StorageLocationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1StorageLocationsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1StorageLocationsGetResponseable instead.
type V1beta1StorageLocationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1StorageLocationsGetResponseable
}
