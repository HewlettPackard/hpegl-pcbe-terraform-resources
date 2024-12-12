package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo resource group properties whose values are defined by, and synchronized with Azure cloud service provider.
type V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider region
    cspRegion *string
}
// NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo instantiates a new V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo()(*V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) {
    m := &V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cspRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspRegion(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfo) SetCspRegion(value *string)() {
    m.cspRegion = value
}
type V1beta1CspAccountsItemCspResourceGroupsGetResponse_items_cspInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCspRegion()(*string)
    SetCspRegion(value *string)()
}
