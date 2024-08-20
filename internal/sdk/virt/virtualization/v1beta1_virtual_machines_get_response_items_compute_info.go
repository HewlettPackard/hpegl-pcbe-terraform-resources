package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_computeInfo compute information of the virtual machine.
type V1beta1VirtualMachinesGetResponse_items_computeInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total memory provisioned.
    memorySizeInMib *string
    // Number of CPU cores provisioned.
    numCpuCores *int32
    // Number of CPU threads provisioned.
    numCpuThreads *int32
}
// NewV1beta1VirtualMachinesGetResponse_items_computeInfo instantiates a new V1beta1VirtualMachinesGetResponse_items_computeInfo and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_computeInfo()(*V1beta1VirtualMachinesGetResponse_items_computeInfo) {
    m := &V1beta1VirtualMachinesGetResponse_items_computeInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_computeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_computeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_computeInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["memorySizeInMib"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemorySizeInMib(val)
        }
        return nil
    }
    res["numCpuCores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumCpuCores(val)
        }
        return nil
    }
    res["numCpuThreads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumCpuThreads(val)
        }
        return nil
    }
    return res
}
// GetMemorySizeInMib gets the memorySizeInMib property value. Total memory provisioned.
// returns a *string when successful
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) GetMemorySizeInMib()(*string) {
    return m.memorySizeInMib
}
// GetNumCpuCores gets the numCpuCores property value. Number of CPU cores provisioned.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) GetNumCpuCores()(*int32) {
    return m.numCpuCores
}
// GetNumCpuThreads gets the numCpuThreads property value. Number of CPU threads provisioned.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) GetNumCpuThreads()(*int32) {
    return m.numCpuThreads
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("memorySizeInMib", m.GetMemorySizeInMib())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numCpuCores", m.GetNumCpuCores())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numCpuThreads", m.GetNumCpuThreads())
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
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMemorySizeInMib sets the memorySizeInMib property value. Total memory provisioned.
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) SetMemorySizeInMib(value *string)() {
    m.memorySizeInMib = value
}
// SetNumCpuCores sets the numCpuCores property value. Number of CPU cores provisioned.
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) SetNumCpuCores(value *int32)() {
    m.numCpuCores = value
}
// SetNumCpuThreads sets the numCpuThreads property value. Number of CPU threads provisioned.
func (m *V1beta1VirtualMachinesGetResponse_items_computeInfo) SetNumCpuThreads(value *int32)() {
    m.numCpuThreads = value
}
type V1beta1VirtualMachinesGetResponse_items_computeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMemorySizeInMib()(*string)
    GetNumCpuCores()(*int32)
    GetNumCpuThreads()(*int32)
    SetMemorySizeInMib(value *string)()
    SetNumCpuCores(value *int32)()
    SetNumCpuThreads(value *int32)()
}
