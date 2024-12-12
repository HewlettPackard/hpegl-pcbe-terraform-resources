package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsGetResponse_items represents a single instance of a hypervisor host (Vmware ESXi)
type V1beta1HypervisorHostsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Application specific information for this host.
    appInfo V1beta1HypervisorHostsGetResponse_items_appInfoable
    // CPU information.
    cpuInfo V1beta1HypervisorHostsGetResponse_items_cpuInfoable
    // CPU sockets information.
    cpuSockets []V1beta1HypervisorHostsGetResponse_items_cpuSocketsable
    // Time in UTC at which the object was created
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // A user-friendly name that identifies the hypervisor host. This will always be same as name since adding or updating hypervisor hosts is not supported when managed from a manager, such as the vCenter.
    displayName *string
    // fqdn value of hypervisor host.
    fqdn *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // UUID string uniquely identifying the HCI server.
    hciServerUuid *string
    // Host network system information.
    hostNetworkSystem V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable
    // Hypervisor host performance metrics.
    hostPerfMetricInfo V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable
    // UUID string uniquely identifying the hypervisor host.
    id *string
    // maintenance mode status of hypervisor host.
    inMaintenanceMode *bool
    // Name of the host as reported by the hypervisor manager.
    name *string
    // An IP address or hostname or FQDN to address the hypervisor host.
    networkAddress *string
    // All the network names associated with this host.
    networksInfo []string
    // Parent of this host. It could be a cluster or folder in case of a VMware ESXi Host. For a Hyper-V host it will be cluster or host group.
    parentInfo V1beta1HypervisorHostsGetResponse_items_parentInfoable
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the hypervisor host
    stateReason *string
    // List of storage adapters associated with this host.
    storageAdaptersInfo []V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable
    // The type of resource.
    typeEscaped *string
    // A hypervisor host provided durable UID.
    uid *string
    // Time in UTC at which the object was last updated
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The hypervisor host version.
    version *string
}
// NewV1beta1HypervisorHostsGetResponse_items instantiates a new V1beta1HypervisorHostsGetResponse_items and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items()(*V1beta1HypervisorHostsGetResponse_items) {
    m := &V1beta1HypervisorHostsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Application specific information for this host.
// returns a V1beta1HypervisorHostsGetResponse_items_appInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetAppInfo()(V1beta1HypervisorHostsGetResponse_items_appInfoable) {
    return m.appInfo
}
// GetCpuInfo gets the cpuInfo property value. CPU information.
// returns a V1beta1HypervisorHostsGetResponse_items_cpuInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetCpuInfo()(V1beta1HypervisorHostsGetResponse_items_cpuInfoable) {
    return m.cpuInfo
}
// GetCpuSockets gets the cpuSockets property value. CPU sockets information.
// returns a []V1beta1HypervisorHostsGetResponse_items_cpuSocketsable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetCpuSockets()([]V1beta1HypervisorHostsGetResponse_items_cpuSocketsable) {
    return m.cpuSockets
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created
// returns a *Time when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the hypervisor host. This will always be same as name since adding or updating hypervisor hosts is not supported when managed from a manager, such as the vCenter.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1HypervisorHostsGetResponse_items_appInfoable))
        }
        return nil
    }
    res["cpuInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_cpuInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuInfo(val.(V1beta1HypervisorHostsGetResponse_items_cpuInfoable))
        }
        return nil
    }
    res["cpuSockets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorHostsGetResponse_items_cpuSocketsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorHostsGetResponse_items_cpuSocketsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorHostsGetResponse_items_cpuSocketsable)
                }
            }
            m.SetCpuSockets(res)
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
    res["fqdn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
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
    res["hciServerUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHciServerUuid(val)
        }
        return nil
    }
    res["hostNetworkSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostNetworkSystem(val.(V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable))
        }
        return nil
    }
    res["hostPerfMetricInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostPerfMetricInfo(val.(V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable))
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable))
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
    res["inMaintenanceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInMaintenanceMode(val)
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
    res["networkAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkAddress(val)
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
    res["parentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsGetResponse_items_parentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentInfo(val.(V1beta1HypervisorHostsGetResponse_items_parentInfoable))
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
    res["storageAdaptersInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable)
                }
            }
            m.SetStorageAdaptersInfo(res)
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
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetFqdn gets the fqdn property value. fqdn value of hypervisor host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetFqdn()(*string) {
    return m.fqdn
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHciServerUuid gets the hciServerUuid property value. UUID string uniquely identifying the HCI server.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetHciServerUuid()(*string) {
    return m.hciServerUuid
}
// GetHostNetworkSystem gets the hostNetworkSystem property value. Host network system information.
// returns a V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetHostNetworkSystem()(V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable) {
    return m.hostNetworkSystem
}
// GetHostPerfMetricInfo gets the hostPerfMetricInfo property value. Hypervisor host performance metrics.
// returns a V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetHostPerfMetricInfo()(V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable) {
    return m.hostPerfMetricInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetHypervisorManagerInfo()(V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetId()(*string) {
    return m.id
}
// GetInMaintenanceMode gets the inMaintenanceMode property value. maintenance mode status of hypervisor host.
// returns a *bool when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetInMaintenanceMode()(*bool) {
    return m.inMaintenanceMode
}
// GetName gets the name property value. Name of the host as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetName()(*string) {
    return m.name
}
// GetNetworkAddress gets the networkAddress property value. An IP address or hostname or FQDN to address the hypervisor host.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetNetworksInfo gets the networksInfo property value. All the network names associated with this host.
// returns a []string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetNetworksInfo()([]string) {
    return m.networksInfo
}
// GetParentInfo gets the parentInfo property value. Parent of this host. It could be a cluster or folder in case of a VMware ESXi Host. For a Hyper-V host it will be cluster or host group.
// returns a V1beta1HypervisorHostsGetResponse_items_parentInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetParentInfo()(V1beta1HypervisorHostsGetResponse_items_parentInfoable) {
    return m.parentInfo
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the hypervisor host
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageAdaptersInfo gets the storageAdaptersInfo property value. List of storage adapters associated with this host.
// returns a []V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetStorageAdaptersInfo()([]V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable) {
    return m.storageAdaptersInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. A hypervisor host provided durable UID.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated
// returns a *Time when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVersion gets the version property value. The hypervisor host version.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cpuInfo", m.GetCpuInfo())
        if err != nil {
            return err
        }
    }
    if m.GetCpuSockets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCpuSockets()))
        for i, v := range m.GetCpuSockets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("cpuSockets", cast)
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
        err := writer.WriteStringValue("fqdn", m.GetFqdn())
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
    {
        err := writer.WriteStringValue("hciServerUuid", m.GetHciServerUuid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hostNetworkSystem", m.GetHostNetworkSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hostPerfMetricInfo", m.GetHostPerfMetricInfo())
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
    {
        err := writer.WriteBoolValue("inMaintenanceMode", m.GetInMaintenanceMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkAddress", m.GetNetworkAddress())
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
        err := writer.WriteObjectValue("parentInfo", m.GetParentInfo())
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
    if m.GetStorageAdaptersInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStorageAdaptersInfo()))
        for i, v := range m.GetStorageAdaptersInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("storageAdaptersInfo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uid", m.GetUid())
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
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1HypervisorHostsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Application specific information for this host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetAppInfo(value V1beta1HypervisorHostsGetResponse_items_appInfoable)() {
    m.appInfo = value
}
// SetCpuInfo sets the cpuInfo property value. CPU information.
func (m *V1beta1HypervisorHostsGetResponse_items) SetCpuInfo(value V1beta1HypervisorHostsGetResponse_items_cpuInfoable)() {
    m.cpuInfo = value
}
// SetCpuSockets sets the cpuSockets property value. CPU sockets information.
func (m *V1beta1HypervisorHostsGetResponse_items) SetCpuSockets(value []V1beta1HypervisorHostsGetResponse_items_cpuSocketsable)() {
    m.cpuSockets = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created
func (m *V1beta1HypervisorHostsGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorHostsGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the hypervisor host. This will always be same as name since adding or updating hypervisor hosts is not supported when managed from a manager, such as the vCenter.
func (m *V1beta1HypervisorHostsGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFqdn sets the fqdn property value. fqdn value of hypervisor host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetFqdn(value *string)() {
    m.fqdn = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorHostsGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1HypervisorHostsGetResponse_items) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHciServerUuid sets the hciServerUuid property value. UUID string uniquely identifying the HCI server.
func (m *V1beta1HypervisorHostsGetResponse_items) SetHciServerUuid(value *string)() {
    m.hciServerUuid = value
}
// SetHostNetworkSystem sets the hostNetworkSystem property value. Host network system information.
func (m *V1beta1HypervisorHostsGetResponse_items) SetHostNetworkSystem(value V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable)() {
    m.hostNetworkSystem = value
}
// SetHostPerfMetricInfo sets the hostPerfMetricInfo property value. Hypervisor host performance metrics.
func (m *V1beta1HypervisorHostsGetResponse_items) SetHostPerfMetricInfo(value V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable)() {
    m.hostPerfMetricInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1HypervisorHostsGetResponse_items) SetHypervisorManagerInfo(value V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetInMaintenanceMode sets the inMaintenanceMode property value. maintenance mode status of hypervisor host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetInMaintenanceMode(value *bool)() {
    m.inMaintenanceMode = value
}
// SetName sets the name property value. Name of the host as reported by the hypervisor manager.
func (m *V1beta1HypervisorHostsGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetNetworkAddress sets the networkAddress property value. An IP address or hostname or FQDN to address the hypervisor host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetNetworksInfo sets the networksInfo property value. All the network names associated with this host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetNetworksInfo(value []string)() {
    m.networksInfo = value
}
// SetParentInfo sets the parentInfo property value. Parent of this host. It could be a cluster or folder in case of a VMware ESXi Host. For a Hyper-V host it will be cluster or host group.
func (m *V1beta1HypervisorHostsGetResponse_items) SetParentInfo(value V1beta1HypervisorHostsGetResponse_items_parentInfoable)() {
    m.parentInfo = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorHostsGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1HypervisorHostsGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the hypervisor host
func (m *V1beta1HypervisorHostsGetResponse_items) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageAdaptersInfo sets the storageAdaptersInfo property value. List of storage adapters associated with this host.
func (m *V1beta1HypervisorHostsGetResponse_items) SetStorageAdaptersInfo(value []V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable)() {
    m.storageAdaptersInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorHostsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. A hypervisor host provided durable UID.
func (m *V1beta1HypervisorHostsGetResponse_items) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated
func (m *V1beta1HypervisorHostsGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVersion sets the version property value. The hypervisor host version.
func (m *V1beta1HypervisorHostsGetResponse_items) SetVersion(value *string)() {
    m.version = value
}
type V1beta1HypervisorHostsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1HypervisorHostsGetResponse_items_appInfoable)
    GetCpuInfo()(V1beta1HypervisorHostsGetResponse_items_cpuInfoable)
    GetCpuSockets()([]V1beta1HypervisorHostsGetResponse_items_cpuSocketsable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetFqdn()(*string)
    GetGeneration()(*int64)
    GetHciClusterUuid()(*string)
    GetHciServerUuid()(*string)
    GetHostNetworkSystem()(V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable)
    GetHostPerfMetricInfo()(V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable)
    GetHypervisorManagerInfo()(V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable)
    GetId()(*string)
    GetInMaintenanceMode()(*bool)
    GetName()(*string)
    GetNetworkAddress()(*string)
    GetNetworksInfo()([]string)
    GetParentInfo()(V1beta1HypervisorHostsGetResponse_items_parentInfoable)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetStorageAdaptersInfo()([]V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetAppInfo(value V1beta1HypervisorHostsGetResponse_items_appInfoable)()
    SetCpuInfo(value V1beta1HypervisorHostsGetResponse_items_cpuInfoable)()
    SetCpuSockets(value []V1beta1HypervisorHostsGetResponse_items_cpuSocketsable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetFqdn(value *string)()
    SetGeneration(value *int64)()
    SetHciClusterUuid(value *string)()
    SetHciServerUuid(value *string)()
    SetHostNetworkSystem(value V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable)()
    SetHostPerfMetricInfo(value V1beta1HypervisorHostsGetResponse_items_hostPerfMetricInfoable)()
    SetHypervisorManagerInfo(value V1beta1HypervisorHostsGetResponse_items_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetInMaintenanceMode(value *bool)()
    SetName(value *string)()
    SetNetworkAddress(value *string)()
    SetNetworksInfo(value []string)()
    SetParentInfo(value V1beta1HypervisorHostsGetResponse_items_parentInfoable)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetStorageAdaptersInfo(value []V1beta1HypervisorHostsGetResponse_items_storageAdaptersInfoable)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
