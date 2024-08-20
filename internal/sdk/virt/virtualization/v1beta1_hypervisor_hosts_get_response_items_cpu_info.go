package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsGetResponse_items_cpuInfo cPU information.
type V1beta1HypervisorHostsGetResponse_items_cpuInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicate whether hyperthreading is active or not.
    hyperthreadingActive *bool
    // Number of logical processors.
    logicalProcessors *int32
    // Number of processor cores.
    processorCores *int32
    // Number of processor sockets.
    processorSockets *int32
    // Processor speed in Hz.
    processorSpeedHz *int64
}
// NewV1beta1HypervisorHostsGetResponse_items_cpuInfo instantiates a new V1beta1HypervisorHostsGetResponse_items_cpuInfo and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items_cpuInfo()(*V1beta1HypervisorHostsGetResponse_items_cpuInfo) {
    m := &V1beta1HypervisorHostsGetResponse_items_cpuInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_items_cpuInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_items_cpuInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items_cpuInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hyperthreadingActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHyperthreadingActive(val)
        }
        return nil
    }
    res["logicalProcessors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogicalProcessors(val)
        }
        return nil
    }
    res["processorCores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCores(val)
        }
        return nil
    }
    res["processorSockets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSockets(val)
        }
        return nil
    }
    res["processorSpeedHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedHz(val)
        }
        return nil
    }
    return res
}
// GetHyperthreadingActive gets the hyperthreadingActive property value. Indicate whether hyperthreading is active or not.
// returns a *bool when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetHyperthreadingActive()(*bool) {
    return m.hyperthreadingActive
}
// GetLogicalProcessors gets the logicalProcessors property value. Number of logical processors.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetLogicalProcessors()(*int32) {
    return m.logicalProcessors
}
// GetProcessorCores gets the processorCores property value. Number of processor cores.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetProcessorCores()(*int32) {
    return m.processorCores
}
// GetProcessorSockets gets the processorSockets property value. Number of processor sockets.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetProcessorSockets()(*int32) {
    return m.processorSockets
}
// GetProcessorSpeedHz gets the processorSpeedHz property value. Processor speed in Hz.
// returns a *int64 when successful
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) GetProcessorSpeedHz()(*int64) {
    return m.processorSpeedHz
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hyperthreadingActive", m.GetHyperthreadingActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("logicalProcessors", m.GetLogicalProcessors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("processorCores", m.GetProcessorCores())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("processorSockets", m.GetProcessorSockets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("processorSpeedHz", m.GetProcessorSpeedHz())
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
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHyperthreadingActive sets the hyperthreadingActive property value. Indicate whether hyperthreading is active or not.
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetHyperthreadingActive(value *bool)() {
    m.hyperthreadingActive = value
}
// SetLogicalProcessors sets the logicalProcessors property value. Number of logical processors.
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetLogicalProcessors(value *int32)() {
    m.logicalProcessors = value
}
// SetProcessorCores sets the processorCores property value. Number of processor cores.
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetProcessorCores(value *int32)() {
    m.processorCores = value
}
// SetProcessorSockets sets the processorSockets property value. Number of processor sockets.
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetProcessorSockets(value *int32)() {
    m.processorSockets = value
}
// SetProcessorSpeedHz sets the processorSpeedHz property value. Processor speed in Hz.
func (m *V1beta1HypervisorHostsGetResponse_items_cpuInfo) SetProcessorSpeedHz(value *int64)() {
    m.processorSpeedHz = value
}
type V1beta1HypervisorHostsGetResponse_items_cpuInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHyperthreadingActive()(*bool)
    GetLogicalProcessors()(*int32)
    GetProcessorCores()(*int32)
    GetProcessorSockets()(*int32)
    GetProcessorSpeedHz()(*int64)
    SetHyperthreadingActive(value *bool)()
    SetLogicalProcessors(value *int32)()
    SetProcessorCores(value *int32)()
    SetProcessorSockets(value *int32)()
    SetProcessorSpeedHz(value *int64)()
}
