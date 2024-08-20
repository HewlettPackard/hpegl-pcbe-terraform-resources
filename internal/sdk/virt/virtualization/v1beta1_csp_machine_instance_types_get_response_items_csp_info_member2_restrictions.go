package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions a restriction because of which SKU cannot be used.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reason for the restriction.
    reasonCode *string
    // Type of restriction.
    typeEscaped *string
    // Restriction value.
    values []string
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reasonCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReasonCode(val)
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
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetReasonCode gets the reasonCode property value. Reason for the restriction.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) GetReasonCode()(*string) {
    return m.reasonCode
}
// GetTypeEscaped gets the type property value. Type of restriction.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValues gets the values property value. Restriction value.
// returns a []string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) GetValues()([]string) {
    return m.values
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reasonCode", m.GetReasonCode())
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
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReasonCode sets the reasonCode property value. Reason for the restriction.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) SetReasonCode(value *string)() {
    m.reasonCode = value
}
// SetTypeEscaped sets the type property value. Type of restriction.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValues sets the values property value. Restriction value.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictions) SetValues(value []string)() {
    m.values = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember2_restrictionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReasonCode()(*string)
    GetTypeEscaped()(*string)
    GetValues()([]string)
    SetReasonCode(value *string)()
    SetTypeEscaped(value *string)()
    SetValues(value []string)()
}
