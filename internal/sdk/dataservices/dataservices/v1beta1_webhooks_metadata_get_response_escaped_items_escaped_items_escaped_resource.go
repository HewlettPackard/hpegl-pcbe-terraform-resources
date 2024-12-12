package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksMetadataGetResponse_items_items_resource additional attributes for a Resource metadata item
type V1beta1WebhooksMetadataGetResponse_items_items_resource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the resource type
    applicationResource *string
    // The permission needed to read instances of this resource type
    readPermission *string
    // Indicates whether or not the permission has scopes
    scopedResource *bool
}
// NewV1beta1WebhooksMetadataGetResponse_items_items_resource instantiates a new V1beta1WebhooksMetadataGetResponse_items_items_resource and sets the default values.
func NewV1beta1WebhooksMetadataGetResponse_items_items_resource()(*V1beta1WebhooksMetadataGetResponse_items_items_resource) {
    m := &V1beta1WebhooksMetadataGetResponse_items_items_resource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksMetadataGetResponse_items_items_resourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksMetadataGetResponse_items_items_resourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksMetadataGetResponse_items_items_resource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationResource gets the applicationResource property value. Identifies the resource type
// returns a *string when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) GetApplicationResource()(*string) {
    return m.applicationResource
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationResource(val)
        }
        return nil
    }
    res["readPermission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadPermission(val)
        }
        return nil
    }
    res["scopedResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopedResource(val)
        }
        return nil
    }
    return res
}
// GetReadPermission gets the readPermission property value. The permission needed to read instances of this resource type
// returns a *string when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) GetReadPermission()(*string) {
    return m.readPermission
}
// GetScopedResource gets the scopedResource property value. Indicates whether or not the permission has scopes
// returns a *bool when successful
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) GetScopedResource()(*bool) {
    return m.scopedResource
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationResource", m.GetApplicationResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("readPermission", m.GetReadPermission())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("scopedResource", m.GetScopedResource())
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
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationResource sets the applicationResource property value. Identifies the resource type
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) SetApplicationResource(value *string)() {
    m.applicationResource = value
}
// SetReadPermission sets the readPermission property value. The permission needed to read instances of this resource type
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) SetReadPermission(value *string)() {
    m.readPermission = value
}
// SetScopedResource sets the scopedResource property value. Indicates whether or not the permission has scopes
func (m *V1beta1WebhooksMetadataGetResponse_items_items_resource) SetScopedResource(value *bool)() {
    m.scopedResource = value
}
type V1beta1WebhooksMetadataGetResponse_items_items_resourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationResource()(*string)
    GetReadPermission()(*string)
    GetScopedResource()(*bool)
    SetApplicationResource(value *string)()
    SetReadPermission(value *string)()
    SetScopedResource(value *bool)()
}
