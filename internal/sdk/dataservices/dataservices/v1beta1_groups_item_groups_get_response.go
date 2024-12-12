package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1GroupsItemGroupsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Array of resource type and count used to represent number of resources for a given typeassociated with a group.
    associatedResourceTypes []V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable
    // Count of total resources associated to this group
    associationCount *int32
    // The "consoleUri" property is the webpage representing the resource in the  Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/groups/c4f8f464-ce09-4847-bfaa-e668de838c66")
    consoleUri *string
    // The time the group was created, conforming to the RFC-3339 format.All times are represented in UTC.
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Primary identifier for the customer (GUID) associated with the group.
    customerId *string
    // Text that can be used to shed light on the meaning or purpose of the group. Description supports unicode
    description *string
    // A monotonically increasing value incremented every time the resource is updated
    generation *int32
    // Type of group
    groupType *int32
    // Unique identifer for the Group, required for all other Group operations.The generation of the 'id' will be the responsibility of this service.
    id *string
    // User will always provide a name for the group resource. It supports unicode. It must be unique within each customer ID.
    name *string
    // The "resourceUri" property identifies the resource's URI. The resourceUri value is a URI encoded string starting with the version of the API (e.g. "/data-services/v1beta1/groups/c4f8f464-ce09-4847-bfaa-e668de838c66").
    resourceUri *string
    // The type of the resource. Used for grouping of like resources.
    typeEscaped *string
    // The last modification time, conforming to the RFC-3339 format.All times are represented in UTC.
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user that performed the most recent modifying operation on the Group.
    updatedBy *string
}
// NewV1beta1GroupsItemGroupsGetResponse instantiates a new V1beta1GroupsItemGroupsGetResponse and sets the default values.
func NewV1beta1GroupsItemGroupsGetResponse()(*V1beta1GroupsItemGroupsGetResponse) {
    m := &V1beta1GroupsItemGroupsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1GroupsItemGroupsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsItemGroupsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsItemGroupsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssociatedResourceTypes gets the associatedResourceTypes property value. Array of resource type and count used to represent number of resources for a given typeassociated with a group.
// returns a []V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetAssociatedResourceTypes()([]V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable) {
    return m.associatedResourceTypes
}
// GetAssociationCount gets the associationCount property value. Count of total resources associated to this group
// returns a *int32 when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetAssociationCount()(*int32) {
    return m.associationCount
}
// GetConsoleUri gets the consoleUri property value. The "consoleUri" property is the webpage representing the resource in the  Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/groups/c4f8f464-ce09-4847-bfaa-e668de838c66")
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCreatedAt gets the createdAt property value. The time the group was created, conforming to the RFC-3339 format.All times are represented in UTC.
// returns a *Time when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. Primary identifier for the customer (GUID) associated with the group.
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDescription gets the description property value. Text that can be used to shed light on the meaning or purpose of the group. Description supports unicode
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["associatedResourceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1GroupsItemGroupsGetResponse_associatedResourceTypesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable)
                }
            }
            m.SetAssociatedResourceTypes(res)
        }
        return nil
    }
    res["associationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociationCount(val)
        }
        return nil
    }
    res["consoleUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsoleUri(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["groupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val)
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
    res["updatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedBy(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value incremented every time the resource is updated
// returns a *int32 when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetGeneration()(*int32) {
    return m.generation
}
// GetGroupType gets the groupType property value. Type of group
// returns a *int32 when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetGroupType()(*int32) {
    return m.groupType
}
// GetId gets the id property value. Unique identifer for the Group, required for all other Group operations.The generation of the 'id' will be the responsibility of this service.
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. User will always provide a name for the group resource. It supports unicode. It must be unique within each customer ID.
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The "resourceUri" property identifies the resource's URI. The resourceUri value is a URI encoded string starting with the version of the API (e.g. "/data-services/v1beta1/groups/c4f8f464-ce09-4847-bfaa-e668de838c66").
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of the resource. Used for grouping of like resources.
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The last modification time, conforming to the RFC-3339 format.All times are represented in UTC.
// returns a *Time when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetUpdatedBy gets the updatedBy property value. The user that performed the most recent modifying operation on the Group.
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse) GetUpdatedBy()(*string) {
    return m.updatedBy
}
// Serialize serializes information the current object
func (m *V1beta1GroupsItemGroupsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssociatedResourceTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedResourceTypes()))
        for i, v := range m.GetAssociatedResourceTypes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("associatedResourceTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("associationCount", m.GetAssociationCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("consoleUri", m.GetConsoleUri())
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
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("groupType", m.GetGroupType())
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
        err := writer.WriteStringValue("name", m.GetName())
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
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
        err := writer.WriteStringValue("updatedBy", m.GetUpdatedBy())
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
func (m *V1beta1GroupsItemGroupsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssociatedResourceTypes sets the associatedResourceTypes property value. Array of resource type and count used to represent number of resources for a given typeassociated with a group.
func (m *V1beta1GroupsItemGroupsGetResponse) SetAssociatedResourceTypes(value []V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable)() {
    m.associatedResourceTypes = value
}
// SetAssociationCount sets the associationCount property value. Count of total resources associated to this group
func (m *V1beta1GroupsItemGroupsGetResponse) SetAssociationCount(value *int32)() {
    m.associationCount = value
}
// SetConsoleUri sets the consoleUri property value. The "consoleUri" property is the webpage representing the resource in the  Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/groups/c4f8f464-ce09-4847-bfaa-e668de838c66")
func (m *V1beta1GroupsItemGroupsGetResponse) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCreatedAt sets the createdAt property value. The time the group was created, conforming to the RFC-3339 format.All times are represented in UTC.
func (m *V1beta1GroupsItemGroupsGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. Primary identifier for the customer (GUID) associated with the group.
func (m *V1beta1GroupsItemGroupsGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDescription sets the description property value. Text that can be used to shed light on the meaning or purpose of the group. Description supports unicode
func (m *V1beta1GroupsItemGroupsGetResponse) SetDescription(value *string)() {
    m.description = value
}
// SetGeneration sets the generation property value. A monotonically increasing value incremented every time the resource is updated
func (m *V1beta1GroupsItemGroupsGetResponse) SetGeneration(value *int32)() {
    m.generation = value
}
// SetGroupType sets the groupType property value. Type of group
func (m *V1beta1GroupsItemGroupsGetResponse) SetGroupType(value *int32)() {
    m.groupType = value
}
// SetId sets the id property value. Unique identifer for the Group, required for all other Group operations.The generation of the 'id' will be the responsibility of this service.
func (m *V1beta1GroupsItemGroupsGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. User will always provide a name for the group resource. It supports unicode. It must be unique within each customer ID.
func (m *V1beta1GroupsItemGroupsGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The "resourceUri" property identifies the resource's URI. The resourceUri value is a URI encoded string starting with the version of the API (e.g. "/data-services/v1beta1/groups/c4f8f464-ce09-4847-bfaa-e668de838c66").
func (m *V1beta1GroupsItemGroupsGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of the resource. Used for grouping of like resources.
func (m *V1beta1GroupsItemGroupsGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The last modification time, conforming to the RFC-3339 format.All times are represented in UTC.
func (m *V1beta1GroupsItemGroupsGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetUpdatedBy sets the updatedBy property value. The user that performed the most recent modifying operation on the Group.
func (m *V1beta1GroupsItemGroupsGetResponse) SetUpdatedBy(value *string)() {
    m.updatedBy = value
}
type V1beta1GroupsItemGroupsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedResourceTypes()([]V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable)
    GetAssociationCount()(*int32)
    GetConsoleUri()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDescription()(*string)
    GetGeneration()(*int32)
    GetGroupType()(*int32)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUpdatedBy()(*string)
    SetAssociatedResourceTypes(value []V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable)()
    SetAssociationCount(value *int32)()
    SetConsoleUri(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDescription(value *string)()
    SetGeneration(value *int32)()
    SetGroupType(value *int32)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUpdatedBy(value *string)()
}
