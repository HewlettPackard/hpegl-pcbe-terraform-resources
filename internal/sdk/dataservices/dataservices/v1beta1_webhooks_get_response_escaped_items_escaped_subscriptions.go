package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksGetResponse_items_subscriptions an event subscription specification
type V1beta1WebhooksGetResponse_items_subscriptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of groups that is used to filter events.Only events associated with these groups arepublished to the webhook endpoint.
    groups []V1beta1WebhooksGetResponse_items_subscriptions_groupsable
    // The resourceType property
    resourceType V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable
}
// NewV1beta1WebhooksGetResponse_items_subscriptions instantiates a new V1beta1WebhooksGetResponse_items_subscriptions and sets the default values.
func NewV1beta1WebhooksGetResponse_items_subscriptions()(*V1beta1WebhooksGetResponse_items_subscriptions) {
    m := &V1beta1WebhooksGetResponse_items_subscriptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksGetResponse_items_subscriptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksGetResponse_items_subscriptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksGetResponse_items_subscriptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksGetResponse_items_subscriptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksGetResponse_items_subscriptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1WebhooksGetResponse_items_subscriptions_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1WebhooksGetResponse_items_subscriptions_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1WebhooksGetResponse_items_subscriptions_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksGetResponse_items_subscriptions_resourceTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val.(V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. List of groups that is used to filter events.Only events associated with these groups arepublished to the webhook endpoint.
// returns a []V1beta1WebhooksGetResponse_items_subscriptions_groupsable when successful
func (m *V1beta1WebhooksGetResponse_items_subscriptions) GetGroups()([]V1beta1WebhooksGetResponse_items_subscriptions_groupsable) {
    return m.groups
}
// GetResourceType gets the resourceType property value. The resourceType property
// returns a V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable when successful
func (m *V1beta1WebhooksGetResponse_items_subscriptions) GetResourceType()(V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable) {
    return m.resourceType
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksGetResponse_items_subscriptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("resourceType", m.GetResourceType())
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
func (m *V1beta1WebhooksGetResponse_items_subscriptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroups sets the groups property value. List of groups that is used to filter events.Only events associated with these groups arepublished to the webhook endpoint.
func (m *V1beta1WebhooksGetResponse_items_subscriptions) SetGroups(value []V1beta1WebhooksGetResponse_items_subscriptions_groupsable)() {
    m.groups = value
}
// SetResourceType sets the resourceType property value. The resourceType property
func (m *V1beta1WebhooksGetResponse_items_subscriptions) SetResourceType(value V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable)() {
    m.resourceType = value
}
type V1beta1WebhooksGetResponse_items_subscriptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]V1beta1WebhooksGetResponse_items_subscriptions_groupsable)
    GetResourceType()(V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable)
    SetGroups(value []V1beta1WebhooksGetResponse_items_subscriptions_groupsable)()
    SetResourceType(value V1beta1WebhooksGetResponse_items_subscriptions_resourceTypeable)()
}
