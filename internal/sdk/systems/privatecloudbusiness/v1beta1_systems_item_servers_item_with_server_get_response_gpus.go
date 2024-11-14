package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersItemWithServerGetResponse_gpus gPU information.
type V1beta1SystemsItemServersItemWithServerGetResponse_gpus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Bus/Driver/Function information.
    bdf *string
    // The driverVersion property
    driverVersion *string
    // The memoryFreeBytes property
    memoryFreeBytes *float64
    // The memoryTotalBytes property
    memoryTotalBytes *float64
    // The memoryUsedBytes property
    memoryUsedBytes *float64
    // The name property
    name *string
    // The vbiosVersion property
    vbiosVersion *string
}
// NewV1beta1SystemsItemServersItemWithServerGetResponse_gpus instantiates a new V1beta1SystemsItemServersItemWithServerGetResponse_gpus and sets the default values.
func NewV1beta1SystemsItemServersItemWithServerGetResponse_gpus()(*V1beta1SystemsItemServersItemWithServerGetResponse_gpus) {
    m := &V1beta1SystemsItemServersItemWithServerGetResponse_gpus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersItemWithServerGetResponse_gpusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersItemWithServerGetResponse_gpusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersItemWithServerGetResponse_gpus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBdf gets the bdf property value. Bus/Driver/Function information.
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetBdf()(*string) {
    return m.bdf
}
// GetDriverVersion gets the driverVersion property value. The driverVersion property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetDriverVersion()(*string) {
    return m.driverVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bdf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBdf(val)
        }
        return nil
    }
    res["driverVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverVersion(val)
        }
        return nil
    }
    res["memoryFreeBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryFreeBytes(val)
        }
        return nil
    }
    res["memoryTotalBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryTotalBytes(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["vbiosVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVbiosVersion(val)
        }
        return nil
    }
    return res
}
// GetMemoryFreeBytes gets the memoryFreeBytes property value. The memoryFreeBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetMemoryFreeBytes()(*float64) {
    return m.memoryFreeBytes
}
// GetMemoryTotalBytes gets the memoryTotalBytes property value. The memoryTotalBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetMemoryTotalBytes()(*float64) {
    return m.memoryTotalBytes
}
// GetMemoryUsedBytes gets the memoryUsedBytes property value. The memoryUsedBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetMemoryUsedBytes()(*float64) {
    return m.memoryUsedBytes
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetName()(*string) {
    return m.name
}
// GetVbiosVersion gets the vbiosVersion property value. The vbiosVersion property
// returns a *string when successful
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) GetVbiosVersion()(*string) {
    return m.vbiosVersion
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bdf", m.GetBdf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("driverVersion", m.GetDriverVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("memoryFreeBytes", m.GetMemoryFreeBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("memoryTotalBytes", m.GetMemoryTotalBytes())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vbiosVersion", m.GetVbiosVersion())
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
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBdf sets the bdf property value. Bus/Driver/Function information.
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetBdf(value *string)() {
    m.bdf = value
}
// SetDriverVersion sets the driverVersion property value. The driverVersion property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetDriverVersion(value *string)() {
    m.driverVersion = value
}
// SetMemoryFreeBytes sets the memoryFreeBytes property value. The memoryFreeBytes property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetMemoryFreeBytes(value *float64)() {
    m.memoryFreeBytes = value
}
// SetMemoryTotalBytes sets the memoryTotalBytes property value. The memoryTotalBytes property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetMemoryTotalBytes(value *float64)() {
    m.memoryTotalBytes = value
}
// SetMemoryUsedBytes sets the memoryUsedBytes property value. The memoryUsedBytes property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetMemoryUsedBytes(value *float64)() {
    m.memoryUsedBytes = value
}
// SetName sets the name property value. The name property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetName(value *string)() {
    m.name = value
}
// SetVbiosVersion sets the vbiosVersion property value. The vbiosVersion property
func (m *V1beta1SystemsItemServersItemWithServerGetResponse_gpus) SetVbiosVersion(value *string)() {
    m.vbiosVersion = value
}
type V1beta1SystemsItemServersItemWithServerGetResponse_gpusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBdf()(*string)
    GetDriverVersion()(*string)
    GetMemoryFreeBytes()(*float64)
    GetMemoryTotalBytes()(*float64)
    GetMemoryUsedBytes()(*float64)
    GetName()(*string)
    GetVbiosVersion()(*string)
    SetBdf(value *string)()
    SetDriverVersion(value *string)()
    SetMemoryFreeBytes(value *float64)()
    SetMemoryTotalBytes(value *float64)()
    SetMemoryUsedBytes(value *float64)()
    SetName(value *string)()
    SetVbiosVersion(value *string)()
}
