package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SettingsGetResponseable instead.
type V1beta1SettingsResponse struct {
    V1beta1SettingsGetResponse
}
// NewV1beta1SettingsResponse instantiates a new V1beta1SettingsResponse and sets the default values.
func NewV1beta1SettingsResponse()(*V1beta1SettingsResponse) {
    m := &V1beta1SettingsResponse{
        V1beta1SettingsGetResponse: *NewV1beta1SettingsGetResponse(),
    }
    return m
}
// CreateV1beta1SettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SettingsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SettingsGetResponseable instead.
type V1beta1SettingsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SettingsGetResponseable
}
