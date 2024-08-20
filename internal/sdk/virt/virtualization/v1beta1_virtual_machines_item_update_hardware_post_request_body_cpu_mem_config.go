package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig reconfigure CPU and memory values for a virtual machine
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reconfigure CPU values for a virtual machine
    cpu V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable
    // Reconfigure memory for a virtual machine. The supported range of memory depends on guest operating system and virtual hardware version of the virtual machine.
    memory V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpu gets the cpu property value. Reconfigure CPU values for a virtual machine
// returns a V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) GetCpu()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable) {
    return m.cpu
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpu"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpu(val.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable))
        }
        return nil
    }
    res["memory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemory(val.(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable))
        }
        return nil
    }
    return res
}
// GetMemory gets the memory property value. Reconfigure memory for a virtual machine. The supported range of memory depends on guest operating system and virtual hardware version of the virtual machine.
// returns a V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) GetMemory()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable) {
    return m.memory
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cpu", m.GetCpu())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("memory", m.GetMemory())
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpu sets the cpu property value. Reconfigure CPU values for a virtual machine
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) SetCpu(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable)() {
    m.cpu = value
}
// SetMemory sets the memory property value. Reconfigure memory for a virtual machine. The supported range of memory depends on guest operating system and virtual hardware version of the virtual machine.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig) SetMemory(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable)() {
    m.memory = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpu()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable)
    GetMemory()(V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable)
    SetCpu(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable)()
    SetMemory(value V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable)()
}
