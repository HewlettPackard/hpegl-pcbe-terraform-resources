package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemResourcePoolsGetResponse_items represents a hypervisor resource pool.
type V1beta1HypervisorManagersItemResourcePoolsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Application specific information for this resource pool.
    appInfo V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable
    // Captures the list of all child resource pools.
    childResourcePools []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable
    // Hypervisor cluster information.
    clusterInfo V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // A user-friendly name that identifies the hypervisor resource pool. This will always be same as name since adding or updating hypervisor resource pools is not supported when managed from a manager, such as the vCenter.
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // List of associated hypervisor hosts.
    hostsInfo []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable
    // UUID string uniquely identifying the hypervisor resource pool resource.
    id *string
    // Name of the resource pool as reported by the hypervisor manager.
    name *string
    // Owner of this resource pool.
    ownerInfo V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // Brief reason for the current state of the hypervisor resource pool.
    stateReason *string
    // The type of resource.
    typeEscaped *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items instantiates a new V1beta1HypervisorManagersItemResourcePoolsGetResponse_items and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items()(*V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) {
    m := &V1beta1HypervisorManagersItemResourcePoolsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemResourcePoolsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppInfo gets the appInfo property value. Application specific information for this resource pool.
// returns a V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetAppInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable) {
    return m.appInfo
}
// GetChildResourcePools gets the childResourcePools property value. Captures the list of all child resource pools.
// returns a []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetChildResourcePools()([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable) {
    return m.childResourcePools
}
// GetClusterInfo gets the clusterInfo property value. Hypervisor cluster information.
// returns a V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetClusterInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable) {
    return m.clusterInfo
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the hypervisor resource pool. This will always be same as name since adding or updating hypervisor resource pools is not supported when managed from a manager, such as the vCenter.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppInfo(val.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable))
        }
        return nil
    }
    res["childResourcePools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable)
                }
            }
            m.SetChildResourcePools(res)
        }
        return nil
    }
    res["clusterInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterInfo(val.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable))
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
    res["hostsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable)
                }
            }
            m.SetHostsInfo(res)
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable))
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
    res["ownerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerInfo(val.(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable))
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
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetHostsInfo gets the hostsInfo property value. List of associated hypervisor hosts.
// returns a []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetHostsInfo()([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable) {
    return m.hostsInfo
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor resource pool resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the resource pool as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetName()(*string) {
    return m.name
}
// GetOwnerInfo gets the ownerInfo property value. Owner of this resource pool.
// returns a V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetOwnerInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable) {
    return m.ownerInfo
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetStateReason gets the stateReason property value. Brief reason for the current state of the hypervisor resource pool.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetStateReason()(*string) {
    return m.stateReason
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appInfo", m.GetAppInfo())
        if err != nil {
            return err
        }
    }
    if m.GetChildResourcePools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildResourcePools()))
        for i, v := range m.GetChildResourcePools() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("childResourcePools", cast)
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
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteObjectValue("ownerInfo", m.GetOwnerInfo())
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
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppInfo sets the appInfo property value. Application specific information for this resource pool.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetAppInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable)() {
    m.appInfo = value
}
// SetChildResourcePools sets the childResourcePools property value. Captures the list of all child resource pools.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetChildResourcePools(value []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable)() {
    m.childResourcePools = value
}
// SetClusterInfo sets the clusterInfo property value. Hypervisor cluster information.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetClusterInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable)() {
    m.clusterInfo = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the hypervisor resource pool. This will always be same as name since adding or updating hypervisor resource pools is not supported when managed from a manager, such as the vCenter.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHostsInfo sets the hostsInfo property value. List of associated hypervisor hosts.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetHostsInfo(value []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable)() {
    m.hostsInfo = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor resource pool resource.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the resource pool as reported by the hypervisor manager.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetOwnerInfo sets the ownerInfo property value. Owner of this resource pool.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetOwnerInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable)() {
    m.ownerInfo = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetStateReason sets the stateReason property value. Brief reason for the current state of the hypervisor resource pool.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetStateReason(value *string)() {
    m.stateReason = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1HypervisorManagersItemResourcePoolsGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1HypervisorManagersItemResourcePoolsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable)
    GetChildResourcePools()([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable)
    GetClusterInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetHostsInfo()([]V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable)
    GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable)
    GetId()(*string)
    GetName()(*string)
    GetOwnerInfo()(V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetStateReason()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAppInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_appInfoable)()
    SetChildResourcePools(value []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_childResourcePoolsable)()
    SetClusterInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_clusterInfoable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetHostsInfo(value []V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hostsInfoable)()
    SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetName(value *string)()
    SetOwnerInfo(value V1beta1HypervisorManagersItemResourcePoolsGetResponse_items_ownerInfoable)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetStateReason(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
