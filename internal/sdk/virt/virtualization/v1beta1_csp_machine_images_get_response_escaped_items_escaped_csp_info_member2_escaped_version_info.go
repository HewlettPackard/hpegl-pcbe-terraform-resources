package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo image version info.
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Merchant's offer.
    offer *string
    // Publisher of the image.
    publisher *string
    // SKU of the image.
    sku *string
    // Version of the image.
    version *string
}
// NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo instantiates a new V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo and sets the default values.
func NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo()(*V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) {
    m := &V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["offer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffer(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["sku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSku(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetOffer gets the offer property value. Merchant's offer.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetOffer()(*string) {
    return m.offer
}
// GetPublisher gets the publisher property value. Publisher of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetPublisher()(*string) {
    return m.publisher
}
// GetSku gets the sku property value. SKU of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetSku()(*string) {
    return m.sku
}
// GetVersion gets the version property value. Version of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("offer", m.GetOffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sku", m.GetSku())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOffer sets the offer property value. Merchant's offer.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) SetOffer(value *string)() {
    m.offer = value
}
// SetPublisher sets the publisher property value. Publisher of the image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) SetPublisher(value *string)() {
    m.publisher = value
}
// SetSku sets the sku property value. SKU of the image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) SetSku(value *string)() {
    m.sku = value
}
// SetVersion sets the version property value. Version of the image.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfo) SetVersion(value *string)() {
    m.version = value
}
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember2_versionInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOffer()(*string)
    GetPublisher()(*string)
    GetSku()(*string)
    GetVersion()(*string)
    SetOffer(value *string)()
    SetPublisher(value *string)()
    SetSku(value *string)()
    SetVersion(value *string)()
}
