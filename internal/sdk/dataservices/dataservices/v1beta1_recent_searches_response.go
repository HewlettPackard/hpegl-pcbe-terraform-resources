package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1RecentSearchesResponse contains a paginated list of recent search items
// Deprecated: This class is obsolete. Use V1beta1RecentSearchesGetResponseable instead.
type V1beta1RecentSearchesResponse struct {
    V1beta1RecentSearchesGetResponse
}
// NewV1beta1RecentSearchesResponse instantiates a new V1beta1RecentSearchesResponse and sets the default values.
func NewV1beta1RecentSearchesResponse()(*V1beta1RecentSearchesResponse) {
    m := &V1beta1RecentSearchesResponse{
        V1beta1RecentSearchesGetResponse: *NewV1beta1RecentSearchesGetResponse(),
    }
    return m
}
// CreateV1beta1RecentSearchesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1RecentSearchesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1RecentSearchesResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1RecentSearchesGetResponseable instead.
type V1beta1RecentSearchesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1RecentSearchesGetResponseable
}
