package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsGetResponse_items_groups storage Cloud Group Identity
type V1beta1SecretAssignmentsGetResponse_items_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifier of the associated group
    id *string
    // Name of the associated group
    name *string
    // URI of the associated group
    resourceUri *string
}
// NewV1beta1SecretAssignmentsGetResponse_items_groups instantiates a new V1beta1SecretAssignmentsGetResponse_items_groups and sets the default values.
func NewV1beta1SecretAssignmentsGetResponse_items_groups()(*V1beta1SecretAssignmentsGetResponse_items_groups) {
    m := &V1beta1SecretAssignmentsGetResponse_items_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretAssignmentsGetResponse_items_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsGetResponse_items_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsGetResponse_items_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetId gets the id property value. Identifier of the associated group
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the associated group
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. URI of the associated group
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) GetResourceUri()(*string) {
    return m.resourceUri
}
// Serialize serializes information the current object
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Identifier of the associated group
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the associated group
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. URI of the associated group
func (m *V1beta1SecretAssignmentsGetResponse_items_groups) SetResourceUri(value *string)() {
    m.resourceUri = value
}
type V1beta1SecretAssignmentsGetResponse_items_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
}
