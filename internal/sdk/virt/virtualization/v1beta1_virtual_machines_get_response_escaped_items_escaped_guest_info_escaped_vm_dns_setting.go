package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting dns settings of the guest operating system on this virtual machine.
type V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The vmDns property
    vmDns V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable
    // The vmDnsValues property
    vmDnsValues V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable
}
// NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting instantiates a new V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting()(*V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) {
    m := &V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["vmDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmDns(val.(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable))
        }
        return nil
    }
    res["vmDnsValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmDnsValues(val.(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable))
        }
        return nil
    }
    return res
}
// GetVmDns gets the vmDns property value. The vmDns property
// returns a V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) GetVmDns()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable) {
    return m.vmDns
}
// GetVmDnsValues gets the vmDnsValues property value. The vmDnsValues property
// returns a V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable when successful
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) GetVmDnsValues()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable) {
    return m.vmDnsValues
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("vmDns", m.GetVmDns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vmDnsValues", m.GetVmDnsValues())
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
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVmDns sets the vmDns property value. The vmDns property
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) SetVmDns(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable)() {
    m.vmDns = value
}
// SetVmDnsValues sets the vmDnsValues property value. The vmDnsValues property
func (m *V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting) SetVmDnsValues(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable)() {
    m.vmDnsValues = value
}
type V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVmDns()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable)
    GetVmDnsValues()(V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable)
    SetVmDns(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsable)()
    SetVmDnsValues(value V1beta1VirtualMachinesGetResponse_items_guestInfo_vmDnsSetting_vmDnsValuesable)()
}
