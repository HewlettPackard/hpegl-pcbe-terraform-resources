package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemStoragePoolsGetResponse_items system's associated storage arrays.
type V1beta1SystemsItemStoragePoolsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The capacity property
    capacity *float64
    // The cloneRatio property
    cloneRatio *float64
    // The compressionRatio property
    compressionRatio *float64
    // The dataReductionRatio property
    dataReductionRatio *float64
    // The dedupeRatio property
    dedupeRatio *float64
    // I/O operations per second observed for this storage pool.
    iops *float64
    // Latency (milliseconds) observed for this storage pool.
    latency *float64
    // The savings property
    savings *float64
    // The savingsClone property
    savingsClone *float64
    // The savingsCompression property
    savingsCompression *float64
    // The savingsDataReduction property
    savingsDataReduction *float64
    // The savingsDedupe property
    savingsDedupe *float64
    // The savingsRatio property
    savingsRatio *float64
    // The savingsVolThinProvisioning property
    savingsVolThinProvisioning *float64
    // Throughput (B/s) observed for this storage pool.
    throughput *float64
    // The uncompressedSnapUsageBytes property
    uncompressedSnapUsageBytes *float64
    // The uncompressedVolUsageBytes property
    uncompressedVolUsageBytes *float64
    // The usage property
    usage *float64
    // Number of VMFS VMs using this storage pool.
    vmfsCount *float64
    // The volThinProvisioningRatio property
    volThinProvisioningRatio *float64
    // Number of VVOL VMs using this storage pool.
    vvolCount *float64
}
// NewV1beta1SystemsItemStoragePoolsGetResponse_items instantiates a new V1beta1SystemsItemStoragePoolsGetResponse_items and sets the default values.
func NewV1beta1SystemsItemStoragePoolsGetResponse_items()(*V1beta1SystemsItemStoragePoolsGetResponse_items) {
    m := &V1beta1SystemsItemStoragePoolsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStoragePoolsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStoragePoolsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStoragePoolsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapacity gets the capacity property value. The capacity property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetCapacity()(*float64) {
    return m.capacity
}
// GetCloneRatio gets the cloneRatio property value. The cloneRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetCloneRatio()(*float64) {
    return m.cloneRatio
}
// GetCompressionRatio gets the compressionRatio property value. The compressionRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetCompressionRatio()(*float64) {
    return m.compressionRatio
}
// GetDataReductionRatio gets the dataReductionRatio property value. The dataReductionRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetDataReductionRatio()(*float64) {
    return m.dataReductionRatio
}
// GetDedupeRatio gets the dedupeRatio property value. The dedupeRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetDedupeRatio()(*float64) {
    return m.dedupeRatio
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["cloneRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloneRatio(val)
        }
        return nil
    }
    res["compressionRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompressionRatio(val)
        }
        return nil
    }
    res["dataReductionRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataReductionRatio(val)
        }
        return nil
    }
    res["dedupeRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDedupeRatio(val)
        }
        return nil
    }
    res["iops"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIops(val)
        }
        return nil
    }
    res["latency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatency(val)
        }
        return nil
    }
    res["savings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavings(val)
        }
        return nil
    }
    res["savingsClone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsClone(val)
        }
        return nil
    }
    res["savingsCompression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsCompression(val)
        }
        return nil
    }
    res["savingsDataReduction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsDataReduction(val)
        }
        return nil
    }
    res["savingsDedupe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsDedupe(val)
        }
        return nil
    }
    res["savingsRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsRatio(val)
        }
        return nil
    }
    res["savingsVolThinProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSavingsVolThinProvisioning(val)
        }
        return nil
    }
    res["throughput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThroughput(val)
        }
        return nil
    }
    res["uncompressedSnapUsageBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUncompressedSnapUsageBytes(val)
        }
        return nil
    }
    res["uncompressedVolUsageBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUncompressedVolUsageBytes(val)
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val)
        }
        return nil
    }
    res["vmfsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmfsCount(val)
        }
        return nil
    }
    res["volThinProvisioningRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolThinProvisioningRatio(val)
        }
        return nil
    }
    res["vvolCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVvolCount(val)
        }
        return nil
    }
    return res
}
// GetIops gets the iops property value. I/O operations per second observed for this storage pool.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetIops()(*float64) {
    return m.iops
}
// GetLatency gets the latency property value. Latency (milliseconds) observed for this storage pool.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetLatency()(*float64) {
    return m.latency
}
// GetSavings gets the savings property value. The savings property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavings()(*float64) {
    return m.savings
}
// GetSavingsClone gets the savingsClone property value. The savingsClone property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsClone()(*float64) {
    return m.savingsClone
}
// GetSavingsCompression gets the savingsCompression property value. The savingsCompression property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsCompression()(*float64) {
    return m.savingsCompression
}
// GetSavingsDataReduction gets the savingsDataReduction property value. The savingsDataReduction property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsDataReduction()(*float64) {
    return m.savingsDataReduction
}
// GetSavingsDedupe gets the savingsDedupe property value. The savingsDedupe property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsDedupe()(*float64) {
    return m.savingsDedupe
}
// GetSavingsRatio gets the savingsRatio property value. The savingsRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsRatio()(*float64) {
    return m.savingsRatio
}
// GetSavingsVolThinProvisioning gets the savingsVolThinProvisioning property value. The savingsVolThinProvisioning property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetSavingsVolThinProvisioning()(*float64) {
    return m.savingsVolThinProvisioning
}
// GetThroughput gets the throughput property value. Throughput (B/s) observed for this storage pool.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetThroughput()(*float64) {
    return m.throughput
}
// GetUncompressedSnapUsageBytes gets the uncompressedSnapUsageBytes property value. The uncompressedSnapUsageBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetUncompressedSnapUsageBytes()(*float64) {
    return m.uncompressedSnapUsageBytes
}
// GetUncompressedVolUsageBytes gets the uncompressedVolUsageBytes property value. The uncompressedVolUsageBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetUncompressedVolUsageBytes()(*float64) {
    return m.uncompressedVolUsageBytes
}
// GetUsage gets the usage property value. The usage property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetUsage()(*float64) {
    return m.usage
}
// GetVmfsCount gets the vmfsCount property value. Number of VMFS VMs using this storage pool.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetVmfsCount()(*float64) {
    return m.vmfsCount
}
// GetVolThinProvisioningRatio gets the volThinProvisioningRatio property value. The volThinProvisioningRatio property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetVolThinProvisioningRatio()(*float64) {
    return m.volThinProvisioningRatio
}
// GetVvolCount gets the vvolCount property value. Number of VVOL VMs using this storage pool.
// returns a *float64 when successful
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) GetVvolCount()(*float64) {
    return m.vvolCount
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("cloneRatio", m.GetCloneRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("compressionRatio", m.GetCompressionRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("dataReductionRatio", m.GetDataReductionRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("dedupeRatio", m.GetDedupeRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("iops", m.GetIops())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latency", m.GetLatency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savings", m.GetSavings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsClone", m.GetSavingsClone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsCompression", m.GetSavingsCompression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsDataReduction", m.GetSavingsDataReduction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsDedupe", m.GetSavingsDedupe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsRatio", m.GetSavingsRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("savingsVolThinProvisioning", m.GetSavingsVolThinProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("throughput", m.GetThroughput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("uncompressedSnapUsageBytes", m.GetUncompressedSnapUsageBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("uncompressedVolUsageBytes", m.GetUncompressedVolUsageBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("usage", m.GetUsage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("vmfsCount", m.GetVmfsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("volThinProvisioningRatio", m.GetVolThinProvisioningRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("vvolCount", m.GetVvolCount())
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
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapacity sets the capacity property value. The capacity property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetCapacity(value *float64)() {
    m.capacity = value
}
// SetCloneRatio sets the cloneRatio property value. The cloneRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetCloneRatio(value *float64)() {
    m.cloneRatio = value
}
// SetCompressionRatio sets the compressionRatio property value. The compressionRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetCompressionRatio(value *float64)() {
    m.compressionRatio = value
}
// SetDataReductionRatio sets the dataReductionRatio property value. The dataReductionRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetDataReductionRatio(value *float64)() {
    m.dataReductionRatio = value
}
// SetDedupeRatio sets the dedupeRatio property value. The dedupeRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetDedupeRatio(value *float64)() {
    m.dedupeRatio = value
}
// SetIops sets the iops property value. I/O operations per second observed for this storage pool.
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetIops(value *float64)() {
    m.iops = value
}
// SetLatency sets the latency property value. Latency (milliseconds) observed for this storage pool.
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetLatency(value *float64)() {
    m.latency = value
}
// SetSavings sets the savings property value. The savings property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavings(value *float64)() {
    m.savings = value
}
// SetSavingsClone sets the savingsClone property value. The savingsClone property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsClone(value *float64)() {
    m.savingsClone = value
}
// SetSavingsCompression sets the savingsCompression property value. The savingsCompression property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsCompression(value *float64)() {
    m.savingsCompression = value
}
// SetSavingsDataReduction sets the savingsDataReduction property value. The savingsDataReduction property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsDataReduction(value *float64)() {
    m.savingsDataReduction = value
}
// SetSavingsDedupe sets the savingsDedupe property value. The savingsDedupe property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsDedupe(value *float64)() {
    m.savingsDedupe = value
}
// SetSavingsRatio sets the savingsRatio property value. The savingsRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsRatio(value *float64)() {
    m.savingsRatio = value
}
// SetSavingsVolThinProvisioning sets the savingsVolThinProvisioning property value. The savingsVolThinProvisioning property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetSavingsVolThinProvisioning(value *float64)() {
    m.savingsVolThinProvisioning = value
}
// SetThroughput sets the throughput property value. Throughput (B/s) observed for this storage pool.
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetThroughput(value *float64)() {
    m.throughput = value
}
// SetUncompressedSnapUsageBytes sets the uncompressedSnapUsageBytes property value. The uncompressedSnapUsageBytes property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetUncompressedSnapUsageBytes(value *float64)() {
    m.uncompressedSnapUsageBytes = value
}
// SetUncompressedVolUsageBytes sets the uncompressedVolUsageBytes property value. The uncompressedVolUsageBytes property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetUncompressedVolUsageBytes(value *float64)() {
    m.uncompressedVolUsageBytes = value
}
// SetUsage sets the usage property value. The usage property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetUsage(value *float64)() {
    m.usage = value
}
// SetVmfsCount sets the vmfsCount property value. Number of VMFS VMs using this storage pool.
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetVmfsCount(value *float64)() {
    m.vmfsCount = value
}
// SetVolThinProvisioningRatio sets the volThinProvisioningRatio property value. The volThinProvisioningRatio property
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetVolThinProvisioningRatio(value *float64)() {
    m.volThinProvisioningRatio = value
}
// SetVvolCount sets the vvolCount property value. Number of VVOL VMs using this storage pool.
func (m *V1beta1SystemsItemStoragePoolsGetResponse_items) SetVvolCount(value *float64)() {
    m.vvolCount = value
}
type V1beta1SystemsItemStoragePoolsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacity()(*float64)
    GetCloneRatio()(*float64)
    GetCompressionRatio()(*float64)
    GetDataReductionRatio()(*float64)
    GetDedupeRatio()(*float64)
    GetIops()(*float64)
    GetLatency()(*float64)
    GetSavings()(*float64)
    GetSavingsClone()(*float64)
    GetSavingsCompression()(*float64)
    GetSavingsDataReduction()(*float64)
    GetSavingsDedupe()(*float64)
    GetSavingsRatio()(*float64)
    GetSavingsVolThinProvisioning()(*float64)
    GetThroughput()(*float64)
    GetUncompressedSnapUsageBytes()(*float64)
    GetUncompressedVolUsageBytes()(*float64)
    GetUsage()(*float64)
    GetVmfsCount()(*float64)
    GetVolThinProvisioningRatio()(*float64)
    GetVvolCount()(*float64)
    SetCapacity(value *float64)()
    SetCloneRatio(value *float64)()
    SetCompressionRatio(value *float64)()
    SetDataReductionRatio(value *float64)()
    SetDedupeRatio(value *float64)()
    SetIops(value *float64)()
    SetLatency(value *float64)()
    SetSavings(value *float64)()
    SetSavingsClone(value *float64)()
    SetSavingsCompression(value *float64)()
    SetSavingsDataReduction(value *float64)()
    SetSavingsDedupe(value *float64)()
    SetSavingsRatio(value *float64)()
    SetSavingsVolThinProvisioning(value *float64)()
    SetThroughput(value *float64)()
    SetUncompressedSnapUsageBytes(value *float64)()
    SetUncompressedVolUsageBytes(value *float64)()
    SetUsage(value *float64)()
    SetVmfsCount(value *float64)()
    SetVolThinProvisioningRatio(value *float64)()
    SetVvolCount(value *float64)()
}
