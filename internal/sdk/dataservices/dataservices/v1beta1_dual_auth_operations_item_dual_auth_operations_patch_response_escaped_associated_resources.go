package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Groups this resource is associated with
    groups []V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable
    // Resource associated with this operation
    resource V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable
}
// NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources instantiates a new V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources and sets the default values.
func NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources()(*V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) {
    m := &V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResourcesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResourcesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Groups this resource is associated with
// returns a []V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) GetGroups()([]V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable) {
    return m.groups
}
// GetResource gets the resource property value. Resource associated with this operation
// returns a V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable when successful
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) GetResource()(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable) {
    return m.resource
}
// Serialize serializes information the current object
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("resource", m.GetResource())
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
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroups sets the groups property value. Groups this resource is associated with
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) SetGroups(value []V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable)() {
    m.groups = value
}
// SetResource sets the resource property value. Resource associated with this operation
func (m *V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources) SetResource(value V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable)() {
    m.resource = value
}
type V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResourcesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroups()([]V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable)
    GetResource()(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable)
    SetGroups(value []V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_groupsable)()
    SetResource(value V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponse_associatedResources_resourceable)()
}
