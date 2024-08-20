package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items represents a single instance of a Virtual Machine
type V1beta1VirtualMachinesGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Application specific information for this virtual machine.
    appInfo V1beta1VirtualMachinesGetResponse_items_appInfoable
    // Size of the virtual machine in bytes.
    capacityInBytes *int64
    // The clusterInfo property
    clusterInfo V1beta1VirtualMachinesGetResponse_items_clusterInfoable
    // Compute information of the virtual machine.
    computeInfo V1beta1VirtualMachinesGetResponse_items_computeInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // A user-friendly name that identifies the virtual machine.
    displayName *string
    // The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
    folderInfo V1beta1VirtualMachinesGetResponse_items_folderInfoable
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Graphics configuration for the virtual machine.
    graphics V1beta1VirtualMachinesGetResponse_items_graphicsable
    // Information of this guest OS running on the virtual machine.
    guestInfo V1beta1VirtualMachinesGetResponse_items_guestInfoable
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // High availability settings of the virtual machine.
    highAvailability V1beta1VirtualMachinesGetResponse_items_highAvailabilityable
    // The hostInfo property
    hostInfo V1beta1VirtualMachinesGetResponse_items_hostInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable
    // UUID string uniquely identifying the virtual machine
    id *string
    // Maximum CPU Usage value of the virtual machine.
    maxCpuUsage *int64
    // Name of the virtual machine as configured in the hypervisor manager.
    name *string
    // The networkAdapters property
    networkAdapters []V1beta1VirtualMachinesGetResponse_items_networkAdaptersable
    // IP address of the virtual machine.
    networkAddress *string
    // Network Addresses associated with virtual machine.
    networkAddresses []string
    // Information about the assigned Protection Policy and the Protection Job.
    protectionJobInfo V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable
    // Describes applied protection policy information.
    protectionPolicyAppliedAtInfo V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable
    // Indicates at least one valid recovery point exists for this resource.
    recoveryPointsExist *bool
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the virtual machine
    stateReason *string
    // storage systems with volume wwn associated with this virtual machine
    storageSystemsInfo []V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable
    // The type of resource.
    typeEscaped *string
    // Unique identifier of the virtual machine as reported by the hypervisor.
    uid *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates it is a vCLS virtual machine or not.
    vclsVm *bool
    // Provides Ahci Disk Controller details of the virtual machine.
    virtualAhciControllers V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable
    // A list of objects encapsulating information about the storage disks provisioned to a virtual machine
    virtualDisks []V1beta1VirtualMachinesGetResponse_items_virtualDisksable
    // Provides Ide Disk Controller details of the virtual machine.
    virtualIdeControllers V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable
    // Provides Nvme Disk Controller details of the virtual machine.
    virtualNvmeControllers V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable
    // Provides SCSI Disk Controller details of the virtual machine.
    virtualScsiControllers V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable
    // VM configuration path of the virtual machine.
    vmConfigPath *string
    // Virtual machine performance metrics.
    vmPerfMetricInfo V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable
    // Protection groups related to the virtual machine.
    vmProtectionGroupsInfo []V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable
    // Volumes associated with this virtual machine.
    volumesInfo []V1beta1VirtualMachinesGetResponse_items_volumesInfoable
}
// NewV1beta1VirtualMachinesGetResponse_items instantiates a new V1beta1VirtualMachinesGetResponse_items and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items()(*V1beta1VirtualMachinesGetResponse_items) {
    m := &V1beta1VirtualMachinesGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Application specific information for this virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_appInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetAppInfo()(V1beta1VirtualMachinesGetResponse_items_appInfoable) {
    return m.appInfo
}
// GetCapacityInBytes gets the capacityInBytes property value. Size of the virtual machine in bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetCapacityInBytes()(*int64) {
    return m.capacityInBytes
}
// GetClusterInfo gets the clusterInfo property value. The clusterInfo property
// returns a V1beta1VirtualMachinesGetResponse_items_clusterInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetClusterInfo()(V1beta1VirtualMachinesGetResponse_items_clusterInfoable) {
    return m.clusterInfo
}
// GetComputeInfo gets the computeInfo property value. Compute information of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_computeInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetComputeInfo()(V1beta1VirtualMachinesGetResponse_items_computeInfoable) {
    return m.computeInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1VirtualMachinesGetResponse_items_appInfoable))
        }
        return nil
    }
    res["capacityInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityInBytes(val)
        }
        return nil
    }
    res["clusterInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_clusterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterInfo(val.(V1beta1VirtualMachinesGetResponse_items_clusterInfoable))
        }
        return nil
    }
    res["computeInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_computeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeInfo(val.(V1beta1VirtualMachinesGetResponse_items_computeInfoable))
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
    res["folderInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_folderInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderInfo(val.(V1beta1VirtualMachinesGetResponse_items_folderInfoable))
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
    res["graphics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_graphicsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGraphics(val.(V1beta1VirtualMachinesGetResponse_items_graphicsable))
        }
        return nil
    }
    res["guestInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_guestInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestInfo(val.(V1beta1VirtualMachinesGetResponse_items_guestInfoable))
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
    res["highAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_highAvailabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighAvailability(val.(V1beta1VirtualMachinesGetResponse_items_highAvailabilityable))
        }
        return nil
    }
    res["hostInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_hostInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostInfo(val.(V1beta1VirtualMachinesGetResponse_items_hostInfoable))
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable))
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
    res["maxCpuUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCpuUsage(val)
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
    res["networkAdapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_networkAdaptersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_networkAdaptersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_networkAdaptersable)
                }
            }
            m.SetNetworkAdapters(res)
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
    res["networkAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNetworkAddresses(res)
        }
        return nil
    }
    res["protectionJobInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_protectionJobInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionJobInfo(val.(V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable))
        }
        return nil
    }
    res["protectionPolicyAppliedAtInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyAppliedAtInfo(val.(V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable))
        }
        return nil
    }
    res["recoveryPointsExist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryPointsExist(val)
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
    res["storageSystemsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_storageSystemsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable)
                }
            }
            m.SetStorageSystemsInfo(res)
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
    res["vclsVm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVclsVm(val)
        }
        return nil
    }
    res["virtualAhciControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualAhciControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualAhciControllers(val.(V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable))
        }
        return nil
    }
    res["virtualDisks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_virtualDisksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_virtualDisksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_virtualDisksable)
                }
            }
            m.SetVirtualDisks(res)
        }
        return nil
    }
    res["virtualIdeControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualIdeControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualIdeControllers(val.(V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable))
        }
        return nil
    }
    res["virtualNvmeControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNvmeControllers(val.(V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable))
        }
        return nil
    }
    res["virtualScsiControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_virtualScsiControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualScsiControllers(val.(V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable))
        }
        return nil
    }
    res["vmConfigPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmConfigPath(val)
        }
        return nil
    }
    res["vmPerfMetricInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmPerfMetricInfo(val.(V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable))
        }
        return nil
    }
    res["vmProtectionGroupsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable)
                }
            }
            m.SetVmProtectionGroupsInfo(res)
        }
        return nil
    }
    res["volumesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesGetResponse_items_volumesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesGetResponse_items_volumesInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesGetResponse_items_volumesInfoable)
                }
            }
            m.SetVolumesInfo(res)
        }
        return nil
    }
    return res
}
// GetFolderInfo gets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
// returns a V1beta1VirtualMachinesGetResponse_items_folderInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetFolderInfo()(V1beta1VirtualMachinesGetResponse_items_folderInfoable) {
    return m.folderInfo
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetGraphics gets the graphics property value. Graphics configuration for the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_graphicsable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetGraphics()(V1beta1VirtualMachinesGetResponse_items_graphicsable) {
    return m.graphics
}
// GetGuestInfo gets the guestInfo property value. Information of this guest OS running on the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_guestInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetGuestInfo()(V1beta1VirtualMachinesGetResponse_items_guestInfoable) {
    return m.guestInfo
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHighAvailability gets the highAvailability property value. High availability settings of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_highAvailabilityable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetHighAvailability()(V1beta1VirtualMachinesGetResponse_items_highAvailabilityable) {
    return m.highAvailability
}
// GetHostInfo gets the hostInfo property value. The hostInfo property
// returns a V1beta1VirtualMachinesGetResponse_items_hostInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetHostInfo()(V1beta1VirtualMachinesGetResponse_items_hostInfoable) {
    return m.hostInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetHypervisorManagerInfo()(V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetId()(*string) {
    return m.id
}
// GetMaxCpuUsage gets the maxCpuUsage property value. Maximum CPU Usage value of the virtual machine.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetMaxCpuUsage()(*int64) {
    return m.maxCpuUsage
}
// GetName gets the name property value. Name of the virtual machine as configured in the hypervisor manager.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetName()(*string) {
    return m.name
}
// GetNetworkAdapters gets the networkAdapters property value. The networkAdapters property
// returns a []V1beta1VirtualMachinesGetResponse_items_networkAdaptersable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetNetworkAdapters()([]V1beta1VirtualMachinesGetResponse_items_networkAdaptersable) {
    return m.networkAdapters
}
// GetNetworkAddress gets the networkAddress property value. IP address of the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetNetworkAddresses gets the networkAddresses property value. Network Addresses associated with virtual machine.
// returns a []string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetNetworkAddresses()([]string) {
    return m.networkAddresses
}
// GetProtectionJobInfo gets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
// returns a V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetProtectionJobInfo()(V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable) {
    return m.protectionJobInfo
}
// GetProtectionPolicyAppliedAtInfo gets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
// returns a V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetProtectionPolicyAppliedAtInfo()(V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable) {
    return m.protectionPolicyAppliedAtInfo
}
// GetRecoveryPointsExist gets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetRecoveryPointsExist()(*bool) {
    return m.recoveryPointsExist
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageSystemsInfo gets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
// returns a []V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetStorageSystemsInfo()([]V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable) {
    return m.storageSystemsInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. Unique identifier of the virtual machine as reported by the hypervisor.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVclsVm gets the vclsVm property value. Indicates it is a vCLS virtual machine or not.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVclsVm()(*bool) {
    return m.vclsVm
}
// GetVirtualAhciControllers gets the virtualAhciControllers property value. Provides Ahci Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVirtualAhciControllers()(V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable) {
    return m.virtualAhciControllers
}
// GetVirtualDisks gets the virtualDisks property value. A list of objects encapsulating information about the storage disks provisioned to a virtual machine
// returns a []V1beta1VirtualMachinesGetResponse_items_virtualDisksable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVirtualDisks()([]V1beta1VirtualMachinesGetResponse_items_virtualDisksable) {
    return m.virtualDisks
}
// GetVirtualIdeControllers gets the virtualIdeControllers property value. Provides Ide Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVirtualIdeControllers()(V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable) {
    return m.virtualIdeControllers
}
// GetVirtualNvmeControllers gets the virtualNvmeControllers property value. Provides Nvme Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVirtualNvmeControllers()(V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable) {
    return m.virtualNvmeControllers
}
// GetVirtualScsiControllers gets the virtualScsiControllers property value. Provides SCSI Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVirtualScsiControllers()(V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable) {
    return m.virtualScsiControllers
}
// GetVmConfigPath gets the vmConfigPath property value. VM configuration path of the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVmConfigPath()(*string) {
    return m.vmConfigPath
}
// GetVmPerfMetricInfo gets the vmPerfMetricInfo property value. Virtual machine performance metrics.
// returns a V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVmPerfMetricInfo()(V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable) {
    return m.vmPerfMetricInfo
}
// GetVmProtectionGroupsInfo gets the vmProtectionGroupsInfo property value. Protection groups related to the virtual machine.
// returns a []V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVmProtectionGroupsInfo()([]V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable) {
    return m.vmProtectionGroupsInfo
}
// GetVolumesInfo gets the volumesInfo property value. Volumes associated with this virtual machine.
// returns a []V1beta1VirtualMachinesGetResponse_items_volumesInfoable when successful
func (m *V1beta1VirtualMachinesGetResponse_items) GetVolumesInfo()([]V1beta1VirtualMachinesGetResponse_items_volumesInfoable) {
    return m.volumesInfo
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("capacityInBytes", m.GetCapacityInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clusterInfo", m.GetClusterInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("computeInfo", m.GetComputeInfo())
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
        err := writer.WriteObjectValue("folderInfo", m.GetFolderInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("graphics", m.GetGraphics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("guestInfo", m.GetGuestInfo())
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
        err := writer.WriteObjectValue("highAvailability", m.GetHighAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hostInfo", m.GetHostInfo())
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
        err := writer.WriteInt64Value("maxCpuUsage", m.GetMaxCpuUsage())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkAdapters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkAdapters()))
        for i, v := range m.GetNetworkAdapters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkAdapters", cast)
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
    if m.GetNetworkAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("networkAddresses", m.GetNetworkAddresses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("protectionJobInfo", m.GetProtectionJobInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("protectionPolicyAppliedAtInfo", m.GetProtectionPolicyAppliedAtInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recoveryPointsExist", m.GetRecoveryPointsExist())
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
    if m.GetStorageSystemsInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStorageSystemsInfo()))
        for i, v := range m.GetStorageSystemsInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("storageSystemsInfo", cast)
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
        err := writer.WriteBoolValue("vclsVm", m.GetVclsVm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("virtualAhciControllers", m.GetVirtualAhciControllers())
        if err != nil {
            return err
        }
    }
    if m.GetVirtualDisks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVirtualDisks()))
        for i, v := range m.GetVirtualDisks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("virtualDisks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("virtualIdeControllers", m.GetVirtualIdeControllers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("virtualNvmeControllers", m.GetVirtualNvmeControllers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("virtualScsiControllers", m.GetVirtualScsiControllers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vmConfigPath", m.GetVmConfigPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vmPerfMetricInfo", m.GetVmPerfMetricInfo())
        if err != nil {
            return err
        }
    }
    if m.GetVmProtectionGroupsInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVmProtectionGroupsInfo()))
        for i, v := range m.GetVmProtectionGroupsInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("vmProtectionGroupsInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVolumesInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVolumesInfo()))
        for i, v := range m.GetVolumesInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("volumesInfo", cast)
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
func (m *V1beta1VirtualMachinesGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Application specific information for this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetAppInfo(value V1beta1VirtualMachinesGetResponse_items_appInfoable)() {
    m.appInfo = value
}
// SetCapacityInBytes sets the capacityInBytes property value. Size of the virtual machine in bytes.
func (m *V1beta1VirtualMachinesGetResponse_items) SetCapacityInBytes(value *int64)() {
    m.capacityInBytes = value
}
// SetClusterInfo sets the clusterInfo property value. The clusterInfo property
func (m *V1beta1VirtualMachinesGetResponse_items) SetClusterInfo(value V1beta1VirtualMachinesGetResponse_items_clusterInfoable)() {
    m.clusterInfo = value
}
// SetComputeInfo sets the computeInfo property value. Compute information of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetComputeInfo(value V1beta1VirtualMachinesGetResponse_items_computeInfoable)() {
    m.computeInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1VirtualMachinesGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1VirtualMachinesGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFolderInfo sets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
func (m *V1beta1VirtualMachinesGetResponse_items) SetFolderInfo(value V1beta1VirtualMachinesGetResponse_items_folderInfoable)() {
    m.folderInfo = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1VirtualMachinesGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGraphics sets the graphics property value. Graphics configuration for the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetGraphics(value V1beta1VirtualMachinesGetResponse_items_graphicsable)() {
    m.graphics = value
}
// SetGuestInfo sets the guestInfo property value. Information of this guest OS running on the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetGuestInfo(value V1beta1VirtualMachinesGetResponse_items_guestInfoable)() {
    m.guestInfo = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1VirtualMachinesGetResponse_items) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHighAvailability sets the highAvailability property value. High availability settings of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetHighAvailability(value V1beta1VirtualMachinesGetResponse_items_highAvailabilityable)() {
    m.highAvailability = value
}
// SetHostInfo sets the hostInfo property value. The hostInfo property
func (m *V1beta1VirtualMachinesGetResponse_items) SetHostInfo(value V1beta1VirtualMachinesGetResponse_items_hostInfoable)() {
    m.hostInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1VirtualMachinesGetResponse_items) SetHypervisorManagerInfo(value V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the virtual machine
func (m *V1beta1VirtualMachinesGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetMaxCpuUsage sets the maxCpuUsage property value. Maximum CPU Usage value of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetMaxCpuUsage(value *int64)() {
    m.maxCpuUsage = value
}
// SetName sets the name property value. Name of the virtual machine as configured in the hypervisor manager.
func (m *V1beta1VirtualMachinesGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetNetworkAdapters sets the networkAdapters property value. The networkAdapters property
func (m *V1beta1VirtualMachinesGetResponse_items) SetNetworkAdapters(value []V1beta1VirtualMachinesGetResponse_items_networkAdaptersable)() {
    m.networkAdapters = value
}
// SetNetworkAddress sets the networkAddress property value. IP address of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetNetworkAddresses sets the networkAddresses property value. Network Addresses associated with virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetNetworkAddresses(value []string)() {
    m.networkAddresses = value
}
// SetProtectionJobInfo sets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
func (m *V1beta1VirtualMachinesGetResponse_items) SetProtectionJobInfo(value V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable)() {
    m.protectionJobInfo = value
}
// SetProtectionPolicyAppliedAtInfo sets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
func (m *V1beta1VirtualMachinesGetResponse_items) SetProtectionPolicyAppliedAtInfo(value V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable)() {
    m.protectionPolicyAppliedAtInfo = value
}
// SetRecoveryPointsExist sets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
func (m *V1beta1VirtualMachinesGetResponse_items) SetRecoveryPointsExist(value *bool)() {
    m.recoveryPointsExist = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1VirtualMachinesGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1VirtualMachinesGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the virtual machine
func (m *V1beta1VirtualMachinesGetResponse_items) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageSystemsInfo sets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
func (m *V1beta1VirtualMachinesGetResponse_items) SetStorageSystemsInfo(value []V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable)() {
    m.storageSystemsInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1VirtualMachinesGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. Unique identifier of the virtual machine as reported by the hypervisor.
func (m *V1beta1VirtualMachinesGetResponse_items) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1VirtualMachinesGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVclsVm sets the vclsVm property value. Indicates it is a vCLS virtual machine or not.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVclsVm(value *bool)() {
    m.vclsVm = value
}
// SetVirtualAhciControllers sets the virtualAhciControllers property value. Provides Ahci Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVirtualAhciControllers(value V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable)() {
    m.virtualAhciControllers = value
}
// SetVirtualDisks sets the virtualDisks property value. A list of objects encapsulating information about the storage disks provisioned to a virtual machine
func (m *V1beta1VirtualMachinesGetResponse_items) SetVirtualDisks(value []V1beta1VirtualMachinesGetResponse_items_virtualDisksable)() {
    m.virtualDisks = value
}
// SetVirtualIdeControllers sets the virtualIdeControllers property value. Provides Ide Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVirtualIdeControllers(value V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable)() {
    m.virtualIdeControllers = value
}
// SetVirtualNvmeControllers sets the virtualNvmeControllers property value. Provides Nvme Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVirtualNvmeControllers(value V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable)() {
    m.virtualNvmeControllers = value
}
// SetVirtualScsiControllers sets the virtualScsiControllers property value. Provides SCSI Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVirtualScsiControllers(value V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable)() {
    m.virtualScsiControllers = value
}
// SetVmConfigPath sets the vmConfigPath property value. VM configuration path of the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVmConfigPath(value *string)() {
    m.vmConfigPath = value
}
// SetVmPerfMetricInfo sets the vmPerfMetricInfo property value. Virtual machine performance metrics.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVmPerfMetricInfo(value V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable)() {
    m.vmPerfMetricInfo = value
}
// SetVmProtectionGroupsInfo sets the vmProtectionGroupsInfo property value. Protection groups related to the virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVmProtectionGroupsInfo(value []V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable)() {
    m.vmProtectionGroupsInfo = value
}
// SetVolumesInfo sets the volumesInfo property value. Volumes associated with this virtual machine.
func (m *V1beta1VirtualMachinesGetResponse_items) SetVolumesInfo(value []V1beta1VirtualMachinesGetResponse_items_volumesInfoable)() {
    m.volumesInfo = value
}
type V1beta1VirtualMachinesGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1VirtualMachinesGetResponse_items_appInfoable)
    GetCapacityInBytes()(*int64)
    GetClusterInfo()(V1beta1VirtualMachinesGetResponse_items_clusterInfoable)
    GetComputeInfo()(V1beta1VirtualMachinesGetResponse_items_computeInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetFolderInfo()(V1beta1VirtualMachinesGetResponse_items_folderInfoable)
    GetGeneration()(*int64)
    GetGraphics()(V1beta1VirtualMachinesGetResponse_items_graphicsable)
    GetGuestInfo()(V1beta1VirtualMachinesGetResponse_items_guestInfoable)
    GetHciClusterUuid()(*string)
    GetHighAvailability()(V1beta1VirtualMachinesGetResponse_items_highAvailabilityable)
    GetHostInfo()(V1beta1VirtualMachinesGetResponse_items_hostInfoable)
    GetHypervisorManagerInfo()(V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable)
    GetId()(*string)
    GetMaxCpuUsage()(*int64)
    GetName()(*string)
    GetNetworkAdapters()([]V1beta1VirtualMachinesGetResponse_items_networkAdaptersable)
    GetNetworkAddress()(*string)
    GetNetworkAddresses()([]string)
    GetProtectionJobInfo()(V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable)
    GetProtectionPolicyAppliedAtInfo()(V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable)
    GetRecoveryPointsExist()(*bool)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetStorageSystemsInfo()([]V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVclsVm()(*bool)
    GetVirtualAhciControllers()(V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable)
    GetVirtualDisks()([]V1beta1VirtualMachinesGetResponse_items_virtualDisksable)
    GetVirtualIdeControllers()(V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable)
    GetVirtualNvmeControllers()(V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable)
    GetVirtualScsiControllers()(V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable)
    GetVmConfigPath()(*string)
    GetVmPerfMetricInfo()(V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable)
    GetVmProtectionGroupsInfo()([]V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable)
    GetVolumesInfo()([]V1beta1VirtualMachinesGetResponse_items_volumesInfoable)
    SetAppInfo(value V1beta1VirtualMachinesGetResponse_items_appInfoable)()
    SetCapacityInBytes(value *int64)()
    SetClusterInfo(value V1beta1VirtualMachinesGetResponse_items_clusterInfoable)()
    SetComputeInfo(value V1beta1VirtualMachinesGetResponse_items_computeInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetFolderInfo(value V1beta1VirtualMachinesGetResponse_items_folderInfoable)()
    SetGeneration(value *int64)()
    SetGraphics(value V1beta1VirtualMachinesGetResponse_items_graphicsable)()
    SetGuestInfo(value V1beta1VirtualMachinesGetResponse_items_guestInfoable)()
    SetHciClusterUuid(value *string)()
    SetHighAvailability(value V1beta1VirtualMachinesGetResponse_items_highAvailabilityable)()
    SetHostInfo(value V1beta1VirtualMachinesGetResponse_items_hostInfoable)()
    SetHypervisorManagerInfo(value V1beta1VirtualMachinesGetResponse_items_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetMaxCpuUsage(value *int64)()
    SetName(value *string)()
    SetNetworkAdapters(value []V1beta1VirtualMachinesGetResponse_items_networkAdaptersable)()
    SetNetworkAddress(value *string)()
    SetNetworkAddresses(value []string)()
    SetProtectionJobInfo(value V1beta1VirtualMachinesGetResponse_items_protectionJobInfoable)()
    SetProtectionPolicyAppliedAtInfo(value V1beta1VirtualMachinesGetResponse_items_protectionPolicyAppliedAtInfoable)()
    SetRecoveryPointsExist(value *bool)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetStorageSystemsInfo(value []V1beta1VirtualMachinesGetResponse_items_storageSystemsInfoable)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVclsVm(value *bool)()
    SetVirtualAhciControllers(value V1beta1VirtualMachinesGetResponse_items_virtualAhciControllersable)()
    SetVirtualDisks(value []V1beta1VirtualMachinesGetResponse_items_virtualDisksable)()
    SetVirtualIdeControllers(value V1beta1VirtualMachinesGetResponse_items_virtualIdeControllersable)()
    SetVirtualNvmeControllers(value V1beta1VirtualMachinesGetResponse_items_virtualNvmeControllersable)()
    SetVirtualScsiControllers(value V1beta1VirtualMachinesGetResponse_items_virtualScsiControllersable)()
    SetVmConfigPath(value *string)()
    SetVmPerfMetricInfo(value V1beta1VirtualMachinesGetResponse_items_vmPerfMetricInfoable)()
    SetVmProtectionGroupsInfo(value []V1beta1VirtualMachinesGetResponse_items_vmProtectionGroupsInfoable)()
    SetVolumesInfo(value []V1beta1VirtualMachinesGetResponse_items_volumesInfoable)()
}
