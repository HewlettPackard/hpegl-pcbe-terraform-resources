package dataservices

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the associated group
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Name of the associated group
    name *string
}
// NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups instantiates a new V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups and sets the default values.
func NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups()(*V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) {
    m := &V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetId gets the id property value. ID of the associated group
// returns a *UUID when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetName gets the name property value. Name of the associated group
// returns a *string when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("id", m.GetId())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. ID of the associated group
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetName sets the name property value. Name of the associated group
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groups) SetName(value *string)() {
    m.name = value
}
type V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_groupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
}
