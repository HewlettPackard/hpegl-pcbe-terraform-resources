package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of resources of the specified type
    count *int32
    // The type of the associated resource. Used for grouping of like resources
    typeEscaped *string
}
// NewV1beta1GroupsItemGroupsGetResponse_associatedResourceTypes instantiates a new V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes and sets the default values.
func NewV1beta1GroupsItemGroupsGetResponse_associatedResourceTypes()(*V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) {
    m := &V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1GroupsItemGroupsGetResponse_associatedResourceTypesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1GroupsItemGroupsGetResponse_associatedResourceTypesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1GroupsItemGroupsGetResponse_associatedResourceTypes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The number of resources of the specified type
// returns a *int32 when successful
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
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
// GetTypeEscaped gets the type property value. The type of the associated resource. Used for grouping of like resources
// returns a *string when successful
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The number of resources of the specified type
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) SetCount(value *int32)() {
    m.count = value
}
// SetTypeEscaped sets the type property value. The type of the associated resource. Used for grouping of like resources
func (m *V1beta1GroupsItemGroupsGetResponse_associatedResourceTypes) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1GroupsItemGroupsGetResponse_associatedResourceTypesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetTypeEscaped()(*string)
    SetCount(value *int32)()
    SetTypeEscaped(value *string)()
}
