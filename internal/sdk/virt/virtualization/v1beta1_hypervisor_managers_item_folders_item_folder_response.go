package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemFoldersItemFolderResponse represents a hypervisor folder resource
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable instead.
type V1beta1HypervisorManagersItemFoldersItemFolderResponse struct {
    V1beta1HypervisorManagersItemFoldersItemFolderGetResponse
}
// NewV1beta1HypervisorManagersItemFoldersItemFolderResponse instantiates a new V1beta1HypervisorManagersItemFoldersItemFolderResponse and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersItemFolderResponse()(*V1beta1HypervisorManagersItemFoldersItemFolderResponse) {
    m := &V1beta1HypervisorManagersItemFoldersItemFolderResponse{
        V1beta1HypervisorManagersItemFoldersItemFolderGetResponse: *NewV1beta1HypervisorManagersItemFoldersItemFolderGetResponse(),
    }
    return m
}
// CreateV1beta1HypervisorManagersItemFoldersItemFolderResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemFoldersItemFolderResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemFoldersItemFolderResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable instead.
type V1beta1HypervisorManagersItemFoldersItemFolderResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable
}
