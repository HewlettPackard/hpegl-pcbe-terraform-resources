package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspVpcsItemWithVpcResponse a virtual private cloud (VPC) whose values are defined by, and synchronized with, a cloud service provider.
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable instead.
type V1beta1CspAccountsItemCspVpcsItemWithVpcResponse struct {
    V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponse
}
// NewV1beta1CspAccountsItemCspVpcsItemWithVpcResponse instantiates a new V1beta1CspAccountsItemCspVpcsItemWithVpcResponse and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsItemWithVpcResponse()(*V1beta1CspAccountsItemCspVpcsItemWithVpcResponse) {
    m := &V1beta1CspAccountsItemCspVpcsItemWithVpcResponse{
        V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponse: *NewV1beta1CspAccountsItemCspVpcsItemWithVpcGetResponse(),
    }
    return m
}
// CreateV1beta1CspAccountsItemCspVpcsItemWithVpcResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspVpcsItemWithVpcResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspVpcsItemWithVpcResponse(), nil
}
// Deprecated: This class is obsolete. Use V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable instead.
type V1beta1CspAccountsItemCspVpcsItemWithVpcResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable
}
