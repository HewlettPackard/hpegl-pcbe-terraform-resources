package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog software Catalog with detailed versions of various software components like HPE Storage Software, hypervisor software, HPE Storage Connection Manager and Server firmware. If the version of this catalog is set to 'Unavailable', it means one or more software component versions are not available. If the version of this catalog is set to 'Non-Compliant', it means the current set of software component versions does not match any of the supported software catalog versions.
type V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reasons if any as to why the current software catalog version cannot be determined
    reasons []string
}
// NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog instantiates a new V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog and sets the default values.
func NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog()(*V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) {
    m := &V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetReasons gets the reasons property value. Reasons if any as to why the current software catalog version cannot be determined
// returns a []string when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) GetReasons()([]string) {
    return m.reasons
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReasons sets the reasons property value. Reasons if any as to why the current software catalog version cannot be determined
func (m *V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalog) SetReasons(value []string)() {
    m.reasons = value
}
type V1beta1SystemsItemGetResponse_softwareInfo_hypervisorClusters_currentSoftwareCatalogable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReasons()([]string)
    SetReasons(value []string)()
}
