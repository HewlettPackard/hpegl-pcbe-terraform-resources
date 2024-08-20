package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AsyncOperationsItemAsyncOperationsResponse the operation resource provides information about progress of asynchronousrequest processing and is where associated resources can be found.
// Deprecated: This class is obsolete. Use V1beta1AsyncOperationsItemAsyncOperationsGetResponseable instead.
type V1beta1AsyncOperationsItemAsyncOperationsResponse struct {
    V1beta1AsyncOperationsItemAsyncOperationsGetResponse
}
// NewV1beta1AsyncOperationsItemAsyncOperationsResponse instantiates a new V1beta1AsyncOperationsItemAsyncOperationsResponse and sets the default values.
func NewV1beta1AsyncOperationsItemAsyncOperationsResponse()(*V1beta1AsyncOperationsItemAsyncOperationsResponse) {
    m := &V1beta1AsyncOperationsItemAsyncOperationsResponse{
        V1beta1AsyncOperationsItemAsyncOperationsGetResponse: *NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse(),
    }
    return m
}
// CreateV1beta1AsyncOperationsItemAsyncOperationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsItemAsyncOperationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsItemAsyncOperationsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1AsyncOperationsItemAsyncOperationsGetResponseable instead.
type V1beta1AsyncOperationsItemAsyncOperationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1AsyncOperationsItemAsyncOperationsGetResponseable
}
