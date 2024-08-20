package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_health server Health information
type V1beta1SystemsItemServersGetResponse_items_health struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The agentlessManagementService property
    agentlessManagementService *string
    // The hbLastUpdateTimestamp property
    hbLastUpdateTimestamp *string
    // The powerState property
    powerState *string
}
// NewV1beta1SystemsItemServersGetResponse_items_health instantiates a new V1beta1SystemsItemServersGetResponse_items_health and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_health()(*V1beta1SystemsItemServersGetResponse_items_health) {
    m := &V1beta1SystemsItemServersGetResponse_items_health{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_healthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_healthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_health(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_health) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAgentlessManagementService gets the agentlessManagementService property value. The agentlessManagementService property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_health) GetAgentlessManagementService()(*string) {
    return m.agentlessManagementService
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_health) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["agentlessManagementService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgentlessManagementService(val)
        }
        return nil
    }
    res["hbLastUpdateTimestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHbLastUpdateTimestamp(val)
        }
        return nil
    }
    res["powerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerState(val)
        }
        return nil
    }
    return res
}
// GetHbLastUpdateTimestamp gets the hbLastUpdateTimestamp property value. The hbLastUpdateTimestamp property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_health) GetHbLastUpdateTimestamp()(*string) {
    return m.hbLastUpdateTimestamp
}
// GetPowerState gets the powerState property value. The powerState property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_health) GetPowerState()(*string) {
    return m.powerState
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_health) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("agentlessManagementService", m.GetAgentlessManagementService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hbLastUpdateTimestamp", m.GetHbLastUpdateTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("powerState", m.GetPowerState())
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
func (m *V1beta1SystemsItemServersGetResponse_items_health) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAgentlessManagementService sets the agentlessManagementService property value. The agentlessManagementService property
func (m *V1beta1SystemsItemServersGetResponse_items_health) SetAgentlessManagementService(value *string)() {
    m.agentlessManagementService = value
}
// SetHbLastUpdateTimestamp sets the hbLastUpdateTimestamp property value. The hbLastUpdateTimestamp property
func (m *V1beta1SystemsItemServersGetResponse_items_health) SetHbLastUpdateTimestamp(value *string)() {
    m.hbLastUpdateTimestamp = value
}
// SetPowerState sets the powerState property value. The powerState property
func (m *V1beta1SystemsItemServersGetResponse_items_health) SetPowerState(value *string)() {
    m.powerState = value
}
type V1beta1SystemsItemServersGetResponse_items_healthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentlessManagementService()(*string)
    GetHbLastUpdateTimestamp()(*string)
    GetPowerState()(*string)
    SetAgentlessManagementService(value *string)()
    SetHbLastUpdateTimestamp(value *string)()
    SetPowerState(value *string)()
}
