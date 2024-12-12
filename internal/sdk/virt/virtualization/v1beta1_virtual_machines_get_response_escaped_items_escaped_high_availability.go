package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_highAvailability high availability settings of the virtual machine.
type V1beta1VirtualMachinesGetResponse_items_highAvailability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // High availability is enabled or not.
    enabled *bool
}
// NewV1beta1VirtualMachinesGetResponse_items_highAvailability instantiates a new V1beta1VirtualMachinesGetResponse_items_highAvailability and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_highAvailability()(*V1beta1VirtualMachinesGetResponse_items_highAvailability) {
    m := &V1beta1VirtualMachinesGetResponse_items_highAvailability{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_highAvailabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_highAvailabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_highAvailability(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. High availability is enabled or not.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. High availability is enabled or not.
func (m *V1beta1VirtualMachinesGetResponse_items_highAvailability) SetEnabled(value *bool)() {
    m.enabled = value
}
type V1beta1VirtualMachinesGetResponse_items_highAvailabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    SetEnabled(value *bool)()
}
