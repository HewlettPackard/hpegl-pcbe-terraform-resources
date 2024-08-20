package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo describes applied protection policy information.
type V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // UUID string uniquely identifying the protection policy.
    id *string
    // Name of the protection policy.
    name *string
    // The URI reference for this resource.
    resourceUri *string
}
// NewV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo instantiates a new V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo and sets the default values.
func NewV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo()(*V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) {
    m := &V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetId gets the id property value. UUID string uniquely identifying the protection policy.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the protection policy.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. UUID string uniquely identifying the protection policy.
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the protection policy.
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
type V1beta1DatastoresGetResponse_items_protectionPolicyAppliedAtInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
}
