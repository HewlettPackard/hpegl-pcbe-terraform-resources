package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SystemsItemAssignSecretPostResponseable instead.
type V1beta1SystemsItemAssignSecretResponse struct {
    V1beta1SystemsItemAssignSecretPostResponse
}
// NewV1beta1SystemsItemAssignSecretResponse instantiates a new V1beta1SystemsItemAssignSecretResponse and sets the default values.
func NewV1beta1SystemsItemAssignSecretResponse()(*V1beta1SystemsItemAssignSecretResponse) {
    m := &V1beta1SystemsItemAssignSecretResponse{
        V1beta1SystemsItemAssignSecretPostResponse: *NewV1beta1SystemsItemAssignSecretPostResponse(),
    }
    return m
}
// CreateV1beta1SystemsItemAssignSecretResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAssignSecretResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAssignSecretResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SystemsItemAssignSecretPostResponseable instead.
type V1beta1SystemsItemAssignSecretResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SystemsItemAssignSecretPostResponseable
}
