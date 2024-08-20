package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody reconfigure virtual machine hardware settings - CPU, memory, network adapters, and disks
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reconfigure CPU and memory values for a virtual machine
    cpuMemConfig V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable
    // Reconfigure network adapter(s) for a virtual machine. ADD/EDIT/DELETE list of network adapters.
    networkAdapters []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable
    // Reconfigure disk(s) for a virtual machine. ADD/EDIT/DELETE list of disks.
    virtualDisks []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuMemConfig gets the cpuMemConfig property value. Reconfigure CPU and memory values for a virtual machine
// returns a V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) GetCpuMemConfig()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable) {
    return m.cpuMemConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpuMemConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuMemConfig(val.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable))
        }
        return nil
    }
    res["networkAdapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable)
                }
            }
            m.SetNetworkAdapters(res)
        }
        return nil
    }
    res["virtualDisks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable)
                }
            }
            m.SetVirtualDisks(res)
        }
        return nil
    }
    return res
}
// GetNetworkAdapters gets the networkAdapters property value. Reconfigure network adapter(s) for a virtual machine. ADD/EDIT/DELETE list of network adapters.
// returns a []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) GetNetworkAdapters()([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable) {
    return m.networkAdapters
}
// GetVirtualDisks gets the virtualDisks property value. Reconfigure disk(s) for a virtual machine. ADD/EDIT/DELETE list of disks.
// returns a []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) GetVirtualDisks()([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable) {
    return m.virtualDisks
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cpuMemConfig", m.GetCpuMemConfig())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkAdapters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkAdapters()))
        for i, v := range m.GetNetworkAdapters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkAdapters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVirtualDisks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVirtualDisks()))
        for i, v := range m.GetVirtualDisks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("virtualDisks", cast)
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuMemConfig sets the cpuMemConfig property value. Reconfigure CPU and memory values for a virtual machine
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) SetCpuMemConfig(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable)() {
    m.cpuMemConfig = value
}
// SetNetworkAdapters sets the networkAdapters property value. Reconfigure network adapter(s) for a virtual machine. ADD/EDIT/DELETE list of network adapters.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) SetNetworkAdapters(value []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable)() {
    m.networkAdapters = value
}
// SetVirtualDisks sets the virtualDisks property value. Reconfigure disk(s) for a virtual machine. ADD/EDIT/DELETE list of disks.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody) SetVirtualDisks(value []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable)() {
    m.virtualDisks = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuMemConfig()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable)
    GetNetworkAdapters()([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable)
    GetVirtualDisks()([]V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable)
    SetCpuMemConfig(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable)()
    SetNetworkAdapters(value []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_networkAdaptersable)()
    SetVirtualDisks(value []V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_virtualDisksable)()
}
