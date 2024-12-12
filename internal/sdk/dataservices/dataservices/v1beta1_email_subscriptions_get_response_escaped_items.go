package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1EmailSubscriptionsGetResponse_items common properties included in all resource models.
type V1beta1EmailSubscriptionsGetResponse_items struct {
    // If true, user activated this email subscription. Otherwise, false.
    activated *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Categories of email notification. This field, along with serviceName, will uniquely identify an email subscription. The match of Categories requires the same number of elements and each element in the array must match exactly. Case sensitivity is observed.
    categories []string
    // The "consoleUri" property is the webpage representing the resource in the Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/email-subscriptions/c4f8f464-ce09-4847-bfaa-e668de838c66")
    consoleUri *string
    // Longer name or description of the subscription that can be provided as a tool tip or when hovering over a subscription in the microapp or plugin.
    description *string
    // Groups the email subscription is associated with.
    groups []V1beta1EmailSubscriptionsGetResponse_items_groupsable
    // The name of the service defining the subscription. This field along with Categories will uniquely identify an email subscription.
    serviceName *string
}
// NewV1beta1EmailSubscriptionsGetResponse_items instantiates a new V1beta1EmailSubscriptionsGetResponse_items and sets the default values.
func NewV1beta1EmailSubscriptionsGetResponse_items()(*V1beta1EmailSubscriptionsGetResponse_items) {
    m := &V1beta1EmailSubscriptionsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1EmailSubscriptionsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1EmailSubscriptionsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1EmailSubscriptionsGetResponse_items(), nil
}
// GetActivated gets the activated property value. If true, user activated this email subscription. Otherwise, false.
// returns a *bool when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetActivated()(*bool) {
    return m.activated
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCategories gets the categories property value. Categories of email notification. This field, along with serviceName, will uniquely identify an email subscription. The match of Categories requires the same number of elements and each element in the array must match exactly. Case sensitivity is observed.
// returns a []string when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetCategories()([]string) {
    return m.categories
}
// GetConsoleUri gets the consoleUri property value. The "consoleUri" property is the webpage representing the resource in the Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/email-subscriptions/c4f8f464-ce09-4847-bfaa-e668de838c66")
// returns a *string when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetDescription gets the description property value. Longer name or description of the subscription that can be provided as a tool tip or when hovering over a subscription in the microapp or plugin.
// returns a *string when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivated(val)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCategories(res)
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1EmailSubscriptionsGetResponse_items_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1EmailSubscriptionsGetResponse_items_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1EmailSubscriptionsGetResponse_items_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Groups the email subscription is associated with.
// returns a []V1beta1EmailSubscriptionsGetResponse_items_groupsable when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetGroups()([]V1beta1EmailSubscriptionsGetResponse_items_groupsable) {
    return m.groups
}
// GetServiceName gets the serviceName property value. The name of the service defining the subscription. This field along with Categories will uniquely identify an email subscription.
// returns a *string when successful
func (m *V1beta1EmailSubscriptionsGetResponse_items) GetServiceName()(*string) {
    return m.serviceName
}
// Serialize serializes information the current object
func (m *V1beta1EmailSubscriptionsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("activated", m.GetActivated())
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
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
        err := writer.WriteStringValue("serviceName", m.GetServiceName())
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
// SetActivated sets the activated property value. If true, user activated this email subscription. Otherwise, false.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetActivated(value *bool)() {
    m.activated = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCategories sets the categories property value. Categories of email notification. This field, along with serviceName, will uniquely identify an email subscription. The match of Categories requires the same number of elements and each element in the array must match exactly. Case sensitivity is observed.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetCategories(value []string)() {
    m.categories = value
}
// SetConsoleUri sets the consoleUri property value. The "consoleUri" property is the webpage representing the resource in the Storage Central UI. The the value of "consoleUri" is the console path including any required query parameters (e.g. "/email-subscriptions/c4f8f464-ce09-4847-bfaa-e668de838c66")
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetDescription sets the description property value. Longer name or description of the subscription that can be provided as a tool tip or when hovering over a subscription in the microapp or plugin.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetDescription(value *string)() {
    m.description = value
}
// SetGroups sets the groups property value. Groups the email subscription is associated with.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetGroups(value []V1beta1EmailSubscriptionsGetResponse_items_groupsable)() {
    m.groups = value
}
// SetServiceName sets the serviceName property value. The name of the service defining the subscription. This field along with Categories will uniquely identify an email subscription.
func (m *V1beta1EmailSubscriptionsGetResponse_items) SetServiceName(value *string)() {
    m.serviceName = value
}
type V1beta1EmailSubscriptionsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivated()(*bool)
    GetCategories()([]string)
    GetConsoleUri()(*string)
    GetDescription()(*string)
    GetGroups()([]V1beta1EmailSubscriptionsGetResponse_items_groupsable)
    GetServiceName()(*string)
    SetActivated(value *bool)()
    SetCategories(value []string)()
    SetConsoleUri(value *string)()
    SetDescription(value *string)()
    SetGroups(value []V1beta1EmailSubscriptionsGetResponse_items_groupsable)()
    SetServiceName(value *string)()
}
