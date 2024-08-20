package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo virtual machine performance metrics.
type V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Average read latency.
    averageReadLatency *int64
    // Average write latency.
    averageWriteLatency *int64
    // CPU allocated in mega hertz.
    cpuAllocatedInMhz *int64
    // CPU used in mega hertz.
    cpuUsedInMhz *int64
    // Memory allocated in mega bytes.
    memoryAllocatedInMb *int64
    // Memory used in mega bytes.
    memoryUsedInMb *int64
    // Storage allocated in kilo bytes.
    storageAllocatedInKb *int64
    // Storage used in bytes.
    storageUsedInBytes *int64
    // Total read IOPS.
    totalReadIops *int64
    // Total write IOPS.
    totalWriteIops *int64
}
// NewV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo instantiates a new V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo()(*V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) {
    m := &V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAverageReadLatency gets the averageReadLatency property value. Average read latency.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetAverageReadLatency()(*int64) {
    return m.averageReadLatency
}
// GetAverageWriteLatency gets the averageWriteLatency property value. Average write latency.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetAverageWriteLatency()(*int64) {
    return m.averageWriteLatency
}
// GetCpuAllocatedInMhz gets the cpuAllocatedInMhz property value. CPU allocated in mega hertz.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetCpuAllocatedInMhz()(*int64) {
    return m.cpuAllocatedInMhz
}
// GetCpuUsedInMhz gets the cpuUsedInMhz property value. CPU used in mega hertz.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetCpuUsedInMhz()(*int64) {
    return m.cpuUsedInMhz
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["averageReadLatency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageReadLatency(val)
        }
        return nil
    }
    res["averageWriteLatency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageWriteLatency(val)
        }
        return nil
    }
    res["cpuAllocatedInMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuAllocatedInMhz(val)
        }
        return nil
    }
    res["cpuUsedInMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuUsedInMhz(val)
        }
        return nil
    }
    res["memoryAllocatedInMb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryAllocatedInMb(val)
        }
        return nil
    }
    res["memoryUsedInMb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryUsedInMb(val)
        }
        return nil
    }
    res["storageAllocatedInKb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAllocatedInKb(val)
        }
        return nil
    }
    res["storageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsedInBytes(val)
        }
        return nil
    }
    res["totalReadIops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalReadIops(val)
        }
        return nil
    }
    res["totalWriteIops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWriteIops(val)
        }
        return nil
    }
    return res
}
// GetMemoryAllocatedInMb gets the memoryAllocatedInMb property value. Memory allocated in mega bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetMemoryAllocatedInMb()(*int64) {
    return m.memoryAllocatedInMb
}
// GetMemoryUsedInMb gets the memoryUsedInMb property value. Memory used in mega bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetMemoryUsedInMb()(*int64) {
    return m.memoryUsedInMb
}
// GetStorageAllocatedInKb gets the storageAllocatedInKb property value. Storage allocated in kilo bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetStorageAllocatedInKb()(*int64) {
    return m.storageAllocatedInKb
}
// GetStorageUsedInBytes gets the storageUsedInBytes property value. Storage used in bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetStorageUsedInBytes()(*int64) {
    return m.storageUsedInBytes
}
// GetTotalReadIops gets the totalReadIops property value. Total read IOPS.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetTotalReadIops()(*int64) {
    return m.totalReadIops
}
// GetTotalWriteIops gets the totalWriteIops property value. Total write IOPS.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) GetTotalWriteIops()(*int64) {
    return m.totalWriteIops
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("averageReadLatency", m.GetAverageReadLatency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("averageWriteLatency", m.GetAverageWriteLatency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("cpuAllocatedInMhz", m.GetCpuAllocatedInMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("cpuUsedInMhz", m.GetCpuUsedInMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("memoryAllocatedInMb", m.GetMemoryAllocatedInMb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("memoryUsedInMb", m.GetMemoryUsedInMb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("storageAllocatedInKb", m.GetStorageAllocatedInKb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("storageUsedInBytes", m.GetStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalReadIops", m.GetTotalReadIops())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalWriteIops", m.GetTotalWriteIops())
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
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAverageReadLatency sets the averageReadLatency property value. Average read latency.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetAverageReadLatency(value *int64)() {
    m.averageReadLatency = value
}
// SetAverageWriteLatency sets the averageWriteLatency property value. Average write latency.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetAverageWriteLatency(value *int64)() {
    m.averageWriteLatency = value
}
// SetCpuAllocatedInMhz sets the cpuAllocatedInMhz property value. CPU allocated in mega hertz.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetCpuAllocatedInMhz(value *int64)() {
    m.cpuAllocatedInMhz = value
}
// SetCpuUsedInMhz sets the cpuUsedInMhz property value. CPU used in mega hertz.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetCpuUsedInMhz(value *int64)() {
    m.cpuUsedInMhz = value
}
// SetMemoryAllocatedInMb sets the memoryAllocatedInMb property value. Memory allocated in mega bytes.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetMemoryAllocatedInMb(value *int64)() {
    m.memoryAllocatedInMb = value
}
// SetMemoryUsedInMb sets the memoryUsedInMb property value. Memory used in mega bytes.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetMemoryUsedInMb(value *int64)() {
    m.memoryUsedInMb = value
}
// SetStorageAllocatedInKb sets the storageAllocatedInKb property value. Storage allocated in kilo bytes.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetStorageAllocatedInKb(value *int64)() {
    m.storageAllocatedInKb = value
}
// SetStorageUsedInBytes sets the storageUsedInBytes property value. Storage used in bytes.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
// SetTotalReadIops sets the totalReadIops property value. Total read IOPS.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetTotalReadIops(value *int64)() {
    m.totalReadIops = value
}
// SetTotalWriteIops sets the totalWriteIops property value. Total write IOPS.
func (m *V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfo) SetTotalWriteIops(value *int64)() {
    m.totalWriteIops = value
}
type V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageReadLatency()(*int64)
    GetAverageWriteLatency()(*int64)
    GetCpuAllocatedInMhz()(*int64)
    GetCpuUsedInMhz()(*int64)
    GetMemoryAllocatedInMb()(*int64)
    GetMemoryUsedInMb()(*int64)
    GetStorageAllocatedInKb()(*int64)
    GetStorageUsedInBytes()(*int64)
    GetTotalReadIops()(*int64)
    GetTotalWriteIops()(*int64)
    SetAverageReadLatency(value *int64)()
    SetAverageWriteLatency(value *int64)()
    SetCpuAllocatedInMhz(value *int64)()
    SetCpuUsedInMhz(value *int64)()
    SetMemoryAllocatedInMb(value *int64)()
    SetMemoryUsedInMb(value *int64)()
    SetStorageAllocatedInKb(value *int64)()
    SetStorageUsedInBytes(value *int64)()
    SetTotalReadIops(value *int64)()
    SetTotalWriteIops(value *int64)()
}
