package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector details of the Data Services Connector software.
type V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewV1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector instantiates a new V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector and sets the default values.
func NewV1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector()(*V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector) {
    m := &V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnector) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type V1beta1SystemsItemGetResponse_softwareInfo_pcForAiSoftwareInfo_currentSoftwareCatalog_dataServicesConnectorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
