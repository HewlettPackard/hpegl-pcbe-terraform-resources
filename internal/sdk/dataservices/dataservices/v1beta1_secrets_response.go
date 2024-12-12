package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsResponse report Secrets Response
// Deprecated: This class is obsolete. Use V1beta1SecretsGetResponseable instead.
type V1beta1SecretsResponse struct {
    V1beta1SecretsGetResponse
}
// NewV1beta1SecretsResponse instantiates a new V1beta1SecretsResponse and sets the default values.
func NewV1beta1SecretsResponse()(*V1beta1SecretsResponse) {
    m := &V1beta1SecretsResponse{
        V1beta1SecretsGetResponse: *NewV1beta1SecretsGetResponse(),
    }
    return m
}
// CreateV1beta1SecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SecretsGetResponseable instead.
type V1beta1SecretsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SecretsGetResponseable
}
