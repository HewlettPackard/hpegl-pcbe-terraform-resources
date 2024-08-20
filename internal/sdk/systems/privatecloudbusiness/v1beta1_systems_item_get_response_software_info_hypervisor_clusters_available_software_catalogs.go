package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs software Catalog with detailed versions of the HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Firmware
type V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Boolean flag indicating whether this is the latest version available for software update.
    isLatest *bool
}
// NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs instantiates a new V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs and sets the default values.
func NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs()(*V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) {
    m := &V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isLatest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLatest(val)
        }
        return nil
    }
    return res
}
// GetIsLatest gets the isLatest property value. Boolean flag indicating whether this is the latest version available for software update.
// returns a *bool when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) GetIsLatest()(*bool) {
    return m.isLatest
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isLatest", m.GetIsLatest())
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
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsLatest sets the isLatest property value. Boolean flag indicating whether this is the latest version available for software update.
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogs) SetIsLatest(value *bool)() {
    m.isLatest = value
}
type V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_availableSoftwareCatalogsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsLatest()(*bool)
    SetIsLatest(value *bool)()
}
