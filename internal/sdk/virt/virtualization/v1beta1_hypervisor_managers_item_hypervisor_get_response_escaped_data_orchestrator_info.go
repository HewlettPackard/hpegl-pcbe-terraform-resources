package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo data Orchestrator specific information.
type V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique string identifying the data orchestrator.
    id *string
    // The URI reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo instantiates a new V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo()(*V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) {
    m := &V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique string identifying the data orchestrator.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) GetId()(*string) {
    return m.id
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
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
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Unique string identifying the data orchestrator.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) SetId(value *string)() {
    m.id = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1HypervisorManagersItemHypervisorGetResponse_dataOrchestratorInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
