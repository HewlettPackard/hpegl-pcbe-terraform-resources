package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsItemHostGetResponse_cpuSockets cPU socket information.
type V1beta1HypervisorHostsItemHostGetResponse_cpuSockets struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // CPU bus speed in Hz.
    cpuBusSpeedHz *int64
    // CPU core speed in Hz.
    cpuCoreSpeedHz *int64
    // Description of CPU socket.
    description *string
    // Vendor of CPU socket.
    vendorEscaped *string
}
// NewV1beta1HypervisorHostsItemHostGetResponse_cpuSockets instantiates a new V1beta1HypervisorHostsItemHostGetResponse_cpuSockets and sets the default values.
func NewV1beta1HypervisorHostsItemHostGetResponse_cpuSockets()(*V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) {
    m := &V1beta1HypervisorHostsItemHostGetResponse_cpuSockets{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsItemHostGetResponse_cpuSocketsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsItemHostGetResponse_cpuSocketsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsItemHostGetResponse_cpuSockets(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuBusSpeedHz gets the cpuBusSpeedHz property value. CPU bus speed in Hz.
// returns a *int64 when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetCpuBusSpeedHz()(*int64) {
    return m.cpuBusSpeedHz
}
// GetCpuCoreSpeedHz gets the cpuCoreSpeedHz property value. CPU core speed in Hz.
// returns a *int64 when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetCpuCoreSpeedHz()(*int64) {
    return m.cpuCoreSpeedHz
}
// GetDescription gets the description property value. Description of CPU socket.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpuBusSpeedHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuBusSpeedHz(val)
        }
        return nil
    }
    res["cpuCoreSpeedHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCoreSpeedHz(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val)
        }
        return nil
    }
    return res
}
// GetVendorEscaped gets the vendor property value. Vendor of CPU socket.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) GetVendorEscaped()(*string) {
    return m.vendorEscaped
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("cpuBusSpeedHz", m.GetCpuBusSpeedHz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("cpuCoreSpeedHz", m.GetCpuCoreSpeedHz())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor", m.GetVendorEscaped())
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
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuBusSpeedHz sets the cpuBusSpeedHz property value. CPU bus speed in Hz.
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) SetCpuBusSpeedHz(value *int64)() {
    m.cpuBusSpeedHz = value
}
// SetCpuCoreSpeedHz sets the cpuCoreSpeedHz property value. CPU core speed in Hz.
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) SetCpuCoreSpeedHz(value *int64)() {
    m.cpuCoreSpeedHz = value
}
// SetDescription sets the description property value. Description of CPU socket.
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) SetDescription(value *string)() {
    m.description = value
}
// SetVendorEscaped sets the vendor property value. Vendor of CPU socket.
func (m *V1beta1HypervisorHostsItemHostGetResponse_cpuSockets) SetVendorEscaped(value *string)() {
    m.vendorEscaped = value
}
type V1beta1HypervisorHostsItemHostGetResponse_cpuSocketsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuBusSpeedHz()(*int64)
    GetCpuCoreSpeedHz()(*int64)
    GetDescription()(*string)
    GetVendorEscaped()(*string)
    SetCpuBusSpeedHz(value *int64)()
    SetCpuCoreSpeedHz(value *int64)()
    SetDescription(value *string)()
    SetVendorEscaped(value *string)()
}
