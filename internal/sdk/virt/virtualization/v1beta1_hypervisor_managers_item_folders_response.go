package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemFoldersGetResponseable instead.
type V1beta1HypervisorManagersItemFoldersResponse struct {
    V1beta1HypervisorManagersItemFoldersGetResponse
}
// NewV1beta1HypervisorManagersItemFoldersResponse instantiates a new V1beta1HypervisorManagersItemFoldersResponse and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersResponse()(*V1beta1HypervisorManagersItemFoldersResponse) {
    m := &V1beta1HypervisorManagersItemFoldersResponse{
        V1beta1HypervisorManagersItemFoldersGetResponse: *NewV1beta1HypervisorManagersItemFoldersGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemFoldersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemFoldersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemFoldersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemFoldersGetResponseable instead.
type V1beta1HypervisorManagersItemFoldersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemFoldersGetResponseable
}
