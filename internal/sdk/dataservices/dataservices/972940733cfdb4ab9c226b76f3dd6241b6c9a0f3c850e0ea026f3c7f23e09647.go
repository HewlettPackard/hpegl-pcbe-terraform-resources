package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource resource associated with this operation
type V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URI for the console screen that displays this resource
    consoleUri *string
    // The name property
    name *string
    // The resourceUri property
    resourceUri *string
    // The type property
    typeEscaped *string
}
// NewV1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource instantiates a new V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource and sets the default values.
func NewV1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource()(*V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) {
    m := &V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. The URI for the console screen that displays this resource
// returns a *string when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The resourceUri property
// returns a *string when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. The URI for the console screen that displays this resource
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetName sets the name property value. The name property
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The resourceUri property
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resource) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1DualAuthOperationsItemDualAuthOperationsGetResponse_associatedResources_resourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetConsoleUri(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
