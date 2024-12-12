package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo information of the Virtual Machine Protection Group.
type V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifier for the Protection Group.
    id *string
    // Name of the Protection Group.
    name *string
    // Reference to resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo instantiates a new V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo and sets the default values.
func NewV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo()(*V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) {
    m := &V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetId gets the id property value. Unique identifier for the Protection Group.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the Protection Group.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. Reference to resource.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Unique identifier for the Protection Group.
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the Protection Group.
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. Reference to resource.
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1DatastoresGetResponse_items_vmProtectionGroupsInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
