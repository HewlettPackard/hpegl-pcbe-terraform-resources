package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemStorageArraysGetResponse_items list of a system's associated storage arrays.For Storage array,  overall health deduced based on the health of itsshelves, controllers, disks, network interfaces, power supply, fan and temperature sensors
type V1beta1SystemsItemStorageArraysGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Capacity metrics of the storage array.
    capacityMetrics V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable
    // Details regarding the Storage Array
    details V1beta1SystemsItemStorageArraysGetResponse_items_detailsable
    // Various storage array components' health.
    health V1beta1SystemsItemStorageArraysGetResponse_items_healthable
}
// NewV1beta1SystemsItemStorageArraysGetResponse_items instantiates a new V1beta1SystemsItemStorageArraysGetResponse_items and sets the default values.
func NewV1beta1SystemsItemStorageArraysGetResponse_items()(*V1beta1SystemsItemStorageArraysGetResponse_items) {
    m := &V1beta1SystemsItemStorageArraysGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStorageArraysGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageArraysGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageArraysGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapacityMetrics gets the capacityMetrics property value. Capacity metrics of the storage array.
// returns a V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) GetCapacityMetrics()(V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable) {
    return m.capacityMetrics
}
// GetDetails gets the details property value. Details regarding the Storage Array
// returns a V1beta1SystemsItemStorageArraysGetResponse_items_detailsable when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) GetDetails()(V1beta1SystemsItemStorageArraysGetResponse_items_detailsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capacityMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacityMetrics(val.(V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable))
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemStorageArraysGetResponse_items_detailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(V1beta1SystemsItemStorageArraysGetResponse_items_detailsable))
        }
        return nil
    }
    res["health"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemStorageArraysGetResponse_items_healthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(V1beta1SystemsItemStorageArraysGetResponse_items_healthable))
        }
        return nil
    }
    return res
}
// GetHealth gets the health property value. Various storage array components' health.
// returns a V1beta1SystemsItemStorageArraysGetResponse_items_healthable when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) GetHealth()(V1beta1SystemsItemStorageArraysGetResponse_items_healthable) {
    return m.health
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("capacityMetrics", m.GetCapacityMetrics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("health", m.GetHealth())
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
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapacityMetrics sets the capacityMetrics property value. Capacity metrics of the storage array.
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) SetCapacityMetrics(value V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable)() {
    m.capacityMetrics = value
}
// SetDetails sets the details property value. Details regarding the Storage Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) SetDetails(value V1beta1SystemsItemStorageArraysGetResponse_items_detailsable)() {
    m.details = value
}
// SetHealth sets the health property value. Various storage array components' health.
func (m *V1beta1SystemsItemStorageArraysGetResponse_items) SetHealth(value V1beta1SystemsItemStorageArraysGetResponse_items_healthable)() {
    m.health = value
}
type V1beta1SystemsItemStorageArraysGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacityMetrics()(V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable)
    GetDetails()(V1beta1SystemsItemStorageArraysGetResponse_items_detailsable)
    GetHealth()(V1beta1SystemsItemStorageArraysGetResponse_items_healthable)
    SetCapacityMetrics(value V1beta1SystemsItemStorageArraysGetResponse_items_capacityMetricsable)()
    SetDetails(value V1beta1SystemsItemStorageArraysGetResponse_items_detailsable)()
    SetHealth(value V1beta1SystemsItemStorageArraysGetResponse_items_healthable)()
}
