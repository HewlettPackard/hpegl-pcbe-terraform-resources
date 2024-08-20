package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1DatastoresItemMountPostResponseable instead.
type V1beta1DatastoresItemMountResponse struct {
    V1beta1DatastoresItemMountPostResponse
}
// NewV1beta1DatastoresItemMountResponse instantiates a new V1beta1DatastoresItemMountResponse and sets the default values.
func NewV1beta1DatastoresItemMountResponse()(*V1beta1DatastoresItemMountResponse) {
    m := &V1beta1DatastoresItemMountResponse{
        V1beta1DatastoresItemMountPostResponse: *NewV1beta1DatastoresItemMountPostResponse(),
    }
    return m
}
// CreateV1beta1DatastoresItemMountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemMountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemMountResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1DatastoresItemMountPostResponseable instead.
type V1beta1DatastoresItemMountResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1DatastoresItemMountPostResponseable
}
