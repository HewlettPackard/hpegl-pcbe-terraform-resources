package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether auto recovery is supported.
    autoRecoverySupported *bool
    // Indicates whether the instance is a bare metal instance type.
    bareMetal *bool
    // Indicates whether the instance type is a burstable performance instance type.
    burstablePerformanceSupported *bool
    // Indicates whether the instance type is current generation.
    currentGeneration *bool
    // Indicates whether Dedicated Hosts are supported on the instance type.
    dedicatedHostsSupported *bool
    // The Amazon EBS settings for the instance type.
    ebsInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable
    // Indicates whether the instance type is eligible for the free tier.
    freeTierEligible *bool
    // Indicates whether On-Demand hibernation is supported.
    hibernationSupported *bool
    // The hypervisor type of the image.
    hypervisor *string
    // Indicates whether instance storage is supported.
    instanceStorageSupported *bool
    // The instance type.
    instanceType *string
    // The memory for the instance type.
    memoryInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable
    // The network settings for the instance type.
    networkInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable
    // The placement group settings for the instance type.
    placementGroupInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable
    // Processor information.
    processorInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable
    // Cloud provider region.
    region *string
    // The vCPU configurations for the instance type.
    vCpuInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable
}
// NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 instantiates a new V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 and sets the default values.
func NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1()(*V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) {
    m := &V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoRecoverySupported gets the autoRecoverySupported property value. Indicates whether auto recovery is supported.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetAutoRecoverySupported()(*bool) {
    return m.autoRecoverySupported
}
// GetBareMetal gets the bareMetal property value. Indicates whether the instance is a bare metal instance type.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetBareMetal()(*bool) {
    return m.bareMetal
}
// GetBurstablePerformanceSupported gets the burstablePerformanceSupported property value. Indicates whether the instance type is a burstable performance instance type.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetBurstablePerformanceSupported()(*bool) {
    return m.burstablePerformanceSupported
}
// GetCurrentGeneration gets the currentGeneration property value. Indicates whether the instance type is current generation.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetCurrentGeneration()(*bool) {
    return m.currentGeneration
}
// GetDedicatedHostsSupported gets the dedicatedHostsSupported property value. Indicates whether Dedicated Hosts are supported on the instance type.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetDedicatedHostsSupported()(*bool) {
    return m.dedicatedHostsSupported
}
// GetEbsInfo gets the ebsInfo property value. The Amazon EBS settings for the instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetEbsInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable) {
    return m.ebsInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["autoRecoverySupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoRecoverySupported(val)
        }
        return nil
    }
    res["bareMetal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBareMetal(val)
        }
        return nil
    }
    res["burstablePerformanceSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBurstablePerformanceSupported(val)
        }
        return nil
    }
    res["currentGeneration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentGeneration(val)
        }
        return nil
    }
    res["dedicatedHostsSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDedicatedHostsSupported(val)
        }
        return nil
    }
    res["ebsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEbsInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable))
        }
        return nil
    }
    res["freeTierEligible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeTierEligible(val)
        }
        return nil
    }
    res["hibernationSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHibernationSupported(val)
        }
        return nil
    }
    res["hypervisor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisor(val)
        }
        return nil
    }
    res["instanceStorageSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceStorageSupported(val)
        }
        return nil
    }
    res["instanceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceType(val)
        }
        return nil
    }
    res["memoryInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable))
        }
        return nil
    }
    res["networkInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable))
        }
        return nil
    }
    res["placementGroupInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlacementGroupInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable))
        }
        return nil
    }
    res["processorInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable))
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["vCpuInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVCpuInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable))
        }
        return nil
    }
    return res
}
// GetFreeTierEligible gets the freeTierEligible property value. Indicates whether the instance type is eligible for the free tier.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetFreeTierEligible()(*bool) {
    return m.freeTierEligible
}
// GetHibernationSupported gets the hibernationSupported property value. Indicates whether On-Demand hibernation is supported.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetHibernationSupported()(*bool) {
    return m.hibernationSupported
}
// GetHypervisor gets the hypervisor property value. The hypervisor type of the image.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetHypervisor()(*string) {
    return m.hypervisor
}
// GetInstanceStorageSupported gets the instanceStorageSupported property value. Indicates whether instance storage is supported.
// returns a *bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetInstanceStorageSupported()(*bool) {
    return m.instanceStorageSupported
}
// GetInstanceType gets the instanceType property value. The instance type.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetInstanceType()(*string) {
    return m.instanceType
}
// GetMemoryInfo gets the memoryInfo property value. The memory for the instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetMemoryInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable) {
    return m.memoryInfo
}
// GetNetworkInfo gets the networkInfo property value. The network settings for the instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetNetworkInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable) {
    return m.networkInfo
}
// GetPlacementGroupInfo gets the placementGroupInfo property value. The placement group settings for the instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetPlacementGroupInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable) {
    return m.placementGroupInfo
}
// GetProcessorInfo gets the processorInfo property value. Processor information.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetProcessorInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable) {
    return m.processorInfo
}
// GetRegion gets the region property value. Cloud provider region.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetRegion()(*string) {
    return m.region
}
// GetVCpuInfo gets the vCpuInfo property value. The vCPU configurations for the instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) GetVCpuInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable) {
    return m.vCpuInfo
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("autoRecoverySupported", m.GetAutoRecoverySupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("bareMetal", m.GetBareMetal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("burstablePerformanceSupported", m.GetBurstablePerformanceSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("currentGeneration", m.GetCurrentGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dedicatedHostsSupported", m.GetDedicatedHostsSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ebsInfo", m.GetEbsInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("freeTierEligible", m.GetFreeTierEligible())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hibernationSupported", m.GetHibernationSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hypervisor", m.GetHypervisor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("instanceStorageSupported", m.GetInstanceStorageSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceType", m.GetInstanceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("memoryInfo", m.GetMemoryInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("networkInfo", m.GetNetworkInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("placementGroupInfo", m.GetPlacementGroupInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("processorInfo", m.GetProcessorInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vCpuInfo", m.GetVCpuInfo())
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
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoRecoverySupported sets the autoRecoverySupported property value. Indicates whether auto recovery is supported.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetAutoRecoverySupported(value *bool)() {
    m.autoRecoverySupported = value
}
// SetBareMetal sets the bareMetal property value. Indicates whether the instance is a bare metal instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetBareMetal(value *bool)() {
    m.bareMetal = value
}
// SetBurstablePerformanceSupported sets the burstablePerformanceSupported property value. Indicates whether the instance type is a burstable performance instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetBurstablePerformanceSupported(value *bool)() {
    m.burstablePerformanceSupported = value
}
// SetCurrentGeneration sets the currentGeneration property value. Indicates whether the instance type is current generation.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetCurrentGeneration(value *bool)() {
    m.currentGeneration = value
}
// SetDedicatedHostsSupported sets the dedicatedHostsSupported property value. Indicates whether Dedicated Hosts are supported on the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetDedicatedHostsSupported(value *bool)() {
    m.dedicatedHostsSupported = value
}
// SetEbsInfo sets the ebsInfo property value. The Amazon EBS settings for the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetEbsInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable)() {
    m.ebsInfo = value
}
// SetFreeTierEligible sets the freeTierEligible property value. Indicates whether the instance type is eligible for the free tier.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetFreeTierEligible(value *bool)() {
    m.freeTierEligible = value
}
// SetHibernationSupported sets the hibernationSupported property value. Indicates whether On-Demand hibernation is supported.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetHibernationSupported(value *bool)() {
    m.hibernationSupported = value
}
// SetHypervisor sets the hypervisor property value. The hypervisor type of the image.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetHypervisor(value *string)() {
    m.hypervisor = value
}
// SetInstanceStorageSupported sets the instanceStorageSupported property value. Indicates whether instance storage is supported.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetInstanceStorageSupported(value *bool)() {
    m.instanceStorageSupported = value
}
// SetInstanceType sets the instanceType property value. The instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetInstanceType(value *string)() {
    m.instanceType = value
}
// SetMemoryInfo sets the memoryInfo property value. The memory for the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetMemoryInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable)() {
    m.memoryInfo = value
}
// SetNetworkInfo sets the networkInfo property value. The network settings for the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetNetworkInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable)() {
    m.networkInfo = value
}
// SetPlacementGroupInfo sets the placementGroupInfo property value. The placement group settings for the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetPlacementGroupInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable)() {
    m.placementGroupInfo = value
}
// SetProcessorInfo sets the processorInfo property value. Processor information.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetProcessorInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable)() {
    m.processorInfo = value
}
// SetRegion sets the region property value. Cloud provider region.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetRegion(value *string)() {
    m.region = value
}
// SetVCpuInfo sets the vCpuInfo property value. The vCPU configurations for the instance type.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1) SetVCpuInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable)() {
    m.vCpuInfo = value
}
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoRecoverySupported()(*bool)
    GetBareMetal()(*bool)
    GetBurstablePerformanceSupported()(*bool)
    GetCurrentGeneration()(*bool)
    GetDedicatedHostsSupported()(*bool)
    GetEbsInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable)
    GetFreeTierEligible()(*bool)
    GetHibernationSupported()(*bool)
    GetHypervisor()(*string)
    GetInstanceStorageSupported()(*bool)
    GetInstanceType()(*string)
    GetMemoryInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable)
    GetNetworkInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable)
    GetPlacementGroupInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable)
    GetProcessorInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable)
    GetRegion()(*string)
    GetVCpuInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable)
    SetAutoRecoverySupported(value *bool)()
    SetBareMetal(value *bool)()
    SetBurstablePerformanceSupported(value *bool)()
    SetCurrentGeneration(value *bool)()
    SetDedicatedHostsSupported(value *bool)()
    SetEbsInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_ebsInfoable)()
    SetFreeTierEligible(value *bool)()
    SetHibernationSupported(value *bool)()
    SetHypervisor(value *string)()
    SetInstanceStorageSupported(value *bool)()
    SetInstanceType(value *string)()
    SetMemoryInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_memoryInfoable)()
    SetNetworkInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_networkInfoable)()
    SetPlacementGroupInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfoable)()
    SetProcessorInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfoable)()
    SetRegion(value *string)()
    SetVCpuInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1_vCpuInfoable)()
}
