package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemStorageReplicationPartnersGetResponseable instead.
type V1beta1SystemsItemStorageReplicationPartnersResponse struct {
    V1beta1SystemsItemStorageReplicationPartnersGetResponse
}
// NewV1beta1SystemsItemStorageReplicationPartnersResponse instantiates a new V1beta1SystemsItemStorageReplicationPartnersResponse and sets the default values.
func NewV1beta1SystemsItemStorageReplicationPartnersResponse()(*V1beta1SystemsItemStorageReplicationPartnersResponse) {
    m := &V1beta1SystemsItemStorageReplicationPartnersResponse{
        V1beta1SystemsItemStorageReplicationPartnersGetResponse: *NewV1beta1SystemsItemStorageReplicationPartnersGetResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemStorageReplicationPartnersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageReplicationPartnersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageReplicationPartnersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemStorageReplicationPartnersGetResponseable instead.
type V1beta1SystemsItemStorageReplicationPartnersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemStorageReplicationPartnersGetResponseable
}
