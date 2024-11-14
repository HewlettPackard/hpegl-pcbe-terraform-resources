package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemRemoveHypervisorClustersPostResponseable instead.
type V1beta1SystemsItemRemoveHypervisorClustersResponse struct {
    V1beta1SystemsItemRemoveHypervisorClustersPostResponse
}
// NewV1beta1SystemsItemRemoveHypervisorClustersResponse instantiates a new V1beta1SystemsItemRemoveHypervisorClustersResponse and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorClustersResponse()(*V1beta1SystemsItemRemoveHypervisorClustersResponse) {
    m := &V1beta1SystemsItemRemoveHypervisorClustersResponse{
        V1beta1SystemsItemRemoveHypervisorClustersPostResponse: *NewV1beta1SystemsItemRemoveHypervisorClustersPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemRemoveHypervisorClustersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemRemoveHypervisorClustersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemRemoveHypervisorClustersResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemRemoveHypervisorClustersPostResponseable instead.
type V1beta1SystemsItemRemoveHypervisorClustersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemRemoveHypervisorClustersPostResponseable
}
