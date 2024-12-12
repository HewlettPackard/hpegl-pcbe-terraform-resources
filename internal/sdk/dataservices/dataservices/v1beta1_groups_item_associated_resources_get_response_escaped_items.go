package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1GroupsItemAssociatedResourcesGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URI for console screen that displays this resource
    consoleUri *string
    // Primary identifier for the customer (GUID) associated with the group.
    customerId *string
    // Unique identifer for the associated resource, usually a UUID
    id *string
    // name could be name/displayName of the resource. It is upto the resource-service to send nameor displayName for this field.
    name *string
    // The 'self' reference for this resource. This is must be unique within the scope DSCC
    resourceUri *string
    // The type of the associated resource. Used for grouping of like resources
    typeEscaped *string
}
// NewV1beta1GroupsItemAssociatedResourcesGetResponse_items instantiates a new V1beta1GroupsItemAssociatedResourcesGetResponse_items and sets the default values.
func NewV1beta1GroupsItemAssociatedResourcesGetResponse_items()(*V1beta1GroupsItemAssociatedResourcesGetResponse_items) {
    m := &V1beta1GroupsItemAssociatedResourcesGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1GroupsItemAssociatedResourcesGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsItemAssociatedResourcesGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsItemAssociatedResourcesGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. The URI for console screen that displays this resource
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCustomerId gets the customerId property value. Primary identifier for the customer (GUID) associated with the group.
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetId gets the id property value. Unique identifer for the associated resource, usually a UUID
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. name could be name/displayName of the resource. It is upto the resource-service to send nameor displayName for this field.
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource. This is must be unique within the scope DSCC
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of the associated resource. Used for grouping of like resources
// returns a *string when successful
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("consoleUri", m.GetConsoleUri())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. The URI for console screen that displays this resource
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCustomerId sets the customerId property value. Primary identifier for the customer (GUID) associated with the group.
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetId sets the id property value. Unique identifer for the associated resource, usually a UUID
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. name could be name/displayName of the resource. It is upto the resource-service to send nameor displayName for this field.
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource. This is must be unique within the scope DSCC
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of the associated resource. Used for grouping of like resources
func (m *V1beta1GroupsItemAssociatedResourcesGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1GroupsItemAssociatedResourcesGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    GetCustomerId()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetConsoleUri(value *string)()
    SetCustomerId(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
