package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo the optimized EBS performance for the instance type.
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
    baselineBandwidthInMbps *int32
    // The baseline input/output storage operations per seconds for an EBS-optimized instance type.
    baselineIops *int32
    // The baseline throughput performance for an EBS-optimized instance type, in MB/s.
    baselineThroughputInMbps *float64
    // The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
    maximumBandwidthInMbps *int32
    // The maximum input/output storage operations per second for an EBS-optimized instance type.
    maximumIops *int32
    // The maximum throughput performance for an EBS-optimized instance type, in MB/s.
    maximumThroughputInMbps *float64
}
// NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo instantiates a new V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo()(*V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) {
    m := &V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaselineBandwidthInMbps gets the baselineBandwidthInMbps property value. The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetBaselineBandwidthInMbps()(*int32) {
    return m.baselineBandwidthInMbps
}
// GetBaselineIops gets the baselineIops property value. The baseline input/output storage operations per seconds for an EBS-optimized instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetBaselineIops()(*int32) {
    return m.baselineIops
}
// GetBaselineThroughputInMbps gets the baselineThroughputInMbps property value. The baseline throughput performance for an EBS-optimized instance type, in MB/s.
// returns a *float64 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetBaselineThroughputInMbps()(*float64) {
    return m.baselineThroughputInMbps
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["baselineBandwidthInMbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaselineBandwidthInMbps(val)
        }
        return nil
    }
    res["baselineIops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaselineIops(val)
        }
        return nil
    }
    res["baselineThroughputInMbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaselineThroughputInMbps(val)
        }
        return nil
    }
    res["maximumBandwidthInMbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumBandwidthInMbps(val)
        }
        return nil
    }
    res["maximumIops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumIops(val)
        }
        return nil
    }
    res["maximumThroughputInMbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumThroughputInMbps(val)
        }
        return nil
    }
    return res
}
// GetMaximumBandwidthInMbps gets the maximumBandwidthInMbps property value. The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetMaximumBandwidthInMbps()(*int32) {
    return m.maximumBandwidthInMbps
}
// GetMaximumIops gets the maximumIops property value. The maximum input/output storage operations per second for an EBS-optimized instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetMaximumIops()(*int32) {
    return m.maximumIops
}
// GetMaximumThroughputInMbps gets the maximumThroughputInMbps property value. The maximum throughput performance for an EBS-optimized instance type, in MB/s.
// returns a *float64 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) GetMaximumThroughputInMbps()(*float64) {
    return m.maximumThroughputInMbps
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("baselineBandwidthInMbps", m.GetBaselineBandwidthInMbps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("baselineIops", m.GetBaselineIops())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("baselineThroughputInMbps", m.GetBaselineThroughputInMbps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximumBandwidthInMbps", m.GetMaximumBandwidthInMbps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximumIops", m.GetMaximumIops())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximumThroughputInMbps", m.GetMaximumThroughputInMbps())
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
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaselineBandwidthInMbps sets the baselineBandwidthInMbps property value. The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetBaselineBandwidthInMbps(value *int32)() {
    m.baselineBandwidthInMbps = value
}
// SetBaselineIops sets the baselineIops property value. The baseline input/output storage operations per seconds for an EBS-optimized instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetBaselineIops(value *int32)() {
    m.baselineIops = value
}
// SetBaselineThroughputInMbps sets the baselineThroughputInMbps property value. The baseline throughput performance for an EBS-optimized instance type, in MB/s.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetBaselineThroughputInMbps(value *float64)() {
    m.baselineThroughputInMbps = value
}
// SetMaximumBandwidthInMbps sets the maximumBandwidthInMbps property value. The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetMaximumBandwidthInMbps(value *int32)() {
    m.maximumBandwidthInMbps = value
}
// SetMaximumIops sets the maximumIops property value. The maximum input/output storage operations per second for an EBS-optimized instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetMaximumIops(value *int32)() {
    m.maximumIops = value
}
// SetMaximumThroughputInMbps sets the maximumThroughputInMbps property value. The maximum throughput performance for an EBS-optimized instance type, in MB/s.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfo) SetMaximumThroughputInMbps(value *float64)() {
    m.maximumThroughputInMbps = value
}
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfo_ebsOptimizedInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaselineBandwidthInMbps()(*int32)
    GetBaselineIops()(*int32)
    GetBaselineThroughputInMbps()(*float64)
    GetMaximumBandwidthInMbps()(*int32)
    GetMaximumIops()(*int32)
    GetMaximumThroughputInMbps()(*float64)
    SetBaselineBandwidthInMbps(value *int32)()
    SetBaselineIops(value *int32)()
    SetBaselineThroughputInMbps(value *float64)()
    SetMaximumBandwidthInMbps(value *int32)()
    SetMaximumIops(value *int32)()
    SetMaximumThroughputInMbps(value *float64)()
}
