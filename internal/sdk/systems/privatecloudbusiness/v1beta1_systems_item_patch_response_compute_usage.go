package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_computeUsage system Compute Usage Information.
type V1beta1SystemsItemPatchResponse_computeUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // CPU Capacity in Mhz of system
    cpuCapacityMhz *float64
    // CPU Usage in Mhz by system
    cpuUsedMhz *float64
    // Memory Capacity in Bytes of system
    memoryCapacityBytes *float64
    // Memory Usage in Bytes by system
    memoryUsedBytes *float64
}
// NewV1beta1SystemsItemPatchResponse_computeUsage instantiates a new V1beta1SystemsItemPatchResponse_computeUsage and sets the default values.
func NewV1beta1SystemsItemPatchResponse_computeUsage()(*V1beta1SystemsItemPatchResponse_computeUsage) {
    m := &V1beta1SystemsItemPatchResponse_computeUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_computeUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_computeUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_computeUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuCapacityMhz gets the cpuCapacityMhz property value. CPU Capacity in Mhz of system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetCpuCapacityMhz()(*float64) {
    return m.cpuCapacityMhz
}
// GetCpuUsedMhz gets the cpuUsedMhz property value. CPU Usage in Mhz by system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetCpuUsedMhz()(*float64) {
    return m.cpuUsedMhz
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpuCapacityMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCapacityMhz(val)
        }
        return nil
    }
    res["cpuUsedMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuUsedMhz(val)
        }
        return nil
    }
    res["memoryCapacityBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryCapacityBytes(val)
        }
        return nil
    }
    res["memoryUsedBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryUsedBytes(val)
        }
        return nil
    }
    return res
}
// GetMemoryCapacityBytes gets the memoryCapacityBytes property value. Memory Capacity in Bytes of system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetMemoryCapacityBytes()(*float64) {
    return m.memoryCapacityBytes
}
// GetMemoryUsedBytes gets the memoryUsedBytes property value. Memory Usage in Bytes by system
// returns a *float64 when successful
func (m *V1beta1SystemsItemPatchResponse_computeUsage) GetMemoryUsedBytes()(*float64) {
    return m.memoryUsedBytes
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_computeUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("cpuCapacityMhz", m.GetCpuCapacityMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("cpuUsedMhz", m.GetCpuUsedMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("memoryCapacityBytes", m.GetMemoryCapacityBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("memoryUsedBytes", m.GetMemoryUsedBytes())
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
func (m *V1beta1SystemsItemPatchResponse_computeUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuCapacityMhz sets the cpuCapacityMhz property value. CPU Capacity in Mhz of system
func (m *V1beta1SystemsItemPatchResponse_computeUsage) SetCpuCapacityMhz(value *float64)() {
    m.cpuCapacityMhz = value
}
// SetCpuUsedMhz sets the cpuUsedMhz property value. CPU Usage in Mhz by system
func (m *V1beta1SystemsItemPatchResponse_computeUsage) SetCpuUsedMhz(value *float64)() {
    m.cpuUsedMhz = value
}
// SetMemoryCapacityBytes sets the memoryCapacityBytes property value. Memory Capacity in Bytes of system
func (m *V1beta1SystemsItemPatchResponse_computeUsage) SetMemoryCapacityBytes(value *float64)() {
    m.memoryCapacityBytes = value
}
// SetMemoryUsedBytes sets the memoryUsedBytes property value. Memory Usage in Bytes by system
func (m *V1beta1SystemsItemPatchResponse_computeUsage) SetMemoryUsedBytes(value *float64)() {
    m.memoryUsedBytes = value
}
type V1beta1SystemsItemPatchResponse_computeUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuCapacityMhz()(*float64)
    GetCpuUsedMhz()(*float64)
    GetMemoryCapacityBytes()(*float64)
    GetMemoryUsedBytes()(*float64)
    SetCpuCapacityMhz(value *float64)()
    SetCpuUsedMhz(value *float64)()
    SetMemoryCapacityBytes(value *float64)()
    SetMemoryUsedBytes(value *float64)()
}
