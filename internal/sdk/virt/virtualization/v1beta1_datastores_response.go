package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DatastoresGetResponseable instead.
type V1beta1DatastoresResponse struct {
    V1beta1DatastoresGetResponse
}
// NewV1beta1DatastoresResponse instantiates a new V1beta1DatastoresResponse and sets the default values.
func NewV1beta1DatastoresResponse()(*V1beta1DatastoresResponse) {
    m := &V1beta1DatastoresResponse{
        V1beta1DatastoresGetResponse: *NewV1beta1DatastoresGetResponse(),
    }
    return m
}
// CreateV1beta1DatastoresResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DatastoresGetResponseable instead.
type V1beta1DatastoresResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DatastoresGetResponseable
}
