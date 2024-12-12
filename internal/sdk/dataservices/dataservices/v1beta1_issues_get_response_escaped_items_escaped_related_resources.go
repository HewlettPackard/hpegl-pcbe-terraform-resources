package dataservices

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1IssuesGetResponse_items_relatedResources reference to another resource and additional id for use with Issues
type V1beta1IssuesGetResponse_items_relatedResources struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // console URI of the related object resource that is the source of the issue
    consoleUri *string
    // The id property
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The name property
    name *string
    // The resourceUri property
    resourceUri *string
    // The type property
    typeEscaped *string
}
// NewV1beta1IssuesGetResponse_items_relatedResources instantiates a new V1beta1IssuesGetResponse_items_relatedResources and sets the default values.
func NewV1beta1IssuesGetResponse_items_relatedResources()(*V1beta1IssuesGetResponse_items_relatedResources) {
    m := &V1beta1IssuesGetResponse_items_relatedResources{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1IssuesGetResponse_items_relatedResourcesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1IssuesGetResponse_items_relatedResourcesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1IssuesGetResponse_items_relatedResources(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. console URI of the related object resource that is the source of the issue
// returns a *string when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetId gets the id property value. The id property
// returns a *UUID when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The resourceUri property
// returns a *string when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *V1beta1IssuesGetResponse_items_relatedResources) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1IssuesGetResponse_items_relatedResources) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("consoleUri", m.GetConsoleUri())
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
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. console URI of the related object resource that is the source of the issue
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetId sets the id property value. The id property
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The resourceUri property
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *V1beta1IssuesGetResponse_items_relatedResources) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1IssuesGetResponse_items_relatedResourcesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetConsoleUri(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
