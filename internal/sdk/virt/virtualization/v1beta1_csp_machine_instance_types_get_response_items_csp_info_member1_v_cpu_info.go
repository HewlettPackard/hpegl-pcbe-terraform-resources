package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo the vCPU configurations for the instance type.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The default number of cores for the instance type.
    defaultCores *int32
    // The default number of threads per core for the instance type.
    defaultThreadsPerCore *int32
    // The default number of vCPUs for the instance type.
    defaultVcpus *int32
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultCores gets the defaultCores property value. The default number of cores for the instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) GetDefaultCores()(*int32) {
    return m.defaultCores
}
// GetDefaultThreadsPerCore gets the defaultThreadsPerCore property value. The default number of threads per core for the instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) GetDefaultThreadsPerCore()(*int32) {
    return m.defaultThreadsPerCore
}
// GetDefaultVcpus gets the defaultVcpus property value. The default number of vCPUs for the instance type.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) GetDefaultVcpus()(*int32) {
    return m.defaultVcpus
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultCores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCores(val)
        }
        return nil
    }
    res["defaultThreadsPerCore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultThreadsPerCore(val)
        }
        return nil
    }
    res["defaultVcpus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultVcpus(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("defaultCores", m.GetDefaultCores())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("defaultThreadsPerCore", m.GetDefaultThreadsPerCore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("defaultVcpus", m.GetDefaultVcpus())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultCores sets the defaultCores property value. The default number of cores for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) SetDefaultCores(value *int32)() {
    m.defaultCores = value
}
// SetDefaultThreadsPerCore sets the defaultThreadsPerCore property value. The default number of threads per core for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) SetDefaultThreadsPerCore(value *int32)() {
    m.defaultThreadsPerCore = value
}
// SetDefaultVcpus sets the defaultVcpus property value. The default number of vCPUs for the instance type.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfo) SetDefaultVcpus(value *int32)() {
    m.defaultVcpus = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_vCpuInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultCores()(*int32)
    GetDefaultThreadsPerCore()(*int32)
    GetDefaultVcpus()(*int32)
    SetDefaultCores(value *int32)()
    SetDefaultThreadsPerCore(value *int32)()
    SetDefaultVcpus(value *int32)()
}
