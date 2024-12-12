package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics capacity metrics of the storage array.
type V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availableBytes property
    availableBytes *float64
    // The rawCapacityBytes property
    rawCapacityBytes *float64
    // The usableCapacityBytes property
    usableCapacityBytes *float64
    // The usageBytes property
    usageBytes *float64
}
// NewV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics instantiates a new V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics and sets the default values.
func NewV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics()(*V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) {
    m := &V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableBytes gets the availableBytes property value. The availableBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetAvailableBytes()(*float64) {
    return m.availableBytes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableBytes(val)
        }
        return nil
    }
    res["rawCapacityBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawCapacityBytes(val)
        }
        return nil
    }
    res["usableCapacityBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsableCapacityBytes(val)
        }
        return nil
    }
    res["usageBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageBytes(val)
        }
        return nil
    }
    return res
}
// GetRawCapacityBytes gets the rawCapacityBytes property value. The rawCapacityBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetRawCapacityBytes()(*float64) {
    return m.rawCapacityBytes
}
// GetUsableCapacityBytes gets the usableCapacityBytes property value. The usableCapacityBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetUsableCapacityBytes()(*float64) {
    return m.usableCapacityBytes
}
// GetUsageBytes gets the usageBytes property value. The usageBytes property
// returns a *float64 when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) GetUsageBytes()(*float64) {
    return m.usageBytes
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("availableBytes", m.GetAvailableBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("rawCapacityBytes", m.GetRawCapacityBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("usableCapacityBytes", m.GetUsableCapacityBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("usageBytes", m.GetUsageBytes())
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
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableBytes sets the availableBytes property value. The availableBytes property
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) SetAvailableBytes(value *float64)() {
    m.availableBytes = value
}
// SetRawCapacityBytes sets the rawCapacityBytes property value. The rawCapacityBytes property
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) SetRawCapacityBytes(value *float64)() {
    m.rawCapacityBytes = value
}
// SetUsableCapacityBytes sets the usableCapacityBytes property value. The usableCapacityBytes property
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) SetUsableCapacityBytes(value *float64)() {
    m.usableCapacityBytes = value
}
// SetUsageBytes sets the usageBytes property value. The usageBytes property
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetrics) SetUsageBytes(value *float64)() {
    m.usageBytes = value
}
type V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableBytes()(*float64)
    GetRawCapacityBytes()(*float64)
    GetUsableCapacityBytes()(*float64)
    GetUsageBytes()(*float64)
    SetAvailableBytes(value *float64)()
    SetRawCapacityBytes(value *float64)()
    SetUsableCapacityBytes(value *float64)()
    SetUsageBytes(value *float64)()
}
