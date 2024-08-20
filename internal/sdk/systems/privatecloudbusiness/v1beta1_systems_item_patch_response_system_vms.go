package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchResponse_systemVms system Virtual Machines Information.  These are not workload virtual machines. These virtual machines are part of the system itself - e.g. the Omnicube virtual machines in a system with stackType SIMPLIVITY.
type V1beta1SystemsItemPatchResponse_systemVms struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the system Virtual Machine.
    name *string
    // Unique Identifier of the system virtual machine, usually a UUID.
    uuid *string
}
// NewV1beta1SystemsItemPatchResponse_systemVms instantiates a new V1beta1SystemsItemPatchResponse_systemVms and sets the default values.
func NewV1beta1SystemsItemPatchResponse_systemVms()(*V1beta1SystemsItemPatchResponse_systemVms) {
    m := &V1beta1SystemsItemPatchResponse_systemVms{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchResponse_systemVmsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchResponse_systemVmsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchResponse_systemVms(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchResponse_systemVms) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchResponse_systemVms) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the system Virtual Machine.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_systemVms) GetName()(*string) {
    return m.name
}
// GetUuid gets the uuid property value. Unique Identifier of the system virtual machine, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemPatchResponse_systemVms) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchResponse_systemVms) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uuid", m.GetUuid())
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
func (m *V1beta1SystemsItemPatchResponse_systemVms) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the system Virtual Machine.
func (m *V1beta1SystemsItemPatchResponse_systemVms) SetName(value *string)() {
    m.name = value
}
// SetUuid sets the uuid property value. Unique Identifier of the system virtual machine, usually a UUID.
func (m *V1beta1SystemsItemPatchResponse_systemVms) SetUuid(value *string)() {
    m.uuid = value
}
type V1beta1SystemsItemPatchResponse_systemVmsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetUuid()(*string)
    SetName(value *string)()
    SetUuid(value *string)()
}
