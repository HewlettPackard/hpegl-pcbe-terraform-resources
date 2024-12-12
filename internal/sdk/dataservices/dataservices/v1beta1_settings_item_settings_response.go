package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1beta1SettingsItemSettingsGetResponseable instead.
type V1beta1SettingsItemSettingsResponse struct {
    V1beta1SettingsItemSettingsGetResponse
}
// NewV1beta1SettingsItemSettingsResponse instantiates a new V1beta1SettingsItemSettingsResponse and sets the default values.
func NewV1beta1SettingsItemSettingsResponse()(*V1beta1SettingsItemSettingsResponse) {
    m := &V1beta1SettingsItemSettingsResponse{
        V1beta1SettingsItemSettingsGetResponse: *NewV1beta1SettingsItemSettingsGetResponse(),
    }
    return m
}
// CreateV1beta1SettingsItemSettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SettingsItemSettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SettingsItemSettingsResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1SettingsItemSettingsGetResponseable instead.
type V1beta1SettingsItemSettingsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1SettingsItemSettingsGetResponseable
}
