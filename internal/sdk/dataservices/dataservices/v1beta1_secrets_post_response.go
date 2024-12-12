package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsPostResponse secret Resource Definition
type V1beta1SecretsPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of associated assignment resources
    assignmentsCount *int32
    // Resource classifier definition
    classifier V1beta1SecretsPostResponse_classifierable
    // Timestamp of the resource creation
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifier of the customer owning the resource
    customerId *string
    // Resource domain definition
    domain V1beta1SecretsPostResponse_domainable
    // Update generation number
    generation *int32
    // Groups associated with the resource
    groups []V1beta1SecretsPostResponse_groupsable
    // Identifier of the resource
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Consumer-defined label of the resource
    label *string
    // Name of the resource
    name *string
    // Secret properties conformant policy
    policy *string
    // Reference of the prototype resource
    prototype V1beta1SecretsPostResponse_prototypeable
    // URI of the resource
    resourceUri *string
    // Name of the service originating the resource
    service *string
    // Status of the resource
    status *string
    // Timestamp of the last status update
    statusUpdatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Resource subclassifier definition
    subclassifier V1beta1SecretsPostResponse_subclassifierable
    // Type of the resource
    typeEscaped *string
    // Timestamp of the last resource update
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SecretsPostResponse instantiates a new V1beta1SecretsPostResponse and sets the default values.
func NewV1beta1SecretsPostResponse()(*V1beta1SecretsPostResponse) {
    m := &V1beta1SecretsPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignmentsCount gets the assignmentsCount property value. Count of associated assignment resources
// returns a *int32 when successful
func (m *V1beta1SecretsPostResponse) GetAssignmentsCount()(*int32) {
    return m.assignmentsCount
}
// GetClassifier gets the classifier property value. Resource classifier definition
// returns a V1beta1SecretsPostResponse_classifierable when successful
func (m *V1beta1SecretsPostResponse) GetClassifier()(V1beta1SecretsPostResponse_classifierable) {
    return m.classifier
}
// GetCreatedAt gets the createdAt property value. Timestamp of the resource creation
// returns a *Time when successful
func (m *V1beta1SecretsPostResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. Identifier of the customer owning the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDomain gets the domain property value. Resource domain definition
// returns a V1beta1SecretsPostResponse_domainable when successful
func (m *V1beta1SecretsPostResponse) GetDomain()(V1beta1SecretsPostResponse_domainable) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentsCount(val)
        }
        return nil
    }
    res["classifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostResponse_classifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassifier(val.(V1beta1SecretsPostResponse_classifierable))
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
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostResponse_domainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val.(V1beta1SecretsPostResponse_domainable))
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SecretsPostResponse_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SecretsPostResponse_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SecretsPostResponse_groupsable)
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
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val)
        }
        return nil
    }
    res["prototype"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostResponse_prototypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrototype(val.(V1beta1SecretsPostResponse_prototypeable))
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
    res["subclassifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SecretsPostResponse_subclassifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubclassifier(val.(V1beta1SecretsPostResponse_subclassifierable))
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
func (m *V1beta1SecretsPostResponse) GetGeneration()(*int32) {
    return m.generation
}
// GetGroups gets the groups property value. Groups associated with the resource
// returns a []V1beta1SecretsPostResponse_groupsable when successful
func (m *V1beta1SecretsPostResponse) GetGroups()([]V1beta1SecretsPostResponse_groupsable) {
    return m.groups
}
// GetId gets the id property value. Identifier of the resource
// returns a *UUID when successful
func (m *V1beta1SecretsPostResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetLabel gets the label property value. Consumer-defined label of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetLabel()(*string) {
    return m.label
}
// GetName gets the name property value. Name of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetName()(*string) {
    return m.name
}
// GetPolicy gets the policy property value. Secret properties conformant policy
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetPolicy()(*string) {
    return m.policy
}
// GetPrototype gets the prototype property value. Reference of the prototype resource
// returns a V1beta1SecretsPostResponse_prototypeable when successful
func (m *V1beta1SecretsPostResponse) GetPrototype()(V1beta1SecretsPostResponse_prototypeable) {
    return m.prototype
}
// GetResourceUri gets the resourceUri property value. URI of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetService gets the service property value. Name of the service originating the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetService()(*string) {
    return m.service
}
// GetStatus gets the status property value. Status of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetStatus()(*string) {
    return m.status
}
// GetStatusUpdatedAt gets the statusUpdatedAt property value. Timestamp of the last status update
// returns a *Time when successful
func (m *V1beta1SecretsPostResponse) GetStatusUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.statusUpdatedAt
}
// GetSubclassifier gets the subclassifier property value. Resource subclassifier definition
// returns a V1beta1SecretsPostResponse_subclassifierable when successful
func (m *V1beta1SecretsPostResponse) GetSubclassifier()(V1beta1SecretsPostResponse_subclassifierable) {
    return m.subclassifier
}
// GetTypeEscaped gets the type property value. Type of the resource
// returns a *string when successful
func (m *V1beta1SecretsPostResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. Timestamp of the last resource update
// returns a *Time when successful
func (m *V1beta1SecretsPostResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SecretsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignmentsCount", m.GetAssignmentsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("classifier", m.GetClassifier())
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
        err := writer.WriteObjectValue("domain", m.GetDomain())
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
        err := writer.WriteStringValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("prototype", m.GetPrototype())
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
        err := writer.WriteObjectValue("subclassifier", m.GetSubclassifier())
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
func (m *V1beta1SecretsPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignmentsCount sets the assignmentsCount property value. Count of associated assignment resources
func (m *V1beta1SecretsPostResponse) SetAssignmentsCount(value *int32)() {
    m.assignmentsCount = value
}
// SetClassifier sets the classifier property value. Resource classifier definition
func (m *V1beta1SecretsPostResponse) SetClassifier(value V1beta1SecretsPostResponse_classifierable)() {
    m.classifier = value
}
// SetCreatedAt sets the createdAt property value. Timestamp of the resource creation
func (m *V1beta1SecretsPostResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. Identifier of the customer owning the resource
func (m *V1beta1SecretsPostResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDomain sets the domain property value. Resource domain definition
func (m *V1beta1SecretsPostResponse) SetDomain(value V1beta1SecretsPostResponse_domainable)() {
    m.domain = value
}
// SetGeneration sets the generation property value. Update generation number
func (m *V1beta1SecretsPostResponse) SetGeneration(value *int32)() {
    m.generation = value
}
// SetGroups sets the groups property value. Groups associated with the resource
func (m *V1beta1SecretsPostResponse) SetGroups(value []V1beta1SecretsPostResponse_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. Identifier of the resource
func (m *V1beta1SecretsPostResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetLabel sets the label property value. Consumer-defined label of the resource
func (m *V1beta1SecretsPostResponse) SetLabel(value *string)() {
    m.label = value
}
// SetName sets the name property value. Name of the resource
func (m *V1beta1SecretsPostResponse) SetName(value *string)() {
    m.name = value
}
// SetPolicy sets the policy property value. Secret properties conformant policy
func (m *V1beta1SecretsPostResponse) SetPolicy(value *string)() {
    m.policy = value
}
// SetPrototype sets the prototype property value. Reference of the prototype resource
func (m *V1beta1SecretsPostResponse) SetPrototype(value V1beta1SecretsPostResponse_prototypeable)() {
    m.prototype = value
}
// SetResourceUri sets the resourceUri property value. URI of the resource
func (m *V1beta1SecretsPostResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetService sets the service property value. Name of the service originating the resource
func (m *V1beta1SecretsPostResponse) SetService(value *string)() {
    m.service = value
}
// SetStatus sets the status property value. Status of the resource
func (m *V1beta1SecretsPostResponse) SetStatus(value *string)() {
    m.status = value
}
// SetStatusUpdatedAt sets the statusUpdatedAt property value. Timestamp of the last status update
func (m *V1beta1SecretsPostResponse) SetStatusUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.statusUpdatedAt = value
}
// SetSubclassifier sets the subclassifier property value. Resource subclassifier definition
func (m *V1beta1SecretsPostResponse) SetSubclassifier(value V1beta1SecretsPostResponse_subclassifierable)() {
    m.subclassifier = value
}
// SetTypeEscaped sets the type property value. Type of the resource
func (m *V1beta1SecretsPostResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. Timestamp of the last resource update
func (m *V1beta1SecretsPostResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SecretsPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentsCount()(*int32)
    GetClassifier()(V1beta1SecretsPostResponse_classifierable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDomain()(V1beta1SecretsPostResponse_domainable)
    GetGeneration()(*int32)
    GetGroups()([]V1beta1SecretsPostResponse_groupsable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLabel()(*string)
    GetName()(*string)
    GetPolicy()(*string)
    GetPrototype()(V1beta1SecretsPostResponse_prototypeable)
    GetResourceUri()(*string)
    GetService()(*string)
    GetStatus()(*string)
    GetStatusUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubclassifier()(V1beta1SecretsPostResponse_subclassifierable)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssignmentsCount(value *int32)()
    SetClassifier(value V1beta1SecretsPostResponse_classifierable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDomain(value V1beta1SecretsPostResponse_domainable)()
    SetGeneration(value *int32)()
    SetGroups(value []V1beta1SecretsPostResponse_groupsable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLabel(value *string)()
    SetName(value *string)()
    SetPolicy(value *string)()
    SetPrototype(value V1beta1SecretsPostResponse_prototypeable)()
    SetResourceUri(value *string)()
    SetService(value *string)()
    SetStatus(value *string)()
    SetStatusUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubclassifier(value V1beta1SecretsPostResponse_subclassifierable)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
