package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1 subnet properties whose values are defined by the AWS cloud service provider.
type V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud provider availability zone
    availabilityZone *string
    // Cloud provider region
    cspRegion *string
}
// NewV1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1 instantiates a new V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1 and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1()(*V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) {
    m := &V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityZone gets the availabilityZone property value. Cloud provider availability zone
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) GetAvailabilityZone()(*string) {
    return m.availabilityZone
}
// GetCspRegion gets the cspRegion property value. Cloud provider region
// returns a *string when successful
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) GetCspRegion()(*string) {
    return m.cspRegion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityZone(val)
        }
        return nil
    }
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
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityZone sets the availabilityZone property value. Cloud provider availability zone
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) SetAvailabilityZone(value *string)() {
    m.availabilityZone = value
}
// SetCspRegion sets the cspRegion property value. Cloud provider region
func (m *V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1) SetCspRegion(value *string)() {
    m.cspRegion = value
}
type V1beta1CspAccountsItemCspSubnetsGetResponse_items_cspInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityZone()(*string)
    GetCspRegion()(*string)
    SetAvailabilityZone(value *string)()
    SetCspRegion(value *string)()
}
