package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksItemWebhooksGetResponse a Webhook resource
type V1beta1WebhooksItemWebhooksGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An object containing an identity that is used to authorize published events
    authorization V1beta1WebhooksItemWebhooksGetResponse_authorizationable
    // URI that opens the webhook in the UI console
    consoleUri *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // List of event destinations
    destinations []V1beta1WebhooksItemWebhooksGetResponse_destinationsable
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // List of group associations
    groups []V1beta1WebhooksItemWebhooksGetResponse_groupsable
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The date & timestamp than an event was last delivered to the destination.The value is null if an event has never been delivered.
    lastEventDeliveredAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // List of subscriptions that identify the events that are subscribed toand will be published to the Webhook destination.
    subscriptions []V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1WebhooksItemWebhooksGetResponse instantiates a new V1beta1WebhooksItemWebhooksGetResponse and sets the default values.
func NewV1beta1WebhooksItemWebhooksGetResponse()(*V1beta1WebhooksItemWebhooksGetResponse) {
    m := &V1beta1WebhooksItemWebhooksGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksItemWebhooksGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksItemWebhooksGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksItemWebhooksGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorization gets the authorization property value. An object containing an identity that is used to authorize published events
// returns a V1beta1WebhooksItemWebhooksGetResponse_authorizationable when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetAuthorization()(V1beta1WebhooksItemWebhooksGetResponse_authorizationable) {
    return m.authorization
}
// GetConsoleUri gets the consoleUri property value. URI that opens the webhook in the UI console
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDestinations gets the destinations property value. List of event destinations
// returns a []V1beta1WebhooksItemWebhooksGetResponse_destinationsable when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetDestinations()([]V1beta1WebhooksItemWebhooksGetResponse_destinationsable) {
    return m.destinations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksItemWebhooksGetResponse_authorizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorization(val.(V1beta1WebhooksItemWebhooksGetResponse_authorizationable))
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
    res["destinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1WebhooksItemWebhooksGetResponse_destinationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1WebhooksItemWebhooksGetResponse_destinationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1WebhooksItemWebhooksGetResponse_destinationsable)
                }
            }
            m.SetDestinations(res)
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
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1WebhooksItemWebhooksGetResponse_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1WebhooksItemWebhooksGetResponse_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1WebhooksItemWebhooksGetResponse_groupsable)
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
    res["lastEventDeliveredAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEventDeliveredAt(val)
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
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1WebhooksItemWebhooksGetResponse_subscriptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable)
                }
            }
            m.SetSubscriptions(res)
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
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetGroups gets the groups property value. List of group associations
// returns a []V1beta1WebhooksItemWebhooksGetResponse_groupsable when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetGroups()([]V1beta1WebhooksItemWebhooksGetResponse_groupsable) {
    return m.groups
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetLastEventDeliveredAt gets the lastEventDeliveredAt property value. The date & timestamp than an event was last delivered to the destination.The value is null if an event has never been delivered.
// returns a *Time when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetLastEventDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastEventDeliveredAt
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSubscriptions gets the subscriptions property value. List of subscriptions that identify the events that are subscribed toand will be published to the Webhook destination.
// returns a []V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetSubscriptions()([]V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable) {
    return m.subscriptions
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1WebhooksItemWebhooksGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksItemWebhooksGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authorization", m.GetAuthorization())
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
    if m.GetDestinations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDestinations()))
        for i, v := range m.GetDestinations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("destinations", cast)
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
    if m.GetSubscriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("subscriptions", cast)
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
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorization sets the authorization property value. An object containing an identity that is used to authorize published events
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetAuthorization(value V1beta1WebhooksItemWebhooksGetResponse_authorizationable)() {
    m.authorization = value
}
// SetConsoleUri sets the consoleUri property value. URI that opens the webhook in the UI console
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDestinations sets the destinations property value. List of event destinations
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetDestinations(value []V1beta1WebhooksItemWebhooksGetResponse_destinationsable)() {
    m.destinations = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGroups sets the groups property value. List of group associations
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetGroups(value []V1beta1WebhooksItemWebhooksGetResponse_groupsable)() {
    m.groups = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetLastEventDeliveredAt sets the lastEventDeliveredAt property value. The date & timestamp than an event was last delivered to the destination.The value is null if an event has never been delivered.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetLastEventDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastEventDeliveredAt = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSubscriptions sets the subscriptions property value. List of subscriptions that identify the events that are subscribed toand will be published to the Webhook destination.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetSubscriptions(value []V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable)() {
    m.subscriptions = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1WebhooksItemWebhooksGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1WebhooksItemWebhooksGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorization()(V1beta1WebhooksItemWebhooksGetResponse_authorizationable)
    GetConsoleUri()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDestinations()([]V1beta1WebhooksItemWebhooksGetResponse_destinationsable)
    GetGeneration()(*int64)
    GetGroups()([]V1beta1WebhooksItemWebhooksGetResponse_groupsable)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastEventDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetResourceUri()(*string)
    GetSubscriptions()([]V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAuthorization(value V1beta1WebhooksItemWebhooksGetResponse_authorizationable)()
    SetConsoleUri(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDestinations(value []V1beta1WebhooksItemWebhooksGetResponse_destinationsable)()
    SetGeneration(value *int64)()
    SetGroups(value []V1beta1WebhooksItemWebhooksGetResponse_groupsable)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastEventDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetSubscriptions(value []V1beta1WebhooksItemWebhooksGetResponse_subscriptionsable)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
