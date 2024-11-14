package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs software Catalog with detailed versions of the HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Firmware
type V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reasons why this catalog is unavailable for update
    reasons []string
}
// NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs instantiates a new V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs and sets the default values.
func NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs()(*V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) {
    m := &V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetReasons(res)
        }
        return nil
    }
    return res
}
// GetReasons gets the reasons property value. Reasons why this catalog is unavailable for update
// returns a []string when successful
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) GetReasons()([]string) {
    return m.reasons
}
// Serialize serializes information the current object
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReasons() != nil {
        err := writer.WriteCollectionOfStringValues("reasons", m.GetReasons())
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
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReasons sets the reasons property value. Reasons why this catalog is unavailable for update
func (m *V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogs) SetReasons(value []string)() {
    m.reasons = value
}
type V1beta1SystemsGetResponse_items_softwareInfo_pcForAiSoftwareInfo_unavailableSoftwareCatalogsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReasons()([]string)
    SetReasons(value []string)()
}
