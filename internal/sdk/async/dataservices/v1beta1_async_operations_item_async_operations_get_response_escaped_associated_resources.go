package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources references to other resources include the name, type and URI of the other resource.
type V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The resourceUri property
    resourceUri *string
    // The type property
    typeEscaped *string
}
// NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources instantiates a new V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources and sets the default values.
func NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources()(*V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) {
    m := &V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResourcesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResourcesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The resourceUri property
// returns a *string when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The resourceUri property
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResources) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1AsyncOperationsItemAsyncOperationsGetResponse_associatedResourcesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
