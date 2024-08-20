package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemTagsItemTagGetResponse represents a hypervisor tag resource
type V1beta1HypervisorManagersItemTagsItemTagGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of resources to which this tag is associated.
    associatedResources []V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable
    // The category name is unique to the currently selected hypervisor manager.
    categoryName *string
    // Time in UTC at which the object was created.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier.
    customerId *string
    // This will always be same as name since add or update of hypervisor tags are not supported when it is managed from a manager like vCenter.
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable
    // UUID string uniquely identifying the hypervisor tag resource.
    id *string
    // Name of the tag as reported by the hypervisor manager.
    name *string
    // The 'self' reference for this resource.
    resourceUri *string
    // List of services this object belongs to.  This list can be used to filter specific services in the UI.
    services []string
    // The type of resource.
    typeEscaped *string
    // Hypervisor provided identifier of the hypervisor tag.
    uid *string
    // Time in UTC at which the object was last updated.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Protection groups related to the hypervisor-tag.
    vmProtectionGroupsInfo []V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable
}
// NewV1beta1HypervisorManagersItemTagsItemTagGetResponse instantiates a new V1beta1HypervisorManagersItemTagsItemTagGetResponse and sets the default values.
func NewV1beta1HypervisorManagersItemTagsItemTagGetResponse()(*V1beta1HypervisorManagersItemTagsItemTagGetResponse) {
    m := &V1beta1HypervisorManagersItemTagsItemTagGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemTagsItemTagGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemTagsItemTagGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemTagsItemTagGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssociatedResources gets the associatedResources property value. List of resources to which this tag is associated.
// returns a []V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetAssociatedResources()([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable) {
    return m.associatedResources
}
// GetCategoryName gets the categoryName property value. The category name is unique to the currently selected hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetCategoryName()(*string) {
    return m.categoryName
}
// GetCreatedAt gets the createdAt property value. Time in UTC at which the object was created.
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. This will always be same as name since add or update of hypervisor tags are not supported when it is managed from a manager like vCenter.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["associatedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable)
                }
            }
            m.SetAssociatedResources(res)
        }
        return nil
    }
    res["categoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryName(val)
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
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable))
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
    res["vmProtectionGroupsInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable)
                }
            }
            m.SetVmProtectionGroupsInfo(res)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor tag resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the tag as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
// returns a []string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetServices()([]string) {
    return m.services
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. Hypervisor provided identifier of the hypervisor tag.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time in UTC at which the object was last updated.
// returns a *Time when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVmProtectionGroupsInfo gets the vmProtectionGroupsInfo property value. Protection groups related to the hypervisor-tag.
// returns a []V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable when successful
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) GetVmProtectionGroupsInfo()([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable) {
    return m.vmProtectionGroupsInfo
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssociatedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedResources()))
        for i, v := range m.GetAssociatedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("associatedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("categoryName", m.GetCategoryName())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssociatedResources sets the associatedResources property value. List of resources to which this tag is associated.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetAssociatedResources(value []V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable)() {
    m.associatedResources = value
}
// SetCategoryName sets the categoryName property value. The category name is unique to the currently selected hypervisor manager.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetCategoryName(value *string)() {
    m.categoryName = value
}
// SetCreatedAt sets the createdAt property value. Time in UTC at which the object was created.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. This will always be same as name since add or update of hypervisor tags are not supported when it is managed from a manager like vCenter.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor tag resource.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the tag as reported by the hypervisor manager.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. List of services this object belongs to.  This list can be used to filter specific services in the UI.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetServices(value []string)() {
    m.services = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. Hypervisor provided identifier of the hypervisor tag.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time in UTC at which the object was last updated.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVmProtectionGroupsInfo sets the vmProtectionGroupsInfo property value. Protection groups related to the hypervisor-tag.
func (m *V1beta1HypervisorManagersItemTagsItemTagGetResponse) SetVmProtectionGroupsInfo(value []V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable)() {
    m.vmProtectionGroupsInfo = value
}
type V1beta1HypervisorManagersItemTagsItemTagGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedResources()([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable)
    GetCategoryName()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVmProtectionGroupsInfo()([]V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable)
    SetAssociatedResources(value []V1beta1HypervisorManagersItemTagsItemTagGetResponse_associatedResourcesable)()
    SetCategoryName(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemTagsItemTagGetResponse_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVmProtectionGroupsInfo(value []V1beta1HypervisorManagersItemTagsItemTagGetResponse_vmProtectionGroupsInfoable)()
}
