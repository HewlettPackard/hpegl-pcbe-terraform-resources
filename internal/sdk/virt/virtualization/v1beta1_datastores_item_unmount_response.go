package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DatastoresItemUnmountPostResponseable instead.
type V1beta1DatastoresItemUnmountResponse struct {
    V1beta1DatastoresItemUnmountPostResponse
}
// NewV1beta1DatastoresItemUnmountResponse instantiates a new V1beta1DatastoresItemUnmountResponse and sets the default values.
func NewV1beta1DatastoresItemUnmountResponse()(*V1beta1DatastoresItemUnmountResponse) {
    m := &V1beta1DatastoresItemUnmountResponse{
        V1beta1DatastoresItemUnmountPostResponse: *NewV1beta1DatastoresItemUnmountPostResponse(),
    }
    return m
}
// CreateV1beta1DatastoresItemUnmountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemUnmountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemUnmountResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DatastoresItemUnmountPostResponseable instead.
type V1beta1DatastoresItemUnmountResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DatastoresItemUnmountPostResponseable
}
