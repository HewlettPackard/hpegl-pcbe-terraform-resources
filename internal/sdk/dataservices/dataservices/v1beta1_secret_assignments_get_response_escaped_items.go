package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsGetResponse_items assignment Resource Definition
type V1beta1SecretAssignmentsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Appliance identifier
    appliance V1beta1SecretAssignmentsGetResponse_items_applianceable
    // Timestamp of the resource creation
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifier of the customer owning the resource
    customerId *string
    // Update generation number
    generation *int32
    // Current goal of the assignment
    goal *string
    // Timestamp of the last goal update
    goalUpdatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Groups associated with the resource
    groups []V1beta1SecretAssignmentsGetResponse_items_groupsable
    // Identifier of the resource
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Consumer-defined label of the assignment
    label *string
    // Name of the resource
    name *string
    // URI of the resource
    resourceUri *string
    // Secret resource reference
    secret V1beta1SecretAssignmentsGetResponse_items_secretable
    // Name of the service originating the resource
    service *string
    // Current status of the assignment
    status *string
    // Timestamp of the last status update
    statusUpdatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Type of the resource
    typeEscaped *string
    // Timestamp of the last resource update
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SecretAssignmentsGetResponse_items instantiates a new V1beta1SecretAssignmentsGetResponse_items and sets the default values.
func NewV1beta1SecretAssignmentsGetResponse_items()(*V1beta1SecretAssignmentsGetResponse_items) {
    m := &V1beta1SecretAssignmentsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretAssignmentsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppliance gets the appliance property value. Appliance identifier
// returns a V1beta1SecretAssignmentsGetResponse_items_applianceable when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetAppliance()(V1beta1SecretAssignmentsGetResponse_items_applianceable) {
    return m.appliance
}
// GetCreatedAt gets the createdAt property value. Timestamp of the resource creation
// returns a *Time when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. Identifier of the customer owning the resource
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretAssignmentsGetResponse_items_applianceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliance(val.(V1beta1SecretAssignmentsGetResponse_items_applianceable))
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["goal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoal(val)
        }
        return nil
    }
    res["goalUpdatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoalUpdatedAt(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SecretAssignmentsGetResponse_items_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SecretAssignmentsGetResponse_items_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SecretAssignmentsGetResponse_items_groupsable)
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
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretAssignmentsGetResponse_items_secretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val.(V1beta1SecretAssignmentsGetResponse_items_secretable))
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["statusUpdatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusUpdatedAt(val)
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
// GetGeneration gets the generation property value. Update generation number
// returns a *int32 when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetGeneration()(*int32) {
    return m.generation
}
// GetGoal gets the goal property value. Current goal of the assignment
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetGoal()(*string) {
    return m.goal
}
// GetGoalUpdatedAt gets the goalUpdatedAt property value. Timestamp of the last goal update
// returns a *Time when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetGoalUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.goalUpdatedAt
}
// GetGroups gets the groups property value. Groups associated with the resource
// returns a []V1beta1SecretAssignmentsGetResponse_items_groupsable when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetGroups()([]V1beta1SecretAssignmentsGetResponse_items_groupsable) {
    return m.groups
}
// GetId gets the id property value. Identifier of the resource
// returns a *UUID when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetLabel gets the label property value. Consumer-defined label of the assignment
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetLabel()(*string) {
    return m.label
}
// GetName gets the name property value. Name of the resource
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. URI of the resource
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSecret gets the secret property value. Secret resource reference
// returns a V1beta1SecretAssignmentsGetResponse_items_secretable when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetSecret()(V1beta1SecretAssignmentsGetResponse_items_secretable) {
    return m.secret
}
// GetService gets the service property value. Name of the service originating the resource
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetService()(*string) {
    return m.service
}
// GetStatus gets the status property value. Current status of the assignment
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetStatus()(*string) {
    return m.status
}
// GetStatusUpdatedAt gets the statusUpdatedAt property value. Timestamp of the last status update
// returns a *Time when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetStatusUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.statusUpdatedAt
}
// GetTypeEscaped gets the type property value. Type of the resource
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. Timestamp of the last resource update
// returns a *Time when successful
func (m *V1beta1SecretAssignmentsGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SecretAssignmentsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appliance", m.GetAppliance())
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
        err := writer.WriteInt32Value("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("goal", m.GetGoal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("goalUpdatedAt", m.GetGoalUpdatedAt())
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
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
        err := writer.WriteObjectValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("statusUpdatedAt", m.GetStatusUpdatedAt())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SecretAssignmentsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppliance sets the appliance property value. Appliance identifier
func (m *V1beta1SecretAssignmentsGetResponse_items) SetAppliance(value V1beta1SecretAssignmentsGetResponse_items_applianceable)() {
    m.appliance = value
}
// SetCreatedAt sets the createdAt property value. Timestamp of the resource creation
func (m *V1beta1SecretAssignmentsGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. Identifier of the customer owning the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. Update generation number
func (m *V1beta1SecretAssignmentsGetResponse_items) SetGeneration(value *int32)() {
    m.generation = value
}
// SetGoal sets the goal property value. Current goal of the assignment
func (m *V1beta1SecretAssignmentsGetResponse_items) SetGoal(value *string)() {
    m.goal = value
}
// SetGoalUpdatedAt sets the goalUpdatedAt property value. Timestamp of the last goal update
func (m *V1beta1SecretAssignmentsGetResponse_items) SetGoalUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.goalUpdatedAt = value
}
// SetGroups sets the groups property value. Groups associated with the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetGroups(value []V1beta1SecretAssignmentsGetResponse_items_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. Identifier of the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetLabel sets the label property value. Consumer-defined label of the assignment
func (m *V1beta1SecretAssignmentsGetResponse_items) SetLabel(value *string)() {
    m.label = value
}
// SetName sets the name property value. Name of the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. URI of the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSecret sets the secret property value. Secret resource reference
func (m *V1beta1SecretAssignmentsGetResponse_items) SetSecret(value V1beta1SecretAssignmentsGetResponse_items_secretable)() {
    m.secret = value
}
// SetService sets the service property value. Name of the service originating the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetService(value *string)() {
    m.service = value
}
// SetStatus sets the status property value. Current status of the assignment
func (m *V1beta1SecretAssignmentsGetResponse_items) SetStatus(value *string)() {
    m.status = value
}
// SetStatusUpdatedAt sets the statusUpdatedAt property value. Timestamp of the last status update
func (m *V1beta1SecretAssignmentsGetResponse_items) SetStatusUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.statusUpdatedAt = value
}
// SetTypeEscaped sets the type property value. Type of the resource
func (m *V1beta1SecretAssignmentsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. Timestamp of the last resource update
func (m *V1beta1SecretAssignmentsGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SecretAssignmentsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliance()(V1beta1SecretAssignmentsGetResponse_items_applianceable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetGeneration()(*int32)
    GetGoal()(*string)
    GetGoalUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroups()([]V1beta1SecretAssignmentsGetResponse_items_groupsable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLabel()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetSecret()(V1beta1SecretAssignmentsGetResponse_items_secretable)
    GetService()(*string)
    GetStatus()(*string)
    GetStatusUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAppliance(value V1beta1SecretAssignmentsGetResponse_items_applianceable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetGeneration(value *int32)()
    SetGoal(value *string)()
    SetGoalUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroups(value []V1beta1SecretAssignmentsGetResponse_items_groupsable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLabel(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetSecret(value V1beta1SecretAssignmentsGetResponse_items_secretable)()
    SetService(value *string)()
    SetStatus(value *string)()
    SetStatusUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
