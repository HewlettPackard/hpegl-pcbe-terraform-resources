package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1StorageLocationsItemStorageLocationsResponse required properties of a storage location.
// Deprecated: This class is obsolete. Use V1beta1StorageLocationsItemStorageLocationsGetResponseable instead.
type V1beta1StorageLocationsItemStorageLocationsResponse struct {
    V1beta1StorageLocationsItemStorageLocationsGetResponse
}
// NewV1beta1StorageLocationsItemStorageLocationsResponse instantiates a new V1beta1StorageLocationsItemStorageLocationsResponse and sets the default values.
func NewV1beta1StorageLocationsItemStorageLocationsResponse()(*V1beta1StorageLocationsItemStorageLocationsResponse) {
    m := &V1beta1StorageLocationsItemStorageLocationsResponse{
        V1beta1StorageLocationsItemStorageLocationsGetResponse: *NewV1beta1StorageLocationsItemStorageLocationsGetResponse(),
    }
    return m
}
// CreateV1beta1StorageLocationsItemStorageLocationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1StorageLocationsItemStorageLocationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1StorageLocationsItemStorageLocationsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1StorageLocationsItemStorageLocationsGetResponseable instead.
type V1beta1StorageLocationsItemStorageLocationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1StorageLocationsItemStorageLocationsGetResponseable
}
