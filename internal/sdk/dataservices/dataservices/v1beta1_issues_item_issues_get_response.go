package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1IssuesItemIssuesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Long description with more details including possible remediations.
    body *string
    // Category of the issue. PERFORMANCE, CAPACITY, etc
    category *string
    // Time when the issue was cleared. RFC 3339 timestamp
    clearedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Groups this issue is associated with.
    groups []V1beta1IssuesItemIssuesGetResponse_groupsable
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The type of the issue. Eg: ISSUE
    issueType *string
    // Time when the issue last occurred. RFC 3339 timestamp
    lastOccurredAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A system specified name for the resource.
    name *string
    // Indicates the number of occurrences of the issue
    occurrenceCount *int32
    // Recommendation to address the underlying issue.
    recommendation *string
    // Details of the resources related to the issue
    relatedResources []V1beta1IssuesItemIssuesGetResponse_relatedResourcesable
    // The self reference for this resource.
    resourceUri *string
    // Details of the services this issue belongs to
    services []string
    // Severity of the issue. For issue: CRITICAL, WARNING, INFO. For reco: HIGH, MEDIUM, LOW
    severity *string
    // An auxiliary calculated attribute to help the end-user filter snoozed and un-snoozed issues
    snoozed *bool
    // The email id of the last user who snoozed this issue.
    snoozedBy *string
    // Date-time until this issue will be considered snoozed/inactive until this time. RFC 3339
    snoozedUntil *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifier for the source resource that is contained in relatedResources
    sourceResourceId *string
    // Type of the source resource that is contained in the relatedResources
    sourceResourceType *string
    // State of the issue. Eg: CREATED, ASSIGNED, CLOSED, SNOOZED, DELETED, etc
    state *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1IssuesItemIssuesGetResponse instantiates a new V1beta1IssuesItemIssuesGetResponse and sets the default values.
func NewV1beta1IssuesItemIssuesGetResponse()(*V1beta1IssuesItemIssuesGetResponse) {
    m := &V1beta1IssuesItemIssuesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1IssuesItemIssuesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1IssuesItemIssuesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1IssuesItemIssuesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Long description with more details including possible remediations.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetBody()(*string) {
    return m.body
}
// GetCategory gets the category property value. Category of the issue. PERFORMANCE, CAPACITY, etc
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetCategory()(*string) {
    return m.category
}
// GetClearedAt gets the clearedAt property value. Time when the issue was cleared. RFC 3339 timestamp
// returns a *Time when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetClearedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.clearedAt
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["clearedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClearedAt(val)
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1IssuesItemIssuesGetResponse_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1IssuesItemIssuesGetResponse_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1IssuesItemIssuesGetResponse_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["issueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueType(val)
        }
        return nil
    }
    res["lastOccurredAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastOccurredAt(val)
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
    res["occurrenceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurrenceCount(val)
        }
        return nil
    }
    res["recommendation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendation(val)
        }
        return nil
    }
    res["relatedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1IssuesItemIssuesGetResponse_relatedResourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1IssuesItemIssuesGetResponse_relatedResourcesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1IssuesItemIssuesGetResponse_relatedResourcesable)
                }
            }
            m.SetRelatedResources(res)
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
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val)
        }
        return nil
    }
    res["snoozed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozed(val)
        }
        return nil
    }
    res["snoozedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozedBy(val)
        }
        return nil
    }
    res["snoozedUntil"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozedUntil(val)
        }
        return nil
    }
    res["sourceResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceResourceId(val)
        }
        return nil
    }
    res["sourceResourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceResourceType(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
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
func (m *V1beta1IssuesItemIssuesGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetGroups gets the groups property value. Groups this issue is associated with.
// returns a []V1beta1IssuesItemIssuesGetResponse_groupsable when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetGroups()([]V1beta1IssuesItemIssuesGetResponse_groupsable) {
    return m.groups
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetIssueType gets the issueType property value. The type of the issue. Eg: ISSUE
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetIssueType()(*string) {
    return m.issueType
}
// GetLastOccurredAt gets the lastOccurredAt property value. Time when the issue last occurred. RFC 3339 timestamp
// returns a *Time when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetLastOccurredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastOccurredAt
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetName()(*string) {
    return m.name
}
// GetOccurrenceCount gets the occurrenceCount property value. Indicates the number of occurrences of the issue
// returns a *int32 when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetOccurrenceCount()(*int32) {
    return m.occurrenceCount
}
// GetRecommendation gets the recommendation property value. Recommendation to address the underlying issue.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetRecommendation()(*string) {
    return m.recommendation
}
// GetRelatedResources gets the relatedResources property value. Details of the resources related to the issue
// returns a []V1beta1IssuesItemIssuesGetResponse_relatedResourcesable when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetRelatedResources()([]V1beta1IssuesItemIssuesGetResponse_relatedResourcesable) {
    return m.relatedResources
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetServices gets the services property value. Details of the services this issue belongs to
// returns a []string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetServices()([]string) {
    return m.services
}
// GetSeverity gets the severity property value. Severity of the issue. For issue: CRITICAL, WARNING, INFO. For reco: HIGH, MEDIUM, LOW
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSeverity()(*string) {
    return m.severity
}
// GetSnoozed gets the snoozed property value. An auxiliary calculated attribute to help the end-user filter snoozed and un-snoozed issues
// returns a *bool when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSnoozed()(*bool) {
    return m.snoozed
}
// GetSnoozedBy gets the snoozedBy property value. The email id of the last user who snoozed this issue.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSnoozedBy()(*string) {
    return m.snoozedBy
}
// GetSnoozedUntil gets the snoozedUntil property value. Date-time until this issue will be considered snoozed/inactive until this time. RFC 3339
// returns a *Time when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSnoozedUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.snoozedUntil
}
// GetSourceResourceId gets the sourceResourceId property value. Identifier for the source resource that is contained in relatedResources
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSourceResourceId()(*string) {
    return m.sourceResourceId
}
// GetSourceResourceType gets the sourceResourceType property value. Type of the source resource that is contained in the relatedResources
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetSourceResourceType()(*string) {
    return m.sourceResourceType
}
// GetState gets the state property value. State of the issue. Eg: CREATED, ASSIGNED, CLOSED, SNOOZED, DELETED, etc
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetState()(*string) {
    return m.state
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1IssuesItemIssuesGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1IssuesItemIssuesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("clearedAt", m.GetClearedAt())
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
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issueType", m.GetIssueType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastOccurredAt", m.GetLastOccurredAt())
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
        err := writer.WriteInt32Value("occurrenceCount", m.GetOccurrenceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recommendation", m.GetRecommendation())
        if err != nil {
            return err
        }
    }
    if m.GetRelatedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelatedResources()))
        for i, v := range m.GetRelatedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("relatedResources", cast)
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
        err := writer.WriteStringValue("severity", m.GetSeverity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("snoozed", m.GetSnoozed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("snoozedBy", m.GetSnoozedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("snoozedUntil", m.GetSnoozedUntil())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceResourceId", m.GetSourceResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceResourceType", m.GetSourceResourceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *V1beta1IssuesItemIssuesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Long description with more details including possible remediations.
func (m *V1beta1IssuesItemIssuesGetResponse) SetBody(value *string)() {
    m.body = value
}
// SetCategory sets the category property value. Category of the issue. PERFORMANCE, CAPACITY, etc
func (m *V1beta1IssuesItemIssuesGetResponse) SetCategory(value *string)() {
    m.category = value
}
// SetClearedAt sets the clearedAt property value. Time when the issue was cleared. RFC 3339 timestamp
func (m *V1beta1IssuesItemIssuesGetResponse) SetClearedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.clearedAt = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1IssuesItemIssuesGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1IssuesItemIssuesGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1IssuesItemIssuesGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGroups sets the groups property value. Groups this issue is associated with.
func (m *V1beta1IssuesItemIssuesGetResponse) SetGroups(value []V1beta1IssuesItemIssuesGetResponse_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1IssuesItemIssuesGetResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetIssueType sets the issueType property value. The type of the issue. Eg: ISSUE
func (m *V1beta1IssuesItemIssuesGetResponse) SetIssueType(value *string)() {
    m.issueType = value
}
// SetLastOccurredAt sets the lastOccurredAt property value. Time when the issue last occurred. RFC 3339 timestamp
func (m *V1beta1IssuesItemIssuesGetResponse) SetLastOccurredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastOccurredAt = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1IssuesItemIssuesGetResponse) SetName(value *string)() {
    m.name = value
}
// SetOccurrenceCount sets the occurrenceCount property value. Indicates the number of occurrences of the issue
func (m *V1beta1IssuesItemIssuesGetResponse) SetOccurrenceCount(value *int32)() {
    m.occurrenceCount = value
}
// SetRecommendation sets the recommendation property value. Recommendation to address the underlying issue.
func (m *V1beta1IssuesItemIssuesGetResponse) SetRecommendation(value *string)() {
    m.recommendation = value
}
// SetRelatedResources sets the relatedResources property value. Details of the resources related to the issue
func (m *V1beta1IssuesItemIssuesGetResponse) SetRelatedResources(value []V1beta1IssuesItemIssuesGetResponse_relatedResourcesable)() {
    m.relatedResources = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1IssuesItemIssuesGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetServices sets the services property value. Details of the services this issue belongs to
func (m *V1beta1IssuesItemIssuesGetResponse) SetServices(value []string)() {
    m.services = value
}
// SetSeverity sets the severity property value. Severity of the issue. For issue: CRITICAL, WARNING, INFO. For reco: HIGH, MEDIUM, LOW
func (m *V1beta1IssuesItemIssuesGetResponse) SetSeverity(value *string)() {
    m.severity = value
}
// SetSnoozed sets the snoozed property value. An auxiliary calculated attribute to help the end-user filter snoozed and un-snoozed issues
func (m *V1beta1IssuesItemIssuesGetResponse) SetSnoozed(value *bool)() {
    m.snoozed = value
}
// SetSnoozedBy sets the snoozedBy property value. The email id of the last user who snoozed this issue.
func (m *V1beta1IssuesItemIssuesGetResponse) SetSnoozedBy(value *string)() {
    m.snoozedBy = value
}
// SetSnoozedUntil sets the snoozedUntil property value. Date-time until this issue will be considered snoozed/inactive until this time. RFC 3339
func (m *V1beta1IssuesItemIssuesGetResponse) SetSnoozedUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.snoozedUntil = value
}
// SetSourceResourceId sets the sourceResourceId property value. Identifier for the source resource that is contained in relatedResources
func (m *V1beta1IssuesItemIssuesGetResponse) SetSourceResourceId(value *string)() {
    m.sourceResourceId = value
}
// SetSourceResourceType sets the sourceResourceType property value. Type of the source resource that is contained in the relatedResources
func (m *V1beta1IssuesItemIssuesGetResponse) SetSourceResourceType(value *string)() {
    m.sourceResourceType = value
}
// SetState sets the state property value. State of the issue. Eg: CREATED, ASSIGNED, CLOSED, SNOOZED, DELETED, etc
func (m *V1beta1IssuesItemIssuesGetResponse) SetState(value *string)() {
    m.state = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1IssuesItemIssuesGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1IssuesItemIssuesGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1IssuesItemIssuesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetCategory()(*string)
    GetClearedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetGroups()([]V1beta1IssuesItemIssuesGetResponse_groupsable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIssueType()(*string)
    GetLastOccurredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetOccurrenceCount()(*int32)
    GetRecommendation()(*string)
    GetRelatedResources()([]V1beta1IssuesItemIssuesGetResponse_relatedResourcesable)
    GetResourceUri()(*string)
    GetServices()([]string)
    GetSeverity()(*string)
    GetSnoozed()(*bool)
    GetSnoozedBy()(*string)
    GetSnoozedUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSourceResourceId()(*string)
    GetSourceResourceType()(*string)
    GetState()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBody(value *string)()
    SetCategory(value *string)()
    SetClearedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetGroups(value []V1beta1IssuesItemIssuesGetResponse_groupsable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIssueType(value *string)()
    SetLastOccurredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetOccurrenceCount(value *int32)()
    SetRecommendation(value *string)()
    SetRelatedResources(value []V1beta1IssuesItemIssuesGetResponse_relatedResourcesable)()
    SetResourceUri(value *string)()
    SetServices(value []string)()
    SetSeverity(value *string)()
    SetSnoozed(value *bool)()
    SetSnoozedBy(value *string)()
    SetSnoozedUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSourceResourceId(value *string)()
    SetSourceResourceType(value *string)()
    SetState(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
