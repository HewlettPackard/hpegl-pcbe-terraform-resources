package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo hypervisor cluster performance metrics.
type V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // CPU allocated in mega hertz.
    cpuCapacityInMhz *int64
    // CPU used in mega hertz.
    cpuUsageInMhz *int64
    // Memory allocated in bytes.
    memorySizeInBytes *int64
    // Memory used in mega bytes.
    memoryUsageInMb *int64
    // Storage allocated in bytes.
    totalStorageInBytes *int64
    // Storage used in bytes.
    usedStorageInBytes *int64
}
// NewV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo instantiates a new V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo and sets the default values.
func NewV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo()(*V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) {
    m := &V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuCapacityInMhz gets the cpuCapacityInMhz property value. CPU allocated in mega hertz.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetCpuCapacityInMhz()(*int64) {
    return m.cpuCapacityInMhz
}
// GetCpuUsageInMhz gets the cpuUsageInMhz property value. CPU used in mega hertz.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetCpuUsageInMhz()(*int64) {
    return m.cpuUsageInMhz
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpuCapacityInMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCapacityInMhz(val)
        }
        return nil
    }
    res["cpuUsageInMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuUsageInMhz(val)
        }
        return nil
    }
    res["memorySizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemorySizeInBytes(val)
        }
        return nil
    }
    res["memoryUsageInMb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryUsageInMb(val)
        }
        return nil
    }
    res["totalStorageInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageInBytes(val)
        }
        return nil
    }
    res["usedStorageInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedStorageInBytes(val)
        }
        return nil
    }
    return res
}
// GetMemorySizeInBytes gets the memorySizeInBytes property value. Memory allocated in bytes.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetMemorySizeInBytes()(*int64) {
    return m.memorySizeInBytes
}
// GetMemoryUsageInMb gets the memoryUsageInMb property value. Memory used in mega bytes.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetMemoryUsageInMb()(*int64) {
    return m.memoryUsageInMb
}
// GetTotalStorageInBytes gets the totalStorageInBytes property value. Storage allocated in bytes.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetTotalStorageInBytes()(*int64) {
    return m.totalStorageInBytes
}
// GetUsedStorageInBytes gets the usedStorageInBytes property value. Storage used in bytes.
// returns a *int64 when successful
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) GetUsedStorageInBytes()(*int64) {
    return m.usedStorageInBytes
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("cpuCapacityInMhz", m.GetCpuCapacityInMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("cpuUsageInMhz", m.GetCpuUsageInMhz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("memorySizeInBytes", m.GetMemorySizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("memoryUsageInMb", m.GetMemoryUsageInMb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalStorageInBytes", m.GetTotalStorageInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("usedStorageInBytes", m.GetUsedStorageInBytes())
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
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuCapacityInMhz sets the cpuCapacityInMhz property value. CPU allocated in mega hertz.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetCpuCapacityInMhz(value *int64)() {
    m.cpuCapacityInMhz = value
}
// SetCpuUsageInMhz sets the cpuUsageInMhz property value. CPU used in mega hertz.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetCpuUsageInMhz(value *int64)() {
    m.cpuUsageInMhz = value
}
// SetMemorySizeInBytes sets the memorySizeInBytes property value. Memory allocated in bytes.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetMemorySizeInBytes(value *int64)() {
    m.memorySizeInBytes = value
}
// SetMemoryUsageInMb sets the memoryUsageInMb property value. Memory used in mega bytes.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetMemoryUsageInMb(value *int64)() {
    m.memoryUsageInMb = value
}
// SetTotalStorageInBytes sets the totalStorageInBytes property value. Storage allocated in bytes.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetTotalStorageInBytes(value *int64)() {
    m.totalStorageInBytes = value
}
// SetUsedStorageInBytes sets the usedStorageInBytes property value. Storage used in bytes.
func (m *V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfo) SetUsedStorageInBytes(value *int64)() {
    m.usedStorageInBytes = value
}
type V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuCapacityInMhz()(*int64)
    GetCpuUsageInMhz()(*int64)
    GetMemorySizeInBytes()(*int64)
    GetMemoryUsageInMb()(*int64)
    GetTotalStorageInBytes()(*int64)
    GetUsedStorageInBytes()(*int64)
    SetCpuCapacityInMhz(value *int64)()
    SetCpuUsageInMhz(value *int64)()
    SetMemorySizeInBytes(value *int64)()
    SetMemoryUsageInMb(value *int64)()
    SetTotalStorageInBytes(value *int64)()
    SetUsedStorageInBytes(value *int64)()
}
