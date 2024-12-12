package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos quality of service.
type V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IOPS limit for this volume. If iopsLimit is set to -1, then the volume has no IOPS limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
    iopsLimit *float64
    // Throughput limit for this volume in MB/s. If mbpsLimit is set to -1, then the volume has no MBPS limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
    mbpsLimit *float64
}
// NewV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos instantiates a new V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos()(*V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) {
    m := &V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qosFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qosFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["iopsLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIopsLimit(val)
        }
        return nil
    }
    res["mbpsLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMbpsLimit(val)
        }
        return nil
    }
    return res
}
// GetIopsLimit gets the iopsLimit property value. IOPS limit for this volume. If iopsLimit is set to -1, then the volume has no IOPS limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
// returns a *float64 when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) GetIopsLimit()(*float64) {
    return m.iopsLimit
}
// GetMbpsLimit gets the mbpsLimit property value. Throughput limit for this volume in MB/s. If mbpsLimit is set to -1, then the volume has no MBPS limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
// returns a *float64 when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) GetMbpsLimit()(*float64) {
    return m.mbpsLimit
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("iopsLimit", m.GetIopsLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("mbpsLimit", m.GetMbpsLimit())
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
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIopsLimit sets the iopsLimit property value. IOPS limit for this volume. If iopsLimit is set to -1, then the volume has no IOPS limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) SetIopsLimit(value *float64)() {
    m.iopsLimit = value
}
// SetMbpsLimit sets the mbpsLimit property value. Throughput limit for this volume in MB/s. If mbpsLimit is set to -1, then the volume has no MBPS limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited. If both iopsLimit and mbpsLimit are specified, mbpsLimit must not be hit before iopsLimit. In other words, IOPS and MBPS limits should honor iopsLimit &lt;= ((mbpsLimit MB/s * 2^20 B/MB) / block_size B). Signed 64-bit integer.
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qos) SetMbpsLimit(value *float64)() {
    m.mbpsLimit = value
}
type V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfo_qosable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIopsLimit()(*float64)
    GetMbpsLimit()(*float64)
    SetIopsLimit(value *float64)()
    SetMbpsLimit(value *float64)()
}
