package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemVirtualMachinesGetResponse represents a single instance of a Virtual Machine
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Application specific information for this virtual machine.
    appInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable
    // Size of the virtual machine in bytes.
    capacityInBytes *int64
    // The clusterInfo property
    clusterInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable
    // Compute information of the virtual machine.
    computeInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // A user-friendly name that identifies the virtual machine.
    displayName *string
    // The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
    folderInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Graphics configuration for the virtual machine.
    graphics V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable
    // Information of this guest OS running on the virtual machine.
    guestInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // High availability settings of the virtual machine.
    highAvailability V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable
    // The hostInfo property
    hostInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable
    // UUID string uniquely identifying the virtual machine
    id *string
    // Maximum CPU Usage value of the virtual machine.
    maxCpuUsage *int64
    // Name of the virtual machine as configured in the hypervisor manager.
    name *string
    // The networkAdapters property
    networkAdapters []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable
    // IP address of the virtual machine.
    networkAddress *string
    // Network Addresses associated with virtual machine.
    networkAddresses []string
    // Information about the assigned Protection Policy and the Protection Job.
    protectionJobInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable
    // Describes applied protection policy information.
    protectionPolicyAppliedAtInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable
    // Indicates at least one valid recovery point exists for this resource.
    recoveryPointsExist *bool
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the virtual machine
    stateReason *string
    // storage systems with volume wwn associated with this virtual machine
    storageSystemsInfo []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable
    // The type of resource.
    typeEscaped *string
    // Unique identifier of the virtual machine as reported by the hypervisor.
    uid *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates it is a vCLS virtual machine or not.
    vclsVm *bool
    // Provides Ahci Disk Controller details of the virtual machine.
    virtualAhciControllers V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable
    // A list of objects encapsulating information about the storage disks provisioned to a virtual machine
    virtualDisks []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable
    // Provides Ide Disk Controller details of the virtual machine.
    virtualIdeControllers V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable
    // Provides Nvme Disk Controller details of the virtual machine.
    virtualNvmeControllers V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable
    // Provides SCSI Disk Controller details of the virtual machine.
    virtualScsiControllers V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable
    // VM configuration path of the virtual machine.
    vmConfigPath *string
    // Virtual machine performance metrics.
    vmPerfMetricInfo V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable
    // Protection groups related to the virtual machine.
    vmProtectionGroupsInfo []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable
    // Volumes associated with this virtual machine.
    volumesInfo []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Application specific information for this virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetAppInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable) {
    return m.appInfo
}
// GetCapacityInBytes gets the capacityInBytes property value. Size of the virtual machine in bytes.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetCapacityInBytes()(*int64) {
    return m.capacityInBytes
}
// GetClusterInfo gets the clusterInfo property value. The clusterInfo property
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetClusterInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable) {
    return m.clusterInfo
}
// GetComputeInfo gets the computeInfo property value. Compute information of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetComputeInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable) {
    return m.computeInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable))
        }
        return nil
    }
    res["computeInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGraphics(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable))
        }
        return nil
    }
    res["guestInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighAvailability(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable))
        }
        return nil
    }
    res["hostInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable))
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable)
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionJobInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable))
        }
        return nil
    }
    res["protectionPolicyAppliedAtInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyAppliedAtInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable)
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualAhciControllers(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable))
        }
        return nil
    }
    res["virtualDisks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable)
                }
            }
            m.SetVirtualDisks(res)
        }
        return nil
    }
    res["virtualIdeControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualIdeControllers(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable))
        }
        return nil
    }
    res["virtualNvmeControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNvmeControllers(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable))
        }
        return nil
    }
    res["virtualScsiControllers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualScsiControllers(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable))
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
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmPerfMetricInfo(val.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable))
        }
        return nil
    }
    res["vmProtectionGroupsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable)
                }
            }
            m.SetVmProtectionGroupsInfo(res)
        }
        return nil
    }
    res["volumesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable)
                }
            }
            m.SetVolumesInfo(res)
        }
        return nil
    }
    return res
}
// GetFolderInfo gets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetFolderInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable) {
    return m.folderInfo
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetGraphics gets the graphics property value. Graphics configuration for the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetGraphics()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable) {
    return m.graphics
}
// GetGuestInfo gets the guestInfo property value. Information of this guest OS running on the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetGuestInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable) {
    return m.guestInfo
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHighAvailability gets the highAvailability property value. High availability settings of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetHighAvailability()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable) {
    return m.highAvailability
}
// GetHostInfo gets the hostInfo property value. The hostInfo property
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetHostInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable) {
    return m.hostInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetHypervisorManagerInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetId()(*string) {
    return m.id
}
// GetMaxCpuUsage gets the maxCpuUsage property value. Maximum CPU Usage value of the virtual machine.
// returns a *int64 when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetMaxCpuUsage()(*int64) {
    return m.maxCpuUsage
}
// GetName gets the name property value. Name of the virtual machine as configured in the hypervisor manager.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetName()(*string) {
    return m.name
}
// GetNetworkAdapters gets the networkAdapters property value. The networkAdapters property
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetNetworkAdapters()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable) {
    return m.networkAdapters
}
// GetNetworkAddress gets the networkAddress property value. IP address of the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetNetworkAddresses gets the networkAddresses property value. Network Addresses associated with virtual machine.
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetNetworkAddresses()([]string) {
    return m.networkAddresses
}
// GetProtectionJobInfo gets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetProtectionJobInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable) {
    return m.protectionJobInfo
}
// GetProtectionPolicyAppliedAtInfo gets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetProtectionPolicyAppliedAtInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable) {
    return m.protectionPolicyAppliedAtInfo
}
// GetRecoveryPointsExist gets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetRecoveryPointsExist()(*bool) {
    return m.recoveryPointsExist
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the virtual machine
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageSystemsInfo gets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetStorageSystemsInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable) {
    return m.storageSystemsInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. Unique identifier of the virtual machine as reported by the hypervisor.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVclsVm gets the vclsVm property value. Indicates it is a vCLS virtual machine or not.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVclsVm()(*bool) {
    return m.vclsVm
}
// GetVirtualAhciControllers gets the virtualAhciControllers property value. Provides Ahci Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVirtualAhciControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable) {
    return m.virtualAhciControllers
}
// GetVirtualDisks gets the virtualDisks property value. A list of objects encapsulating information about the storage disks provisioned to a virtual machine
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVirtualDisks()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable) {
    return m.virtualDisks
}
// GetVirtualIdeControllers gets the virtualIdeControllers property value. Provides Ide Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVirtualIdeControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable) {
    return m.virtualIdeControllers
}
// GetVirtualNvmeControllers gets the virtualNvmeControllers property value. Provides Nvme Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVirtualNvmeControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable) {
    return m.virtualNvmeControllers
}
// GetVirtualScsiControllers gets the virtualScsiControllers property value. Provides SCSI Disk Controller details of the virtual machine.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVirtualScsiControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable) {
    return m.virtualScsiControllers
}
// GetVmConfigPath gets the vmConfigPath property value. VM configuration path of the virtual machine.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVmConfigPath()(*string) {
    return m.vmConfigPath
}
// GetVmPerfMetricInfo gets the vmPerfMetricInfo property value. Virtual machine performance metrics.
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVmPerfMetricInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable) {
    return m.vmPerfMetricInfo
}
// GetVmProtectionGroupsInfo gets the vmProtectionGroupsInfo property value. Protection groups related to the virtual machine.
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVmProtectionGroupsInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable) {
    return m.vmProtectionGroupsInfo
}
// GetVolumesInfo gets the volumesInfo property value. Volumes associated with this virtual machine.
// returns a []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) GetVolumesInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable) {
    return m.volumesInfo
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Application specific information for this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetAppInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable)() {
    m.appInfo = value
}
// SetCapacityInBytes sets the capacityInBytes property value. Size of the virtual machine in bytes.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetCapacityInBytes(value *int64)() {
    m.capacityInBytes = value
}
// SetClusterInfo sets the clusterInfo property value. The clusterInfo property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetClusterInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable)() {
    m.clusterInfo = value
}
// SetComputeInfo sets the computeInfo property value. Compute information of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetComputeInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable)() {
    m.computeInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFolderInfo sets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetFolderInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable)() {
    m.folderInfo = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGraphics sets the graphics property value. Graphics configuration for the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetGraphics(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable)() {
    m.graphics = value
}
// SetGuestInfo sets the guestInfo property value. Information of this guest OS running on the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetGuestInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable)() {
    m.guestInfo = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHighAvailability sets the highAvailability property value. High availability settings of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetHighAvailability(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable)() {
    m.highAvailability = value
}
// SetHostInfo sets the hostInfo property value. The hostInfo property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetHostInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable)() {
    m.hostInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetHypervisorManagerInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the virtual machine
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetId(value *string)() {
    m.id = value
}
// SetMaxCpuUsage sets the maxCpuUsage property value. Maximum CPU Usage value of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetMaxCpuUsage(value *int64)() {
    m.maxCpuUsage = value
}
// SetName sets the name property value. Name of the virtual machine as configured in the hypervisor manager.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetName(value *string)() {
    m.name = value
}
// SetNetworkAdapters sets the networkAdapters property value. The networkAdapters property
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetNetworkAdapters(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable)() {
    m.networkAdapters = value
}
// SetNetworkAddress sets the networkAddress property value. IP address of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetNetworkAddresses sets the networkAddresses property value. Network Addresses associated with virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetNetworkAddresses(value []string)() {
    m.networkAddresses = value
}
// SetProtectionJobInfo sets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetProtectionJobInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable)() {
    m.protectionJobInfo = value
}
// SetProtectionPolicyAppliedAtInfo sets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetProtectionPolicyAppliedAtInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable)() {
    m.protectionPolicyAppliedAtInfo = value
}
// SetRecoveryPointsExist sets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetRecoveryPointsExist(value *bool)() {
    m.recoveryPointsExist = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the virtual machine
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageSystemsInfo sets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetStorageSystemsInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable)() {
    m.storageSystemsInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. Unique identifier of the virtual machine as reported by the hypervisor.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVclsVm sets the vclsVm property value. Indicates it is a vCLS virtual machine or not.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVclsVm(value *bool)() {
    m.vclsVm = value
}
// SetVirtualAhciControllers sets the virtualAhciControllers property value. Provides Ahci Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVirtualAhciControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable)() {
    m.virtualAhciControllers = value
}
// SetVirtualDisks sets the virtualDisks property value. A list of objects encapsulating information about the storage disks provisioned to a virtual machine
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVirtualDisks(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable)() {
    m.virtualDisks = value
}
// SetVirtualIdeControllers sets the virtualIdeControllers property value. Provides Ide Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVirtualIdeControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable)() {
    m.virtualIdeControllers = value
}
// SetVirtualNvmeControllers sets the virtualNvmeControllers property value. Provides Nvme Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVirtualNvmeControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable)() {
    m.virtualNvmeControllers = value
}
// SetVirtualScsiControllers sets the virtualScsiControllers property value. Provides SCSI Disk Controller details of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVirtualScsiControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable)() {
    m.virtualScsiControllers = value
}
// SetVmConfigPath sets the vmConfigPath property value. VM configuration path of the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVmConfigPath(value *string)() {
    m.vmConfigPath = value
}
// SetVmPerfMetricInfo sets the vmPerfMetricInfo property value. Virtual machine performance metrics.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVmPerfMetricInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable)() {
    m.vmPerfMetricInfo = value
}
// SetVmProtectionGroupsInfo sets the vmProtectionGroupsInfo property value. Protection groups related to the virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVmProtectionGroupsInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable)() {
    m.vmProtectionGroupsInfo = value
}
// SetVolumesInfo sets the volumesInfo property value. Volumes associated with this virtual machine.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse) SetVolumesInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable)() {
    m.volumesInfo = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable)
    GetCapacityInBytes()(*int64)
    GetClusterInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable)
    GetComputeInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetFolderInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable)
    GetGeneration()(*int64)
    GetGraphics()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable)
    GetGuestInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable)
    GetHciClusterUuid()(*string)
    GetHighAvailability()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable)
    GetHostInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable)
    GetHypervisorManagerInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable)
    GetId()(*string)
    GetMaxCpuUsage()(*int64)
    GetName()(*string)
    GetNetworkAdapters()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable)
    GetNetworkAddress()(*string)
    GetNetworkAddresses()([]string)
    GetProtectionJobInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable)
    GetProtectionPolicyAppliedAtInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable)
    GetRecoveryPointsExist()(*bool)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetStorageSystemsInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVclsVm()(*bool)
    GetVirtualAhciControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable)
    GetVirtualDisks()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable)
    GetVirtualIdeControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable)
    GetVirtualNvmeControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable)
    GetVirtualScsiControllers()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable)
    GetVmConfigPath()(*string)
    GetVmPerfMetricInfo()(V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable)
    GetVmProtectionGroupsInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable)
    GetVolumesInfo()([]V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable)
    SetAppInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_appInfoable)()
    SetCapacityInBytes(value *int64)()
    SetClusterInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_clusterInfoable)()
    SetComputeInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_computeInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetFolderInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_folderInfoable)()
    SetGeneration(value *int64)()
    SetGraphics(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_graphicsable)()
    SetGuestInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_guestInfoable)()
    SetHciClusterUuid(value *string)()
    SetHighAvailability(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_highAvailabilityable)()
    SetHostInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hostInfoable)()
    SetHypervisorManagerInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetMaxCpuUsage(value *int64)()
    SetName(value *string)()
    SetNetworkAdapters(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdaptersable)()
    SetNetworkAddress(value *string)()
    SetNetworkAddresses(value []string)()
    SetProtectionJobInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionJobInfoable)()
    SetProtectionPolicyAppliedAtInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_protectionPolicyAppliedAtInfoable)()
    SetRecoveryPointsExist(value *bool)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetStorageSystemsInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_storageSystemsInfoable)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVclsVm(value *bool)()
    SetVirtualAhciControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualAhciControllersable)()
    SetVirtualDisks(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualDisksable)()
    SetVirtualIdeControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualIdeControllersable)()
    SetVirtualNvmeControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualNvmeControllersable)()
    SetVirtualScsiControllers(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_virtualScsiControllersable)()
    SetVmConfigPath(value *string)()
    SetVmPerfMetricInfo(value V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmPerfMetricInfoable)()
    SetVmProtectionGroupsInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_vmProtectionGroupsInfoable)()
    SetVolumesInfo(value []V1beta1VirtualMachinesItemVirtualMachinesGetResponse_volumesInfoable)()
}
