package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresItemDatastoresGetResponse represents a single instance of a Datastore
type V1beta1DatastoresItemDatastoresGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unused storage of the datastore in bytes.
    capacityFree *int64
    // Size of the datastore in bytes.
    capacityInBytes *int64
    // Uncommitted storage of the datastore in bytes.
    capacityUncommitted *int64
    // The clusterInfo property
    clusterInfo V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // List of datacenters to which the datastore is presented to.
    datacentersInfo []V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable
    // A user-friendly name that identifies the datastore.
    displayName *string
    // The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
    folderInfo V1beta1DatastoresItemDatastoresGetResponse_folderInfoable
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // UUID string uniquely identifying the HCI cluster.
    hciClusterUuid *string
    // List of hypervisor hosts to which the datastore is presented to.
    hostsInfo []V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable
    // UUID string uniquely identifying the datastore
    id *string
    // VMware provided moref of the data store
    moref *string
    // Name of the datastore as reported by the hypervisor manager.
    name *string
    // Information about the assigned Protection Policy and the Protection Job.
    protectionJobInfo V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable
    // Describes applied protection policy information.
    protectionPolicyAppliedAtInfo V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable
    // Describes provisioning policy information.
    provisioningPolicyInfo V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable
    // Indicates at least one valid recovery point exists for this resource.
    recoveryPointsExist *bool
    // Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
    replicationInfo V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the datastore
    stateReason *string
    // storage systems with volume wwn associated with this virtual machine
    storageSystemsInfo []V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable
    // The type of resource.
    typeEscaped *string
    // VMware provided uuid of the datastore.
    uid *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of virtual machines associated with the datastore.
    vmCount *int64
    // List of virtual machine protection groups.
    vmProtectionGroupsInfo []V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable
    // Volumes associated with datastore.
    volumesInfo []V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable
}
// NewV1beta1DatastoresItemDatastoresGetResponse instantiates a new V1beta1DatastoresItemDatastoresGetResponse and sets the default values.
func NewV1beta1DatastoresItemDatastoresGetResponse()(*V1beta1DatastoresItemDatastoresGetResponse) {
    m := &V1beta1DatastoresItemDatastoresGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresItemDatastoresGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemDatastoresGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemDatastoresGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapacityFree gets the capacityFree property value. Unused storage of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetCapacityFree()(*int64) {
    return m.capacityFree
}
// GetCapacityInBytes gets the capacityInBytes property value. Size of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetCapacityInBytes()(*int64) {
    return m.capacityInBytes
}
// GetCapacityUncommitted gets the capacityUncommitted property value. Uncommitted storage of the datastore in bytes.
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetCapacityUncommitted()(*int64) {
    return m.capacityUncommitted
}
// GetClusterInfo gets the clusterInfo property value. The clusterInfo property
// returns a V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetClusterInfo()(V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable) {
    return m.clusterInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDatacentersInfo gets the datacentersInfo property value. List of datacenters to which the datastore is presented to.
// returns a []V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetDatacentersInfo()([]V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable) {
    return m.datacentersInfo
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the datastore.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_clusterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_datacentersInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable)
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
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_folderInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_folderInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_hostsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable)
                }
            }
            m.SetHostsInfo(res)
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionJobInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable))
        }
        return nil
    }
    res["protectionPolicyAppliedAtInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyAppliedAtInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable))
        }
        return nil
    }
    res["provisioningPolicyInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable))
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
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresGetResponse_replicationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplicationInfo(val.(V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable))
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable)
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable)
                }
            }
            m.SetVmProtectionGroupsInfo(res)
        }
        return nil
    }
    res["volumesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DatastoresItemDatastoresGetResponse_volumesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable)
                }
            }
            m.SetVolumesInfo(res)
        }
        return nil
    }
    return res
}
// GetFolderInfo gets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
// returns a V1beta1DatastoresItemDatastoresGetResponse_folderInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetFolderInfo()(V1beta1DatastoresItemDatastoresGetResponse_folderInfoable) {
    return m.folderInfo
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHciClusterUuid gets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetHciClusterUuid()(*string) {
    return m.hciClusterUuid
}
// GetHostsInfo gets the hostsInfo property value. List of hypervisor hosts to which the datastore is presented to.
// returns a []V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetHostsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable) {
    return m.hostsInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetHypervisorManagerInfo()(V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the datastore
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetId()(*string) {
    return m.id
}
// GetMoref gets the moref property value. VMware provided moref of the data store
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetMoref()(*string) {
    return m.moref
}
// GetName gets the name property value. Name of the datastore as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetName()(*string) {
    return m.name
}
// GetProtectionJobInfo gets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
// returns a V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetProtectionJobInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable) {
    return m.protectionJobInfo
}
// GetProtectionPolicyAppliedAtInfo gets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
// returns a V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetProtectionPolicyAppliedAtInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable) {
    return m.protectionPolicyAppliedAtInfo
}
// GetProvisioningPolicyInfo gets the provisioningPolicyInfo property value. Describes provisioning policy information.
// returns a V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetProvisioningPolicyInfo()(V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable) {
    return m.provisioningPolicyInfo
}
// GetRecoveryPointsExist gets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
// returns a *bool when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetRecoveryPointsExist()(*bool) {
    return m.recoveryPointsExist
}
// GetReplicationInfo gets the replicationInfo property value. Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
// returns a V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetReplicationInfo()(V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable) {
    return m.replicationInfo
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the datastore
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetStateReason()(*string) {
    return m.stateReason
}
// GetStorageSystemsInfo gets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
// returns a []V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetStorageSystemsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable) {
    return m.storageSystemsInfo
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. VMware provided uuid of the datastore.
// returns a *string when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVmCount gets the vmCount property value. Number of virtual machines associated with the datastore.
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetVmCount()(*int64) {
    return m.vmCount
}
// GetVmProtectionGroupsInfo gets the vmProtectionGroupsInfo property value. List of virtual machine protection groups.
// returns a []V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetVmProtectionGroupsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable) {
    return m.vmProtectionGroupsInfo
}
// GetVolumesInfo gets the volumesInfo property value. Volumes associated with datastore.
// returns a []V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable when successful
func (m *V1beta1DatastoresItemDatastoresGetResponse) GetVolumesInfo()([]V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable) {
    return m.volumesInfo
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresItemDatastoresGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapacityFree sets the capacityFree property value. Unused storage of the datastore in bytes.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetCapacityFree(value *int64)() {
    m.capacityFree = value
}
// SetCapacityInBytes sets the capacityInBytes property value. Size of the datastore in bytes.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetCapacityInBytes(value *int64)() {
    m.capacityInBytes = value
}
// SetCapacityUncommitted sets the capacityUncommitted property value. Uncommitted storage of the datastore in bytes.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetCapacityUncommitted(value *int64)() {
    m.capacityUncommitted = value
}
// SetClusterInfo sets the clusterInfo property value. The clusterInfo property
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetClusterInfo(value V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable)() {
    m.clusterInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDatacentersInfo sets the datacentersInfo property value. List of datacenters to which the datastore is presented to.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetDatacentersInfo(value []V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable)() {
    m.datacentersInfo = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the datastore.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFolderInfo sets the folderInfo property value. The immediate parent folder on which this resource is hosted in the inventory of hypervisor-manager.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetFolderInfo(value V1beta1DatastoresItemDatastoresGetResponse_folderInfoable)() {
    m.folderInfo = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHciClusterUuid sets the hciClusterUuid property value. UUID string uniquely identifying the HCI cluster.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetHciClusterUuid(value *string)() {
    m.hciClusterUuid = value
}
// SetHostsInfo sets the hostsInfo property value. List of hypervisor hosts to which the datastore is presented to.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetHostsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable)() {
    m.hostsInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetHypervisorManagerInfo(value V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the datastore
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetId(value *string)() {
    m.id = value
}
// SetMoref sets the moref property value. VMware provided moref of the data store
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetMoref(value *string)() {
    m.moref = value
}
// SetName sets the name property value. Name of the datastore as reported by the hypervisor manager.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetName(value *string)() {
    m.name = value
}
// SetProtectionJobInfo sets the protectionJobInfo property value. Information about the assigned Protection Policy and the Protection Job.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetProtectionJobInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable)() {
    m.protectionJobInfo = value
}
// SetProtectionPolicyAppliedAtInfo sets the protectionPolicyAppliedAtInfo property value. Describes applied protection policy information.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetProtectionPolicyAppliedAtInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable)() {
    m.protectionPolicyAppliedAtInfo = value
}
// SetProvisioningPolicyInfo sets the provisioningPolicyInfo property value. Describes provisioning policy information.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetProvisioningPolicyInfo(value V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable)() {
    m.provisioningPolicyInfo = value
}
// SetRecoveryPointsExist sets the recoveryPointsExist property value. Indicates at least one valid recovery point exists for this resource.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetRecoveryPointsExist(value *bool)() {
    m.recoveryPointsExist = value
}
// SetReplicationInfo sets the replicationInfo property value. Replication groups information containing details of all replication partners. Applicable only for Protection Group type 'STORAGE_REPLICATION_GROUP'
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetReplicationInfo(value V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable)() {
    m.replicationInfo = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the datastore
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetStorageSystemsInfo sets the storageSystemsInfo property value. storage systems with volume wwn associated with this virtual machine
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetStorageSystemsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable)() {
    m.storageSystemsInfo = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. VMware provided uuid of the datastore.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVmCount sets the vmCount property value. Number of virtual machines associated with the datastore.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetVmCount(value *int64)() {
    m.vmCount = value
}
// SetVmProtectionGroupsInfo sets the vmProtectionGroupsInfo property value. List of virtual machine protection groups.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetVmProtectionGroupsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable)() {
    m.vmProtectionGroupsInfo = value
}
// SetVolumesInfo sets the volumesInfo property value. Volumes associated with datastore.
func (m *V1beta1DatastoresItemDatastoresGetResponse) SetVolumesInfo(value []V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable)() {
    m.volumesInfo = value
}
type V1beta1DatastoresItemDatastoresGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacityFree()(*int64)
    GetCapacityInBytes()(*int64)
    GetCapacityUncommitted()(*int64)
    GetClusterInfo()(V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDatacentersInfo()([]V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable)
    GetDisplayName()(*string)
    GetFolderInfo()(V1beta1DatastoresItemDatastoresGetResponse_folderInfoable)
    GetGeneration()(*int64)
    GetHciClusterUuid()(*string)
    GetHostsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable)
    GetHypervisorManagerInfo()(V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable)
    GetId()(*string)
    GetMoref()(*string)
    GetName()(*string)
    GetProtectionJobInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable)
    GetProtectionPolicyAppliedAtInfo()(V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable)
    GetProvisioningPolicyInfo()(V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable)
    GetRecoveryPointsExist()(*bool)
    GetReplicationInfo()(V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetStorageSystemsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVmCount()(*int64)
    GetVmProtectionGroupsInfo()([]V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable)
    GetVolumesInfo()([]V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable)
    SetCapacityFree(value *int64)()
    SetCapacityInBytes(value *int64)()
    SetCapacityUncommitted(value *int64)()
    SetClusterInfo(value V1beta1DatastoresItemDatastoresGetResponse_clusterInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDatacentersInfo(value []V1beta1DatastoresItemDatastoresGetResponse_datacentersInfoable)()
    SetDisplayName(value *string)()
    SetFolderInfo(value V1beta1DatastoresItemDatastoresGetResponse_folderInfoable)()
    SetGeneration(value *int64)()
    SetHciClusterUuid(value *string)()
    SetHostsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_hostsInfoable)()
    SetHypervisorManagerInfo(value V1beta1DatastoresItemDatastoresGetResponse_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetMoref(value *string)()
    SetName(value *string)()
    SetProtectionJobInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionJobInfoable)()
    SetProtectionPolicyAppliedAtInfo(value V1beta1DatastoresItemDatastoresGetResponse_protectionPolicyAppliedAtInfoable)()
    SetProvisioningPolicyInfo(value V1beta1DatastoresItemDatastoresGetResponse_provisioningPolicyInfoable)()
    SetRecoveryPointsExist(value *bool)()
    SetReplicationInfo(value V1beta1DatastoresItemDatastoresGetResponse_replicationInfoable)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetStorageSystemsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_storageSystemsInfoable)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVmCount(value *int64)()
    SetVmProtectionGroupsInfo(value []V1beta1DatastoresItemDatastoresGetResponse_vmProtectionGroupsInfoable)()
    SetVolumesInfo(value []V1beta1DatastoresItemDatastoresGetResponse_volumesInfoable)()
}
