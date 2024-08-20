package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresItemDatastoresResponse represents a single instance of a Datastore
// Deprecated: This class is obsolete. Use V1beta1DatastoresItemDatastoresGetResponseable instead.
type V1beta1DatastoresItemDatastoresResponse struct {
    V1beta1DatastoresItemDatastoresGetResponse
}
// NewV1beta1DatastoresItemDatastoresResponse instantiates a new V1beta1DatastoresItemDatastoresResponse and sets the default values.
func NewV1beta1DatastoresItemDatastoresResponse()(*V1beta1DatastoresItemDatastoresResponse) {
    m := &V1beta1DatastoresItemDatastoresResponse{
        V1beta1DatastoresItemDatastoresGetResponse: *NewV1beta1DatastoresItemDatastoresGetResponse(),
    }
    return m
}
// CreateV1beta1DatastoresItemDatastoresResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemDatastoresResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemDatastoresResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DatastoresItemDatastoresGetResponseable instead.
type V1beta1DatastoresItemDatastoresResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DatastoresItemDatastoresGetResponseable
}
