package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretAssignmentsGetResponse_items_appliance appliance Identity
type V1beta1SecretAssignmentsGetResponse_items_appliance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identity of the appliance
    id *string
}
// NewV1beta1SecretAssignmentsGetResponse_items_appliance instantiates a new V1beta1SecretAssignmentsGetResponse_items_appliance and sets the default values.
func NewV1beta1SecretAssignmentsGetResponse_items_appliance()(*V1beta1SecretAssignmentsGetResponse_items_appliance) {
    m := &V1beta1SecretAssignmentsGetResponse_items_appliance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretAssignmentsGetResponse_items_applianceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretAssignmentsGetResponse_items_applianceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretAssignmentsGetResponse_items_appliance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetId gets the id property value. Identity of the appliance
// returns a *string when successful
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Identity of the appliance
func (m *V1beta1SecretAssignmentsGetResponse_items_appliance) SetId(value *string)() {
    m.id = value
}
type V1beta1SecretAssignmentsGetResponse_items_applianceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    SetId(value *string)()
}