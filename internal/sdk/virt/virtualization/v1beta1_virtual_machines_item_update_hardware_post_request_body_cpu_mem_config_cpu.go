package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu reconfigure CPU values for a virtual machine
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of CPU cores per socket in the virtual machine. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket.
    numOfCoresPerSocket *int32
    // The number of CPU cores in the virtual machine. The supported range of CPU depends on guest operating system and virtual hardware version of the virtual machine.
    numOfCpus *int32
}
// NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu instantiates a new V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu()(*V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) {
    m := &V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["numOfCoresPerSocket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumOfCoresPerSocket(val)
        }
        return nil
    }
    res["numOfCpus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumOfCpus(val)
        }
        return nil
    }
    return res
}
// GetNumOfCoresPerSocket gets the numOfCoresPerSocket property value. The number of CPU cores per socket in the virtual machine. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) GetNumOfCoresPerSocket()(*int32) {
    return m.numOfCoresPerSocket
}
// GetNumOfCpus gets the numOfCpus property value. The number of CPU cores in the virtual machine. The supported range of CPU depends on guest operating system and virtual hardware version of the virtual machine.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) GetNumOfCpus()(*int32) {
    return m.numOfCpus
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("numOfCoresPerSocket", m.GetNumOfCoresPerSocket())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numOfCpus", m.GetNumOfCpus())
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
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumOfCoresPerSocket sets the numOfCoresPerSocket property value. The number of CPU cores per socket in the virtual machine. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) SetNumOfCoresPerSocket(value *int32)() {
    m.numOfCoresPerSocket = value
}
// SetNumOfCpus sets the numOfCpus property value. The number of CPU cores in the virtual machine. The supported range of CPU depends on guest operating system and virtual hardware version of the virtual machine.
func (m *V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpu) SetNumOfCpus(value *int32)() {
    m.numOfCpus = value
}
type V1beta1VirtualMachinesItemUpdateHardwarePostRequestBody_cpuMemConfig_cpuable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumOfCoresPerSocket()(*int32)
    GetNumOfCpus()(*int32)
    SetNumOfCoresPerSocket(value *int32)()
    SetNumOfCpus(value *int32)()
}
