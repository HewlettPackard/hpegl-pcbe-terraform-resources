package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DualAuthOperationsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Resources associated with this operation
    associatedResources []V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable
    // Time when this operation was checked. RFC 3339 Timestamp
    checkedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email of the user who checked (second authorization) this operation
    checkedByEmail *string
    // URI of the user who checked (second authorization) this operation
    checkedByUri *string
    // The URI for console screen that displays this resource. Deprecated - use associatedResources instead
    // Deprecated: 
    consoleUri *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // Detailed description of the operation
    description *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // Groups this operation is associated with. Deprecated - use associatedResources instead
    // Deprecated: 
    groups []V1beta1DualAuthOperationsGetResponse_items_groupsable
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // A system specified name for the resource.
    name *string
    // Operation resource on which the operation is taking place. Deprecated - use associatedResources instead
    // Deprecated: 
    operationResource V1beta1DualAuthOperationsGetResponse_items_operationResourceable
    // Time when this operation was requested. RFC 3339 Timestamp
    requestedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email of the user who performed this operation
    requestedByEmail *string
    // URI of the user who performed this operation
    requestedByUri *string
    // One word description of the operation
    requestedOperation *string
    // The self reference for this resource.
    resourceUri *string
    // External Service Name from where this request was sent
    sourceServiceExternalName *string
    // state of this operation
    state *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1DualAuthOperationsGetResponse_items instantiates a new V1beta1DualAuthOperationsGetResponse_items and sets the default values.
func NewV1beta1DualAuthOperationsGetResponse_items()(*V1beta1DualAuthOperationsGetResponse_items) {
    m := &V1beta1DualAuthOperationsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DualAuthOperationsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DualAuthOperationsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DualAuthOperationsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssociatedResources gets the associatedResources property value. Resources associated with this operation
// returns a []V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetAssociatedResources()([]V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable) {
    return m.associatedResources
}
// GetCheckedAt gets the checkedAt property value. Time when this operation was checked. RFC 3339 Timestamp
// returns a *Time when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetCheckedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.checkedAt
}
// GetCheckedByEmail gets the checkedByEmail property value. Email of the user who checked (second authorization) this operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetCheckedByEmail()(*string) {
    return m.checkedByEmail
}
// GetCheckedByUri gets the checkedByUri property value. URI of the user who checked (second authorization) this operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetCheckedByUri()(*string) {
    return m.checkedByUri
}
// GetConsoleUri gets the consoleUri property value. The URI for console screen that displays this resource. Deprecated - use associatedResources instead
// Deprecated: 
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDescription gets the description property value. Detailed description of the operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["associatedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DualAuthOperationsGetResponse_items_associatedResourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable)
                }
            }
            m.SetAssociatedResources(res)
        }
        return nil
    }
    res["checkedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedAt(val)
        }
        return nil
    }
    res["checkedByEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedByEmail(val)
        }
        return nil
    }
    res["checkedByUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedByUri(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DualAuthOperationsGetResponse_items_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DualAuthOperationsGetResponse_items_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DualAuthOperationsGetResponse_items_groupsable)
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
    res["operationResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DualAuthOperationsGetResponse_items_operationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationResource(val.(V1beta1DualAuthOperationsGetResponse_items_operationResourceable))
        }
        return nil
    }
    res["requestedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedAt(val)
        }
        return nil
    }
    res["requestedByEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedByEmail(val)
        }
        return nil
    }
    res["requestedByUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedByUri(val)
        }
        return nil
    }
    res["requestedOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedOperation(val)
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
    res["sourceServiceExternalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceServiceExternalName(val)
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
func (m *V1beta1DualAuthOperationsGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetGroups gets the groups property value. Groups this operation is associated with. Deprecated - use associatedResources instead
// Deprecated: 
// returns a []V1beta1DualAuthOperationsGetResponse_items_groupsable when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetGroups()([]V1beta1DualAuthOperationsGetResponse_items_groupsable) {
    return m.groups
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetName()(*string) {
    return m.name
}
// GetOperationResource gets the operationResource property value. Operation resource on which the operation is taking place. Deprecated - use associatedResources instead
// Deprecated: 
// returns a V1beta1DualAuthOperationsGetResponse_items_operationResourceable when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetOperationResource()(V1beta1DualAuthOperationsGetResponse_items_operationResourceable) {
    return m.operationResource
}
// GetRequestedAt gets the requestedAt property value. Time when this operation was requested. RFC 3339 Timestamp
// returns a *Time when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetRequestedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestedAt
}
// GetRequestedByEmail gets the requestedByEmail property value. Email of the user who performed this operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetRequestedByEmail()(*string) {
    return m.requestedByEmail
}
// GetRequestedByUri gets the requestedByUri property value. URI of the user who performed this operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetRequestedByUri()(*string) {
    return m.requestedByUri
}
// GetRequestedOperation gets the requestedOperation property value. One word description of the operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetRequestedOperation()(*string) {
    return m.requestedOperation
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSourceServiceExternalName gets the sourceServiceExternalName property value. External Service Name from where this request was sent
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetSourceServiceExternalName()(*string) {
    return m.sourceServiceExternalName
}
// GetState gets the state property value. state of this operation
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetState()(*string) {
    return m.state
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1DualAuthOperationsGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1DualAuthOperationsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteTimeValue("checkedAt", m.GetCheckedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("checkedByEmail", m.GetCheckedByEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("checkedByUri", m.GetCheckedByUri())
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
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("operationResource", m.GetOperationResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("requestedAt", m.GetRequestedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("requestedByEmail", m.GetRequestedByEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("requestedByUri", m.GetRequestedByUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("requestedOperation", m.GetRequestedOperation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceServiceExternalName", m.GetSourceServiceExternalName())
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
func (m *V1beta1DualAuthOperationsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssociatedResources sets the associatedResources property value. Resources associated with this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetAssociatedResources(value []V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable)() {
    m.associatedResources = value
}
// SetCheckedAt sets the checkedAt property value. Time when this operation was checked. RFC 3339 Timestamp
func (m *V1beta1DualAuthOperationsGetResponse_items) SetCheckedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.checkedAt = value
}
// SetCheckedByEmail sets the checkedByEmail property value. Email of the user who checked (second authorization) this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetCheckedByEmail(value *string)() {
    m.checkedByEmail = value
}
// SetCheckedByUri sets the checkedByUri property value. URI of the user who checked (second authorization) this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetCheckedByUri(value *string)() {
    m.checkedByUri = value
}
// SetConsoleUri sets the consoleUri property value. The URI for console screen that displays this resource. Deprecated - use associatedResources instead
// Deprecated: 
func (m *V1beta1DualAuthOperationsGetResponse_items) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1DualAuthOperationsGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1DualAuthOperationsGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDescription sets the description property value. Detailed description of the operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetDescription(value *string)() {
    m.description = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1DualAuthOperationsGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGroups sets the groups property value. Groups this operation is associated with. Deprecated - use associatedResources instead
// Deprecated: 
func (m *V1beta1DualAuthOperationsGetResponse_items) SetGroups(value []V1beta1DualAuthOperationsGetResponse_items_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1DualAuthOperationsGetResponse_items) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1DualAuthOperationsGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetOperationResource sets the operationResource property value. Operation resource on which the operation is taking place. Deprecated - use associatedResources instead
// Deprecated: 
func (m *V1beta1DualAuthOperationsGetResponse_items) SetOperationResource(value V1beta1DualAuthOperationsGetResponse_items_operationResourceable)() {
    m.operationResource = value
}
// SetRequestedAt sets the requestedAt property value. Time when this operation was requested. RFC 3339 Timestamp
func (m *V1beta1DualAuthOperationsGetResponse_items) SetRequestedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedAt = value
}
// SetRequestedByEmail sets the requestedByEmail property value. Email of the user who performed this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetRequestedByEmail(value *string)() {
    m.requestedByEmail = value
}
// SetRequestedByUri sets the requestedByUri property value. URI of the user who performed this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetRequestedByUri(value *string)() {
    m.requestedByUri = value
}
// SetRequestedOperation sets the requestedOperation property value. One word description of the operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetRequestedOperation(value *string)() {
    m.requestedOperation = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1DualAuthOperationsGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSourceServiceExternalName sets the sourceServiceExternalName property value. External Service Name from where this request was sent
func (m *V1beta1DualAuthOperationsGetResponse_items) SetSourceServiceExternalName(value *string)() {
    m.sourceServiceExternalName = value
}
// SetState sets the state property value. state of this operation
func (m *V1beta1DualAuthOperationsGetResponse_items) SetState(value *string)() {
    m.state = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DualAuthOperationsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1DualAuthOperationsGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1DualAuthOperationsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedResources()([]V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable)
    GetCheckedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCheckedByEmail()(*string)
    GetCheckedByUri()(*string)
    GetConsoleUri()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDescription()(*string)
    GetGeneration()(*int64)
    GetGroups()([]V1beta1DualAuthOperationsGetResponse_items_groupsable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    GetOperationResource()(V1beta1DualAuthOperationsGetResponse_items_operationResourceable)
    GetRequestedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedByEmail()(*string)
    GetRequestedByUri()(*string)
    GetRequestedOperation()(*string)
    GetResourceUri()(*string)
    GetSourceServiceExternalName()(*string)
    GetState()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssociatedResources(value []V1beta1DualAuthOperationsGetResponse_items_associatedResourcesable)()
    SetCheckedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCheckedByEmail(value *string)()
    SetCheckedByUri(value *string)()
    SetConsoleUri(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDescription(value *string)()
    SetGeneration(value *int64)()
    SetGroups(value []V1beta1DualAuthOperationsGetResponse_items_groupsable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
    SetOperationResource(value V1beta1DualAuthOperationsGetResponse_items_operationResourceable)()
    SetRequestedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedByEmail(value *string)()
    SetRequestedByUri(value *string)()
    SetRequestedOperation(value *string)()
    SetResourceUri(value *string)()
    SetSourceServiceExternalName(value *string)()
    SetState(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
