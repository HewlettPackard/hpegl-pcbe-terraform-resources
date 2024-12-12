package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory reconfigure memory for a virtual machine. The supported range of memory depends on guest operating system and virtual hardware version of the virtual machine.
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // New memory size in mebibytes.
    memoryInMb *int32
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["memoryInMb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryInMb(val)
        }
        return nil
    }
    return res
}
// GetMemoryInMb gets the memoryInMb property value. New memory size in mebibytes.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) GetMemoryInMb()(*int32) {
    return m.memoryInMb
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("memoryInMb", m.GetMemoryInMb())
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMemoryInMb sets the memoryInMb property value. New memory size in mebibytes.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memory) SetMemoryInMb(value *int32)() {
    m.memoryInMb = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_memoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMemoryInMb()(*int32)
    SetMemoryInMb(value *int32)()
}
