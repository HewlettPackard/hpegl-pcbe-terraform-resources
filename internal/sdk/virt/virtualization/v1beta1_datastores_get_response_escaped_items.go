package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresGetResponse_items represents a single instance of a Datastore
type V1beta1DatastoresGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unused storage of the datastore in bytes.
    capacityFree *int64
    // Size of the datastore in bytes.
    capacityInBytes *int64
    // Uncommitted storage of the datastore in bytes.
    capacityUncommitted *int64
    // The clusterInfo property
    clusterInfo V1beta1DatastoresGetResponse_items_clusterInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // List of datacenters to which the datastore is presented to.
    datacentersInfo []V1beta1DatastoresGetResponse_items_datacentersInfoable
    // A user-friendly name that identifies the datastore.
    displayName *string
    // The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
    folderInfo V1beta1DatastoresGetResponse_items_folderInfoable
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // List of hypervisor hosts to which the datastore is presented to.
    hostsInfo []V1beta1DatastoresGetResponse_items_hostsInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable
    // UUID string uniquely identifying the datastore
    id *string
    // VMware provided moref of the data store
    moref *string
    // Name of the datastore as reported by the hypervisor manager.
    name *string
    // Information about the assigned Protection Policy and the Protection Job.
    protectionJobInfo V1beta1DatastoresGetResponse_items_protectionJobInfoable
    // Describes applied protection policy information.
    protectionPolicyAppliedAtInfo V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable
    // Describes provisioning policy information.
    provisioningPolicyInfo V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable
    // Indicates at least one valid recovery point exists for this resource.
    recoveryPointsExist *bool
    // Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
    replicationInfo V1beta1DatastoresGetResponse_items_replicationInfoable
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the datastore
    stateReason *string
    // storage systems with volume wwn associated with this virtual machine
    storageSystemsInfo []V1beta1DatastoresGetResponse_items_storageSystemsInfoable
    // The type of resource.
    typeEscaped *string
    // VMware provided uuid of the datastore.
    uid *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of virtual machines associated with the datastore.
    vmCount *int64
    // List of virtual machine protection groups.
    vmProtectionGroupsInfo []V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable
    // Volumes associated with datastore.
    volumesInfo []V1beta1DatastoresGetResponse_items_volumesInfoable
}
// NewV1beta1DatastoresGetResponse_items instantiates a new V1beta1DatastoresGetResponse_items and sets the default values.
func NewV1beta1DatastoresGetResponse_items()(*V1beta1DatastoresGetResponse_items) {
    m := &V1beta1DatastoresGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapacityFree gets the capacityFree property value. Unused storage of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresGetResponse_items) GetCapacityFree()(*int64) {
    return m.capacityFree
}
// GetCapacityInBytes gets the capacityInBytes property value. Size of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresGetResponse_items) GetCapacityInBytes()(*int64) {
    return m.capacityInBytes
}
// GetCapacityUncommitted gets the capacityUncommitted property value. Uncommitted storage of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresGetResponse_items) GetCapacityUncommitted()(*int64) {
    return m.capacityUncommitted
}
// GetClusterInfo gets the clusterInfo property value. The clusterInfo property
// returns a V1beta1DatastoresGetResponse_items_clusterInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetClusterInfo()(V1beta1DatastoresGetResponse_items_clusterInfoable) {
    return m.clusterInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1DatastoresGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDatacentersInfo gets the datacentersInfo property value. List of datacenters to which the datastore is presented to.
// returns a []V1beta1DatastoresGetResponse_items_datacentersInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetDatacentersInfo()([]V1beta1DatastoresGetResponse_items_datacentersInfoable) {
    return m.datacentersInfo
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the datastore.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capacityFree"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityFree(val)
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
    res["capacityUncommitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityUncommitted(val)
        }
        return nil
    }
    res["clusterInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_clusterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterInfo(val.(V1beta1DatastoresGetResponse_items_clusterInfoable))
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
    res["datacentersInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresGetResponse_items_datacentersInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresGetResponse_items_datacentersInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresGetResponse_items_datacentersInfoable)
                }
            }
            m.SetDatacentersInfo(res)
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
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_folderInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderInfo(val.(V1beta1DatastoresGetResponse_items_folderInfoable))
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
    res["hostsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresGetResponse_items_hostsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresGetResponse_items_hostsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresGetResponse_items_hostsInfoable)
                }
            }
            m.SetHostsInfo(res)
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable))
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
    res["moref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoref(val)
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
    res["protectionJobInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_protectionJobInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionJobInfo(val.(V1beta1DatastoresGetResponse_items_protectionJobInfoable))
        }
        return nil
    }
    res["protectionPolicyAppliedAtInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyAppliedAtInfo(val.(V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable))
        }
        return nil
    }
    res["provisioningPolicyInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_provisioningPolicyInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyInfo(val.(V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable))
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
    res["replicationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresGetResponse_items_replicationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplicationInfo(val.(V1beta1DatastoresGetResponse_items_replicationInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresGetResponse_items_storageSystemsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresGetResponse_items_storageSystemsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresGetResponse_items_storageSystemsInfoable)
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
    res["vmCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmCount(val)
        }
        return nil
    }
    res["vmProtectionGroupsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable)
                }
            }
            m.SetVmProtectionGroupsInfo(res)
        }
        return nil
    }
    res["volumesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresGetResponse_items_volumesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresGetResponse_items_volumesInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresGetResponse_items_volumesInfoable)
                }
            }
            m.SetVolumesInfo(res)
        }
        return nil
    }
    return res
}
// GetFolderInfo gets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
// returns a V1beta1DatastoresGetResponse_items_folderInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetFolderInfo()(V1beta1DatastoresGetResponse_items_folderInfoable) {
    return m.folderInfo
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1DatastoresGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHostsInfo gets the hostsInfo property value. List of hypervisor hosts to which the datastore is presented to.
// returns a []V1beta1DatastoresGetResponse_items_hostsInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetHostsInfo()([]V1beta1DatastoresGetResponse_items_hostsInfoable) {
    return m.hostsInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetHypervisorManagerInfo()(V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the datastore
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetId()(*string) {
    return m.id
}
// GetMoref gets the moref property value. VMware provided moref of the data store
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetMoref()(*string) {
    return m.moref
}
// GetName gets the name property value. Name of the datastore as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetName()(*string) {
    return m.name
}
// GetProtectionJobInfo gets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
// returns a V1beta1DatastoresGetResponse_items_protectionJobInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetProtectionJobInfo()(V1beta1DatastoresGetResponse_items_protectionJobInfoable) {
    return m.protectionJobInfo
}
// GetProtectionPolicyAppliedAtInfo gets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
// returns a V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetProtectionPolicyAppliedAtInfo()(V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable) {
    return m.protectionPolicyAppliedAtInfo
}
// GetProvisioningPolicyInfo gets the provisioningPolicyInfo property value. Describes provisioning policy information.
// returns a V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetProvisioningPolicyInfo()(V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable) {
    return m.provisioningPolicyInfo
}
// GetRecoveryPointsExist gets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
// returns a *bool when successful
func (m *V1beta1DatastoresGetResponse_items) GetRecoveryPointsExist()(*bool) {
    return m.recoveryPointsExist
}
// GetReplicationInfo gets the replicationInfo property value. Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
// returns a V1beta1DatastoresGetResponse_items_replicationInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetReplicationInfo()(V1beta1DatastoresGetResponse_items_replicationInfoable) {
    return m.replicationInfo
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1DatastoresGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the datastore
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageSystemsInfo gets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
// returns a []V1beta1DatastoresGetResponse_items_storageSystemsInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetStorageSystemsInfo()([]V1beta1DatastoresGetResponse_items_storageSystemsInfoable) {
    return m.storageSystemsInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. VMware provided uuid of the datastore.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1DatastoresGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVmCount gets the vmCount property value. Number of virtual machines associated with the datastore.
// returns a *int64 when successful
func (m *V1beta1DatastoresGetResponse_items) GetVmCount()(*int64) {
    return m.vmCount
}
// GetVmProtectionGroupsInfo gets the vmProtectionGroupsInfo property value. List of virtual machine protection groups.
// returns a []V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetVmProtectionGroupsInfo()([]V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable) {
    return m.vmProtectionGroupsInfo
}
// GetVolumesInfo gets the volumesInfo property value. Volumes associated with datastore.
// returns a []V1beta1DatastoresGetResponse_items_volumesInfoable when successful
func (m *V1beta1DatastoresGetResponse_items) GetVolumesInfo()([]V1beta1DatastoresGetResponse_items_volumesInfoable) {
    return m.volumesInfo
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("capacityFree", m.GetCapacityFree())
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
        err := writer.WriteInt64Value("capacityUncommitted", m.GetCapacityUncommitted())
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
    if m.GetDatacentersInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDatacentersInfo()))
        for i, v := range m.GetDatacentersInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("datacentersInfo", cast)
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
        err := writer.WriteStringValue("hciClusterUuid", m.GetHciClusterUuid())
        if err != nil {
            return err
        }
    }
    if m.GetHostsInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostsInfo()))
        for i, v := range m.GetHostsInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("hostsInfo", cast)
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
        err := writer.WriteStringValue("moref", m.GetMoref())
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
        err := writer.WriteObjectValue("provisioningPolicyInfo", m.GetProvisioningPolicyInfo())
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
        err := writer.WriteObjectValue("replicationInfo", m.GetReplicationInfo())
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
        err := writer.WriteInt64Value("vmCount", m.GetVmCount())
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
func (m *V1beta1DatastoresGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapacityFree sets the capacityFree property value. Unused storage of the datastore in bytes.
func (m *V1beta1DatastoresGetResponse_items) SetCapacityFree(value *int64)() {
    m.capacityFree = value
}
// SetCapacityInBytes sets the capacityInBytes property value. Size of the datastore in bytes.
func (m *V1beta1DatastoresGetResponse_items) SetCapacityInBytes(value *int64)() {
    m.capacityInBytes = value
}
// SetCapacityUncommitted sets the capacityUncommitted property value. Uncommitted storage of the datastore in bytes.
func (m *V1beta1DatastoresGetResponse_items) SetCapacityUncommitted(value *int64)() {
    m.capacityUncommitted = value
}
// SetClusterInfo sets the clusterInfo property value. The clusterInfo property
func (m *V1beta1DatastoresGetResponse_items) SetClusterInfo(value V1beta1DatastoresGetResponse_items_clusterInfoable)() {
    m.clusterInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1DatastoresGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1DatastoresGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDatacentersInfo sets the datacentersInfo property value. List of datacenters to which the datastore is presented to.
func (m *V1beta1DatastoresGetResponse_items) SetDatacentersInfo(value []V1beta1DatastoresGetResponse_items_datacentersInfoable)() {
    m.datacentersInfo = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the datastore.
func (m *V1beta1DatastoresGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFolderInfo sets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
func (m *V1beta1DatastoresGetResponse_items) SetFolderInfo(value V1beta1DatastoresGetResponse_items_folderInfoable)() {
    m.folderInfo = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1DatastoresGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1DatastoresGetResponse_items) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHostsInfo sets the hostsInfo property value. List of hypervisor hosts to which the datastore is presented to.
func (m *V1beta1DatastoresGetResponse_items) SetHostsInfo(value []V1beta1DatastoresGetResponse_items_hostsInfoable)() {
    m.hostsInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1DatastoresGetResponse_items) SetHypervisorManagerInfo(value V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the datastore
func (m *V1beta1DatastoresGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetMoref sets the moref property value. VMware provided moref of the data store
func (m *V1beta1DatastoresGetResponse_items) SetMoref(value *string)() {
    m.moref = value
}
// SetName sets the name property value. Name of the datastore as reported by the hypervisor manager.
func (m *V1beta1DatastoresGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetProtectionJobInfo sets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
func (m *V1beta1DatastoresGetResponse_items) SetProtectionJobInfo(value V1beta1DatastoresGetResponse_items_protectionJobInfoable)() {
    m.protectionJobInfo = value
}
// SetProtectionPolicyAppliedAtInfo sets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
func (m *V1beta1DatastoresGetResponse_items) SetProtectionPolicyAppliedAtInfo(value V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable)() {
    m.protectionPolicyAppliedAtInfo = value
}
// SetProvisioningPolicyInfo sets the provisioningPolicyInfo property value. Describes provisioning policy information.
func (m *V1beta1DatastoresGetResponse_items) SetProvisioningPolicyInfo(value V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable)() {
    m.provisioningPolicyInfo = value
}
// SetRecoveryPointsExist sets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
func (m *V1beta1DatastoresGetResponse_items) SetRecoveryPointsExist(value *bool)() {
    m.recoveryPointsExist = value
}
// SetReplicationInfo sets the replicationInfo property value. Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
func (m *V1beta1DatastoresGetResponse_items) SetReplicationInfo(value V1beta1DatastoresGetResponse_items_replicationInfoable)() {
    m.replicationInfo = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1DatastoresGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1DatastoresGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the datastore
func (m *V1beta1DatastoresGetResponse_items) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageSystemsInfo sets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
func (m *V1beta1DatastoresGetResponse_items) SetStorageSystemsInfo(value []V1beta1DatastoresGetResponse_items_storageSystemsInfoable)() {
    m.storageSystemsInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DatastoresGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. VMware provided uuid of the datastore.
func (m *V1beta1DatastoresGetResponse_items) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1DatastoresGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVmCount sets the vmCount property value. Number of virtual machines associated with the datastore.
func (m *V1beta1DatastoresGetResponse_items) SetVmCount(value *int64)() {
    m.vmCount = value
}
// SetVmProtectionGroupsInfo sets the vmProtectionGroupsInfo property value. List of virtual machine protection groups.
func (m *V1beta1DatastoresGetResponse_items) SetVmProtectionGroupsInfo(value []V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable)() {
    m.vmProtectionGroupsInfo = value
}
// SetVolumesInfo sets the volumesInfo property value. Volumes associated with datastore.
func (m *V1beta1DatastoresGetResponse_items) SetVolumesInfo(value []V1beta1DatastoresGetResponse_items_volumesInfoable)() {
    m.volumesInfo = value
}
type V1beta1DatastoresGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacityFree()(*int64)
    GetCapacityInBytes()(*int64)
    GetCapacityUncommitted()(*int64)
    GetClusterInfo()(V1beta1DatastoresGetResponse_items_clusterInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDatacentersInfo()([]V1beta1DatastoresGetResponse_items_datacentersInfoable)
    GetDisplayName()(*string)
    GetFolderInfo()(V1beta1DatastoresGetResponse_items_folderInfoable)
    GetGeneration()(*int64)
    GetHciClusterUuid()(*string)
    GetHostsInfo()([]V1beta1DatastoresGetResponse_items_hostsInfoable)
    GetHypervisorManagerInfo()(V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable)
    GetId()(*string)
    GetMoref()(*string)
    GetName()(*string)
    GetProtectionJobInfo()(V1beta1DatastoresGetResponse_items_protectionJobInfoable)
    GetProtectionPolicyAppliedAtInfo()(V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable)
    GetProvisioningPolicyInfo()(V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable)
    GetRecoveryPointsExist()(*bool)
    GetReplicationInfo()(V1beta1DatastoresGetResponse_items_replicationInfoable)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetStorageSystemsInfo()([]V1beta1DatastoresGetResponse_items_storageSystemsInfoable)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVmCount()(*int64)
    GetVmProtectionGroupsInfo()([]V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable)
    GetVolumesInfo()([]V1beta1DatastoresGetResponse_items_volumesInfoable)
    SetCapacityFree(value *int64)()
    SetCapacityInBytes(value *int64)()
    SetCapacityUncommitted(value *int64)()
    SetClusterInfo(value V1beta1DatastoresGetResponse_items_clusterInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDatacentersInfo(value []V1beta1DatastoresGetResponse_items_datacentersInfoable)()
    SetDisplayName(value *string)()
    SetFolderInfo(value V1beta1DatastoresGetResponse_items_folderInfoable)()
    SetGeneration(value *int64)()
    SetHciClusterUuid(value *string)()
    SetHostsInfo(value []V1beta1DatastoresGetResponse_items_hostsInfoable)()
    SetHypervisorManagerInfo(value V1beta1DatastoresGetResponse_items_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetMoref(value *string)()
    SetName(value *string)()
    SetProtectionJobInfo(value V1beta1DatastoresGetResponse_items_protectionJobInfoable)()
    SetProtectionPolicyAppliedAtInfo(value V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable)()
    SetProvisioningPolicyInfo(value V1beta1DatastoresGetResponse_items_provisioningPolicyInfoable)()
    SetRecoveryPointsExist(value *bool)()
    SetReplicationInfo(value V1beta1DatastoresGetResponse_items_replicationInfoable)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetStorageSystemsInfo(value []V1beta1DatastoresGetResponse_items_storageSystemsInfoable)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVmCount(value *int64)()
    SetVmProtectionGroupsInfo(value []V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable)()
    SetVolumesInfo(value []V1beta1DatastoresGetResponse_items_volumesInfoable)()
}
