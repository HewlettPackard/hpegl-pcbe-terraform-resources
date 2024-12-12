package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsResponse report Assignments Response
// Deprecated: This class is obsolete. Use V1beta1SecretAssignmentsGetResponseable instead.
type V1beta1SecretAssignmentsResponse struct {
    V1beta1SecretAssignmentsGetResponse
}
// NewV1beta1SecretAssignmentsResponse instantiates a new V1beta1SecretAssignmentsResponse and sets the default values.
func NewV1beta1SecretAssignmentsResponse()(*V1beta1SecretAssignmentsResponse) {
    m := &V1beta1SecretAssignmentsResponse{
        V1beta1SecretAssignmentsGetResponse: *NewV1beta1SecretAssignmentsGetResponse(),
    }
    return m
}
// CreateV1beta1SecretAssignmentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SecretAssignmentsGetResponseable instead.
type V1beta1SecretAssignmentsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SecretAssignmentsGetResponseable
}
