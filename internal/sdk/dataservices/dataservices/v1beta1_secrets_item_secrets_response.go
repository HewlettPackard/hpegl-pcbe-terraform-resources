package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsItemSecretsResponse secret Resource Definition
// Deprecated: This class is obsolete. Use V1beta1SecretsItemSecretsGetResponseable instead.
type V1beta1SecretsItemSecretsResponse struct {
    V1beta1SecretsItemSecretsGetResponse
}
// NewV1beta1SecretsItemSecretsResponse instantiates a new V1beta1SecretsItemSecretsResponse and sets the default values.
func NewV1beta1SecretsItemSecretsResponse()(*V1beta1SecretsItemSecretsResponse) {
    m := &V1beta1SecretsItemSecretsResponse{
        V1beta1SecretsItemSecretsGetResponse: *NewV1beta1SecretsItemSecretsGetResponse(),
    }
    return m
}
// CreateV1beta1SecretsItemSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecretsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SecretsItemSecretsGetResponseable instead.
type V1beta1SecretsItemSecretsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SecretsItemSecretsGetResponseable
}
