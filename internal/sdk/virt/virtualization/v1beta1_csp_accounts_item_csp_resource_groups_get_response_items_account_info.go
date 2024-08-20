package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo information about the cloud provider account where the resource is located.
type V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo instantiates a new V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo()(*V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo) {
    m := &V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_accountInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
