package privatecloudbusiness

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse details of the system with select information.
type V1beta1SystemsItemGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Associated Resource Information of system.
    associatedResourceCounts V1beta1SystemsItemGetResponse_associatedResourceCountsable
    // List of compute clusters in the system.
    computeClusters []V1beta1SystemsItemGetResponse_computeClustersable
    // System Compute Usage Information.
    computeUsage V1beta1SystemsItemGetResponse_computeUsageable
    // Status of last run of configuration analysis job.
    configAnalysisStatus V1beta1SystemsItemGetResponse_configAnalysisStatusable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Aspects of system health.Deduced health of storage subsystem based on associated arrays, controllers, disks,shelves, power supply, network interfaces, fan and temperature sensors.Aggregated health of servers based on health statuses of multiple servers in the system.Aggregated health of network based on health statues of multiple switches used in the system.Deduced/Aggregated overall health of the system based on storage, servers and networking health.
    health V1beta1SystemsItemGetResponse_healthable
    // List of hypervisor clusters in the system. To be deprecated soon.
    hypervisorClusters []V1beta1SystemsItemGetResponse_hypervisorClustersable
    // An identifier for the resource, usually a UUID.
    id *string
    // System Location Information.
    location V1beta1SystemsItemGetResponse_locationable
    // A system specified name for the resource.
    name *string
    // An identifier of the on prem agent for the system.
    onPremAgentId *string
    // The self reference for this resource.
    resourceUri *string
    // system software information.
    softwareInfo V1beta1SystemsItemGetResponse_softwareInfoable
    // Storage Information of system.
    storageSystem V1beta1SystemsItemGetResponse_storageSystemable
    // Storage Usage Information of system.
    storageUsage V1beta1SystemsItemGetResponse_storageUsageable
    // List of system virtual machine information.
    systemVms []V1beta1SystemsItemGetResponse_systemVmsable
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SystemsItemGetResponse instantiates a new V1beta1SystemsItemGetResponse and sets the default values.
func NewV1beta1SystemsItemGetResponse()(*V1beta1SystemsItemGetResponse) {
    m := &V1beta1SystemsItemGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssociatedResourceCounts gets the associatedResourceCounts property value. Associated Resource Information of system.
// returns a V1beta1SystemsItemGetResponse_associatedResourceCountsable when successful
func (m *V1beta1SystemsItemGetResponse) GetAssociatedResourceCounts()(V1beta1SystemsItemGetResponse_associatedResourceCountsable) {
    return m.associatedResourceCounts
}
// GetComputeClusters gets the computeClusters property value. List of compute clusters in the system.
// returns a []V1beta1SystemsItemGetResponse_computeClustersable when successful
func (m *V1beta1SystemsItemGetResponse) GetComputeClusters()([]V1beta1SystemsItemGetResponse_computeClustersable) {
    return m.computeClusters
}
// GetComputeUsage gets the computeUsage property value. System Compute Usage Information.
// returns a V1beta1SystemsItemGetResponse_computeUsageable when successful
func (m *V1beta1SystemsItemGetResponse) GetComputeUsage()(V1beta1SystemsItemGetResponse_computeUsageable) {
    return m.computeUsage
}
// GetConfigAnalysisStatus gets the configAnalysisStatus property value. Status of last run of configuration analysis job.
// returns a V1beta1SystemsItemGetResponse_configAnalysisStatusable when successful
func (m *V1beta1SystemsItemGetResponse) GetConfigAnalysisStatus()(V1beta1SystemsItemGetResponse_configAnalysisStatusable) {
    return m.configAnalysisStatus
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["associatedResourceCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_associatedResourceCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedResourceCounts(val.(V1beta1SystemsItemGetResponse_associatedResourceCountsable))
        }
        return nil
    }
    res["computeClusters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemGetResponse_computeClustersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemGetResponse_computeClustersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemGetResponse_computeClustersable)
                }
            }
            m.SetComputeClusters(res)
        }
        return nil
    }
    res["computeUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_computeUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeUsage(val.(V1beta1SystemsItemGetResponse_computeUsageable))
        }
        return nil
    }
    res["configAnalysisStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_configAnalysisStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigAnalysisStatus(val.(V1beta1SystemsItemGetResponse_configAnalysisStatusable))
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["health"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_healthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(V1beta1SystemsItemGetResponse_healthable))
        }
        return nil
    }
    res["hypervisorClusters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemGetResponse_hypervisorClustersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemGetResponse_hypervisorClustersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemGetResponse_hypervisorClustersable)
                }
            }
            m.SetHypervisorClusters(res)
        }
        return nil
    }
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
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_locationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(V1beta1SystemsItemGetResponse_locationable))
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
    res["onPremAgentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremAgentId(val)
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
    res["softwareInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_softwareInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareInfo(val.(V1beta1SystemsItemGetResponse_softwareInfoable))
        }
        return nil
    }
    res["storageSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_storageSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSystem(val.(V1beta1SystemsItemGetResponse_storageSystemable))
        }
        return nil
    }
    res["storageUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemGetResponse_storageUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsage(val.(V1beta1SystemsItemGetResponse_storageUsageable))
        }
        return nil
    }
    res["systemVms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemGetResponse_systemVmsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemGetResponse_systemVmsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemGetResponse_systemVmsable)
                }
            }
            m.SetSystemVms(res)
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
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1SystemsItemGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHealth gets the health property value. Aspects of system health.Deduced health of storage subsystem based on associated arrays, controllers, disks,shelves, power supply, network interfaces, fan and temperature sensors.Aggregated health of servers based on health statuses of multiple servers in the system.Aggregated health of network based on health statues of multiple switches used in the system.Deduced/Aggregated overall health of the system based on storage, servers and networking health.
// returns a V1beta1SystemsItemGetResponse_healthable when successful
func (m *V1beta1SystemsItemGetResponse) GetHealth()(V1beta1SystemsItemGetResponse_healthable) {
    return m.health
}
// GetHypervisorClusters gets the hypervisorClusters property value. List of hypervisor clusters in the system. To be deprecated soon.
// returns a []V1beta1SystemsItemGetResponse_hypervisorClustersable when successful
func (m *V1beta1SystemsItemGetResponse) GetHypervisorClusters()([]V1beta1SystemsItemGetResponse_hypervisorClustersable) {
    return m.hypervisorClusters
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetId()(*string) {
    return m.id
}
// GetLocation gets the location property value. System Location Information.
// returns a V1beta1SystemsItemGetResponse_locationable when successful
func (m *V1beta1SystemsItemGetResponse) GetLocation()(V1beta1SystemsItemGetResponse_locationable) {
    return m.location
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetName()(*string) {
    return m.name
}
// GetOnPremAgentId gets the onPremAgentId property value. An identifier of the on prem agent for the system.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetOnPremAgentId()(*string) {
    return m.onPremAgentId
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSoftwareInfo gets the softwareInfo property value. system software information.
// returns a V1beta1SystemsItemGetResponse_softwareInfoable when successful
func (m *V1beta1SystemsItemGetResponse) GetSoftwareInfo()(V1beta1SystemsItemGetResponse_softwareInfoable) {
    return m.softwareInfo
}
// GetStorageSystem gets the storageSystem property value. Storage Information of system.
// returns a V1beta1SystemsItemGetResponse_storageSystemable when successful
func (m *V1beta1SystemsItemGetResponse) GetStorageSystem()(V1beta1SystemsItemGetResponse_storageSystemable) {
    return m.storageSystem
}
// GetStorageUsage gets the storageUsage property value. Storage Usage Information of system.
// returns a V1beta1SystemsItemGetResponse_storageUsageable when successful
func (m *V1beta1SystemsItemGetResponse) GetStorageUsage()(V1beta1SystemsItemGetResponse_storageUsageable) {
    return m.storageUsage
}
// GetSystemVms gets the systemVms property value. List of system virtual machine information.
// returns a []V1beta1SystemsItemGetResponse_systemVmsable when successful
func (m *V1beta1SystemsItemGetResponse) GetSystemVms()([]V1beta1SystemsItemGetResponse_systemVmsable) {
    return m.systemVms
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SystemsItemGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SystemsItemGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("associatedResourceCounts", m.GetAssociatedResourceCounts())
        if err != nil {
            return err
        }
    }
    if m.GetComputeClusters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComputeClusters()))
        for i, v := range m.GetComputeClusters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("computeClusters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("computeUsage", m.GetComputeUsage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("configAnalysisStatus", m.GetConfigAnalysisStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("health", m.GetHealth())
        if err != nil {
            return err
        }
    }
    if m.GetHypervisorClusters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHypervisorClusters()))
        for i, v := range m.GetHypervisorClusters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hypervisorClusters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("location", m.GetLocation())
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
        err := writer.WriteObjectValue("softwareInfo", m.GetSoftwareInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageSystem", m.GetStorageSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("storageUsage", m.GetStorageUsage())
        if err != nil {
            return err
        }
    }
    if m.GetSystemVms() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemVms()))
        for i, v := range m.GetSystemVms() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("systemVms", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
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
func (m *V1beta1SystemsItemGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssociatedResourceCounts sets the associatedResourceCounts property value. Associated Resource Information of system.
func (m *V1beta1SystemsItemGetResponse) SetAssociatedResourceCounts(value V1beta1SystemsItemGetResponse_associatedResourceCountsable)() {
    m.associatedResourceCounts = value
}
// SetComputeClusters sets the computeClusters property value. List of compute clusters in the system.
func (m *V1beta1SystemsItemGetResponse) SetComputeClusters(value []V1beta1SystemsItemGetResponse_computeClustersable)() {
    m.computeClusters = value
}
// SetComputeUsage sets the computeUsage property value. System Compute Usage Information.
func (m *V1beta1SystemsItemGetResponse) SetComputeUsage(value V1beta1SystemsItemGetResponse_computeUsageable)() {
    m.computeUsage = value
}
// SetConfigAnalysisStatus sets the configAnalysisStatus property value. Status of last run of configuration analysis job.
func (m *V1beta1SystemsItemGetResponse) SetConfigAnalysisStatus(value V1beta1SystemsItemGetResponse_configAnalysisStatusable)() {
    m.configAnalysisStatus = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SystemsItemGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SystemsItemGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SystemsItemGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHealth sets the health property value. Aspects of system health.Deduced health of storage subsystem based on associated arrays, controllers, disks,shelves, power supply, network interfaces, fan and temperature sensors.Aggregated health of servers based on health statuses of multiple servers in the system.Aggregated health of network based on health statues of multiple switches used in the system.Deduced/Aggregated overall health of the system based on storage, servers and networking health.
func (m *V1beta1SystemsItemGetResponse) SetHealth(value V1beta1SystemsItemGetResponse_healthable)() {
    m.health = value
}
// SetHypervisorClusters sets the hypervisorClusters property value. List of hypervisor clusters in the system. To be deprecated soon.
func (m *V1beta1SystemsItemGetResponse) SetHypervisorClusters(value []V1beta1SystemsItemGetResponse_hypervisorClustersable)() {
    m.hypervisorClusters = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SystemsItemGetResponse) SetId(value *string)() {
    m.id = value
}
// SetLocation sets the location property value. System Location Information.
func (m *V1beta1SystemsItemGetResponse) SetLocation(value V1beta1SystemsItemGetResponse_locationable)() {
    m.location = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SystemsItemGetResponse) SetName(value *string)() {
    m.name = value
}
// SetOnPremAgentId sets the onPremAgentId property value. An identifier of the on prem agent for the system.
func (m *V1beta1SystemsItemGetResponse) SetOnPremAgentId(value *string)() {
    m.onPremAgentId = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SystemsItemGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSoftwareInfo sets the softwareInfo property value. system software information.
func (m *V1beta1SystemsItemGetResponse) SetSoftwareInfo(value V1beta1SystemsItemGetResponse_softwareInfoable)() {
    m.softwareInfo = value
}
// SetStorageSystem sets the storageSystem property value. Storage Information of system.
func (m *V1beta1SystemsItemGetResponse) SetStorageSystem(value V1beta1SystemsItemGetResponse_storageSystemable)() {
    m.storageSystem = value
}
// SetStorageUsage sets the storageUsage property value. Storage Usage Information of system.
func (m *V1beta1SystemsItemGetResponse) SetStorageUsage(value V1beta1SystemsItemGetResponse_storageUsageable)() {
    m.storageUsage = value
}
// SetSystemVms sets the systemVms property value. List of system virtual machine information.
func (m *V1beta1SystemsItemGetResponse) SetSystemVms(value []V1beta1SystemsItemGetResponse_systemVmsable)() {
    m.systemVms = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SystemsItemGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SystemsItemGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SystemsItemGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedResourceCounts()(V1beta1SystemsItemGetResponse_associatedResourceCountsable)
    GetComputeClusters()([]V1beta1SystemsItemGetResponse_computeClustersable)
    GetComputeUsage()(V1beta1SystemsItemGetResponse_computeUsageable)
    GetConfigAnalysisStatus()(V1beta1SystemsItemGetResponse_configAnalysisStatusable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetHealth()(V1beta1SystemsItemGetResponse_healthable)
    GetHypervisorClusters()([]V1beta1SystemsItemGetResponse_hypervisorClustersable)
    GetId()(*string)
    GetLocation()(V1beta1SystemsItemGetResponse_locationable)
    GetName()(*string)
    GetOnPremAgentId()(*string)
    GetResourceUri()(*string)
    GetSoftwareInfo()(V1beta1SystemsItemGetResponse_softwareInfoable)
    GetStorageSystem()(V1beta1SystemsItemGetResponse_storageSystemable)
    GetStorageUsage()(V1beta1SystemsItemGetResponse_storageUsageable)
    GetSystemVms()([]V1beta1SystemsItemGetResponse_systemVmsable)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssociatedResourceCounts(value V1beta1SystemsItemGetResponse_associatedResourceCountsable)()
    SetComputeClusters(value []V1beta1SystemsItemGetResponse_computeClustersable)()
    SetComputeUsage(value V1beta1SystemsItemGetResponse_computeUsageable)()
    SetConfigAnalysisStatus(value V1beta1SystemsItemGetResponse_configAnalysisStatusable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetHealth(value V1beta1SystemsItemGetResponse_healthable)()
    SetHypervisorClusters(value []V1beta1SystemsItemGetResponse_hypervisorClustersable)()
    SetId(value *string)()
    SetLocation(value V1beta1SystemsItemGetResponse_locationable)()
    SetName(value *string)()
    SetOnPremAgentId(value *string)()
    SetResourceUri(value *string)()
    SetSoftwareInfo(value V1beta1SystemsItemGetResponse_softwareInfoable)()
    SetStorageSystem(value V1beta1SystemsItemGetResponse_storageSystemable)()
    SetStorageUsage(value V1beta1SystemsItemGetResponse_storageUsageable)()
    SetSystemVms(value []V1beta1SystemsItemGetResponse_systemVmsable)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
