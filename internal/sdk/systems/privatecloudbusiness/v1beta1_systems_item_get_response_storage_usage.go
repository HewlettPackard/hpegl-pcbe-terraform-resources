package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemGetResponse_storageUsage storage Usage Information of system.
type V1beta1SystemsItemGetResponse_storageUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Storage Capacity of the system
    sizeInBytes *float64
    // Storage Usage of the system
    usageInBytes *float64
}
// NewV1beta1SystemsItemGetResponse_storageUsage instantiates a new V1beta1SystemsItemGetResponse_storageUsage and sets the default values.
func NewV1beta1SystemsItemGetResponse_storageUsage()(*V1beta1SystemsItemGetResponse_storageUsage) {
    m := &V1beta1SystemsItemGetResponse_storageUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemGetResponse_storageUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemGetResponse_storageUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemGetResponse_storageUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemGetResponse_storageUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemGetResponse_storageUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["usageInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageInBytes(val)
        }
        return nil
    }
    return res
}
// GetSizeInBytes gets the sizeInBytes property value. Storage Capacity of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageUsage) GetSizeInBytes()(*float64) {
    return m.sizeInBytes
}
// GetUsageInBytes gets the usageInBytes property value. Storage Usage of the system
// returns a *float64 when successful
func (m *V1beta1SystemsItemGetResponse_storageUsage) GetUsageInBytes()(*float64) {
    return m.usageInBytes
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemGetResponse_storageUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("usageInBytes", m.GetUsageInBytes())
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
func (m *V1beta1SystemsItemGetResponse_storageUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSizeInBytes sets the sizeInBytes property value. Storage Capacity of the system
func (m *V1beta1SystemsItemGetResponse_storageUsage) SetSizeInBytes(value *float64)() {
    m.sizeInBytes = value
}
// SetUsageInBytes sets the usageInBytes property value. Storage Usage of the system
func (m *V1beta1SystemsItemGetResponse_storageUsage) SetUsageInBytes(value *float64)() {
    m.usageInBytes = value
}
type V1beta1SystemsItemGetResponse_storageUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSizeInBytes()(*float64)
    GetUsageInBytes()(*float64)
    SetSizeInBytes(value *float64)()
    SetUsageInBytes(value *float64)()
}
