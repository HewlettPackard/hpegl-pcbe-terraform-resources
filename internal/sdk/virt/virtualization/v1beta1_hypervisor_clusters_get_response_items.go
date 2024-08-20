package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorClustersGetResponse_items represents a single instance of a hypervisor cluster (Vmware ESXi Cluster)
type V1beta1HypervisorClustersGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Application specific information for this cluster.
    appInfo V1beta1HypervisorClustersGetResponse_items_appInfoable
    // Hypervisor cluster performance metrics.
    clusterPerfMetricInfo V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // A user-friendly name that identifies the hypervisor cluster. This will always be same as name since adding or updating hypervisor clusters is not supported when managed from a manager, such as the vCenter.
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // Captures the list of all hosts of this cluster. In a VMWare cluster, this entity maps to ESXi hosts.
    hypervisorHosts []V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable
    // UUID string uniquely identifying the hypervisor cluster.
    id *string
    // Name of the cluster as reported by the hypervisor manager.
    name *string
    // All the network names associated with this cluster.
    networksInfo []string
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the hypervisor cluster.
    stateReason *string
    // The type of resource.
    typeEscaped *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1HypervisorClustersGetResponse_items instantiates a new V1beta1HypervisorClustersGetResponse_items and sets the default values.
func NewV1beta1HypervisorClustersGetResponse_items()(*V1beta1HypervisorClustersGetResponse_items) {
    m := &V1beta1HypervisorClustersGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorClustersGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorClustersGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorClustersGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Application specific information for this cluster.
// returns a V1beta1HypervisorClustersGetResponse_items_appInfoable when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetAppInfo()(V1beta1HypervisorClustersGetResponse_items_appInfoable) {
    return m.appInfo
}
// GetClusterPerfMetricInfo gets the clusterPerfMetricInfo property value. Hypervisor cluster performance metrics.
// returns a V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetClusterPerfMetricInfo()(V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable) {
    return m.clusterPerfMetricInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the hypervisor cluster. This will always be same as name since adding or updating hypervisor clusters is not supported when managed from a manager, such as the vCenter.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorClustersGetResponse_items_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1HypervisorClustersGetResponse_items_appInfoable))
        }
        return nil
    }
    res["clusterPerfMetricInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterPerfMetricInfo(val.(V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable))
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["hciClusterUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHciClusterUuid(val)
        }
        return nil
    }
    res["hypervisorHosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorClustersGetResponse_items_hypervisorHostsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable)
                }
            }
            m.SetHypervisorHosts(res)
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable))
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
    res["networksInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNetworksInfo(res)
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
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["stateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val)
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
func (m *V1beta1HypervisorClustersGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHypervisorHosts gets the hypervisorHosts property value. Captures the list of all hosts of this cluster. In a VMWare cluster, this entity maps to ESXi hosts.
// returns a []V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetHypervisorHosts()([]V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable) {
    return m.hypervisorHosts
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetHypervisorManagerInfo()(V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor cluster.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the cluster as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetName()(*string) {
    return m.name
}
// GetNetworksInfo gets the networksInfo property value. All the network names associated with this cluster.
// returns a []string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetNetworksInfo()([]string) {
    return m.networksInfo
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the hypervisor cluster.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetStateReason()(*string) {
    return m.stateReason
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1HypervisorClustersGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorClustersGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clusterPerfMetricInfo", m.GetClusterPerfMetricInfo())
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
        err := writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hciClusterUuid", m.GetHciClusterUuid())
        if err != nil {
            return err
        }
    }
    if m.GetHypervisorHosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHypervisorHosts()))
        for i, v := range m.GetHypervisorHosts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hypervisorHosts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorManagerInfo", m.GetHypervisorManagerInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetNetworksInfo() != nil {
        err := writer.WriteCollectionOfStringValues("networksInfo", m.GetNetworksInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        err := writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateReason", m.GetStateReason())
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
func (m *V1beta1HypervisorClustersGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Application specific information for this cluster.
func (m *V1beta1HypervisorClustersGetResponse_items) SetAppInfo(value V1beta1HypervisorClustersGetResponse_items_appInfoable)() {
    m.appInfo = value
}
// SetClusterPerfMetricInfo sets the clusterPerfMetricInfo property value. Hypervisor cluster performance metrics.
func (m *V1beta1HypervisorClustersGetResponse_items) SetClusterPerfMetricInfo(value V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable)() {
    m.clusterPerfMetricInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1HypervisorClustersGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorClustersGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the hypervisor cluster. This will always be same as name since adding or updating hypervisor clusters is not supported when managed from a manager, such as the vCenter.
func (m *V1beta1HypervisorClustersGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorClustersGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1HypervisorClustersGetResponse_items) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHypervisorHosts sets the hypervisorHosts property value. Captures the list of all hosts of this cluster. In a VMWare cluster, this entity maps to ESXi hosts.
func (m *V1beta1HypervisorClustersGetResponse_items) SetHypervisorHosts(value []V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable)() {
    m.hypervisorHosts = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1HypervisorClustersGetResponse_items) SetHypervisorManagerInfo(value V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor cluster.
func (m *V1beta1HypervisorClustersGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the cluster as reported by the hypervisor manager.
func (m *V1beta1HypervisorClustersGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetNetworksInfo sets the networksInfo property value. All the network names associated with this cluster.
func (m *V1beta1HypervisorClustersGetResponse_items) SetNetworksInfo(value []string)() {
    m.networksInfo = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorClustersGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1HypervisorClustersGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the hypervisor cluster.
func (m *V1beta1HypervisorClustersGetResponse_items) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorClustersGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1HypervisorClustersGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1HypervisorClustersGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1HypervisorClustersGetResponse_items_appInfoable)
    GetClusterPerfMetricInfo()(V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetHciClusterUuid()(*string)
    GetHypervisorHosts()([]V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable)
    GetHypervisorManagerInfo()(V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable)
    GetId()(*string)
    GetName()(*string)
    GetNetworksInfo()([]string)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAppInfo(value V1beta1HypervisorClustersGetResponse_items_appInfoable)()
    SetClusterPerfMetricInfo(value V1beta1HypervisorClustersGetResponse_items_clusterPerfMetricInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetHciClusterUuid(value *string)()
    SetHypervisorHosts(value []V1beta1HypervisorClustersGetResponse_items_hypervisorHostsable)()
    SetHypervisorManagerInfo(value V1beta1HypervisorClustersGetResponse_items_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetName(value *string)()
    SetNetworksInfo(value []string)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
