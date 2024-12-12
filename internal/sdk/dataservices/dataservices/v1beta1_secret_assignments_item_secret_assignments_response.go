package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsItemSecretAssignmentsResponse assignment Resource Definition
// Deprecated: This class is obsolete. Use V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable instead.
type V1beta1SecretAssignmentsItemSecretAssignmentsResponse struct {
    V1beta1SecretAssignmentsItemSecretAssignmentsGetResponse
}
// NewV1beta1SecretAssignmentsItemSecretAssignmentsResponse instantiates a new V1beta1SecretAssignmentsItemSecretAssignmentsResponse and sets the default values.
func NewV1beta1SecretAssignmentsItemSecretAssignmentsResponse()(*V1beta1SecretAssignmentsItemSecretAssignmentsResponse) {
    m := &V1beta1SecretAssignmentsItemSecretAssignmentsResponse{
        V1beta1SecretAssignmentsItemSecretAssignmentsGetResponse: *NewV1beta1SecretAssignmentsItemSecretAssignmentsGetResponse(),
    }
    return m
}
// CreateV1beta1SecretAssignmentsItemSecretAssignmentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsItemSecretAssignmentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsItemSecretAssignmentsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable instead.
type V1beta1SecretAssignmentsItemSecretAssignmentsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable
}
